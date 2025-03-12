package event

import (
	"context"
	"fmt"
	"math/big"
	"time"

	"github.com/ethereum/go-ethereum/accounts/abi/bind"
	"github.com/ethereum/go-ethereum/common"
	"github.com/ethereum/go-ethereum/core/types"
	"github.com/ethereum/go-ethereum/crypto"
	"github.com/haitao-sun03/donation/backend-end/abi"
	"github.com/haitao-sun03/donation/backend-end/config"
	"github.com/haitao-sun03/donation/backend-end/model"
	"github.com/haitao-sun03/donation/backend-end/utils"
	log "github.com/sirupsen/logrus"
	"gorm.io/gorm"
	"gorm.io/gorm/clause"
)

// CampaignHandler 处理捐赠相关事件
type CampaignHandler struct {
}

// StartCampaignRecord 活动创建事件
type StartCampaignRecord struct {
	CampaignID     *big.Int
	Starter        common.Address
	Title          string
	Goal           *big.Int
	StartTime      *big.Int
	EndTime        *big.Int
	Status         uint8
	Nature         uint8
	Beneficiary    uint8
	Purpose        string
	ExpectedImpact string
	BaseEvent
}

func (h *CampaignHandler) Parse(vLog types.Log) (interface{}, error) {
	donationManage := config.DonationManageContract.DonationManageFilterer
	switch vLog.Topics[0] {

	case crypto.Keccak256Hash([]byte("StartCampaign(uint256,address,string,uint256,uint256,uint256,uint8,uint8,string,string,uint8)")):
		return donationManage.ParseStartCampaign(vLog)
	case crypto.Keccak256Hash([]byte("ActiveCampaign(uint256,address,uint256)")):
		return donationManage.ParseActiveCampaign(vLog)
	case crypto.Keccak256Hash([]byte("CompletedCampaign(uint256,address,uint256)")):
		return donationManage.ParseCompletedCampaign(vLog)
	case crypto.Keccak256Hash([]byte("CancellCampaign(uint256,address,uint256)")):
		return donationManage.ParseCancellCampaign(vLog)
	case crypto.Keccak256Hash([]byte("Withdraw(uint256,address,uint256,uint256)")):
		return donationManage.ParseWithdraw(vLog)

	default:
		return nil, fmt.Errorf("unsupported campaign event")
	}
}

func (h *CampaignHandler) Handle(data any) error {
	switch event := data.(type) {
	case *abi.DonationManageStartCampaign:
		return saveCampaignRecordAndUpdateAuth(event)
	case *abi.DonationManageActiveCampaign:
		return updateCampaignStatus(event)
	case *abi.DonationManageCompletedCampaign:
		return updateCampaignStatus(event)
	case *abi.DonationManageCancellCampaign:
		return updateCampaignStatus(event)
	case *abi.DonationManageWithdraw:
		return updateCampaignWithdraw(event)

	default:
		return fmt.Errorf("invalid campaign event data")
	}
}

func saveCampaignRecordAndUpdateAuth(donationManageStartCampaign *abi.DonationManageStartCampaign) error {
	// 获取带上下文和超时的DB实例
	db := config.GetDB().WithContext(context.Background())
	return db.Transaction(func(tx *gorm.DB) error {
		if err := tx.Create(&model.CampaignModel{
			CampaignID:     donationManageStartCampaign.Id.String(),
			Title:          donationManageStartCampaign.Title,
			Goal:           donationManageStartCampaign.Goal.String(),                   // *big.Int to string
			StartTime:      time.Unix(donationManageStartCampaign.StartTime.Int64(), 0), // *big.Int to time.Time
			EndTime:        time.Unix(donationManageStartCampaign.EndTime.Int64(), 0),   // *big.Int to time.Time
			Status:         donationManageStartCampaign.Status,
			Starter:        donationManageStartCampaign.Starter.Hex(), // common.Address to string
			Nature:         donationManageStartCampaign.Nature,
			Beneficiary:    donationManageStartCampaign.Beneficiary,
			Purpose:        donationManageStartCampaign.Purpose,
			ExpectedImpact: donationManageStartCampaign.ExpectedImpact,
		}).Error; err != nil {
			return err
		}

		if !donationManageStartCampaign.Id.IsInt64() {
			return fmt.Errorf("CampaignID %s exceeds int64 range", donationManageStartCampaign.Id.String())
		}
		campaignIdInt64 := donationManageStartCampaign.Id.Int64()

		if res := tx.Exec("INSERT INTO user_activity_roles (address, campaign_id, role) VALUES (?, ?, ?)",
			donationManageStartCampaign.Starter.Hex(),
			campaignIdInt64,
			model.StarterRole); res.Error != nil {
			return res.Error
		}

		return nil
	})

}

// CampaignIsWithdrawRecord 活动取款
type CampaignIsWithdrawRecord struct {
	CampaignID *big.Int
	Caller     common.Address
	Timestamp  *big.Int
	IsWithdraw uint8
	BaseEvent
}

type CampaignEvent interface {
	GetCampaignID() *big.Int
	GetStatus() uint8
}

func updateCampaignStatus(record CampaignEvent) error {
	// 实现活动状态更新逻辑
	db := config.GetDB().WithContext(context.Background())
	return db.Transaction(func(tx *gorm.DB) error {
		var existing model.CampaignModel

		// 查询现有记录（带行锁）
		err := tx.
			Clauses(clause.Locking{Strength: "UPDATE"}).
			Where("campaign_id = ?", record.GetCampaignID().String()).
			First(&existing).Error

		// db.Transaction回调方法中需要返回error，若为nil，则提交事务，否则rollback
		switch {
		// err包括记录没有查找到的情况
		case err != nil:
			return err
		default:
			// 更新记录
			if err = tx.Model(&existing).
				Updates(map[string]interface{}{
					"status": record.GetStatus(),
				}).Error; err != nil {
				return err
			}
		}
		// 如果活动状态为已完成，则检查是否需要铸造NFT
		if record.GetStatus() == model.CampaignStatusCompleted {
			err = checkMintNFTOfTheCampaign(tx, record, existing)
			if err != nil {
				return err
			}
		}
		return nil
	})
}

func checkMintNFTOfTheCampaign(tx *gorm.DB, record CampaignEvent, campaign model.CampaignModel) error {
	// 查询该活动的所有donations
	var donations []model.DonationModel
	if err := tx.Where("campaign_id = ?", record.GetCampaignID().String()).Find(&donations).Error; err != nil {
		return err
	}

	for _, donation := range donations {
		donationAmount, _ := new(big.Int).SetString(donation.Amount, 10)
		campaignGoal, _ := new(big.Int).SetString(campaign.Goal, 10)
		ratio := new(big.Float).Quo(new(big.Float).SetInt(donationAmount), new(big.Float).SetInt(campaignGoal))
		percentage := new(big.Float).Mul(ratio, big.NewFloat(100))

		thresholds := []struct {
			percent int
			level   string
		}{
			{80, model.Diamond},
			{50, model.Gold},
			{20, model.Silver},
		}
		for _, threshold := range thresholds {
			if percentage.Cmp(big.NewFloat(float64(threshold.percent))) >= 0 {
				// 达到或超过阈值，铸造NFT
				log.Infof("Minting NFT for %s at %d,donor :%s", threshold.level, threshold.percent, donation.Donor)
				// ABI 编码
				NftABI := getContractABI(NftContract)
				if _, ok := NftABI.Methods["safeMint"]; !ok {
					log.Error("safeMint method not found in ABI")
					return fmt.Errorf("invalid ABI: safeMint not defined")
				}
				callData, err := NftABI.Pack("safeMint", common.HexToAddress(donation.Donor), threshold.level)
				if err != nil {
					log.WithError(err).Error("Failed to pack ABI data")
					return err
				}
				log.Debugf("Packed callData: %x", callData)
				auth, err := utils.GetTransactOpts(0, common.HexToAddress(config.Config.Geth.NftContract.Address), big.NewInt(0), callData)
				if err != nil {
					log.WithError(err).Error("Failed to get transact opts")
					return err
				}
				tran, err := config.NFTContract.SafeMint(auth, common.HexToAddress(donation.Donor), threshold.level)
				if err != nil {
					log.WithError(err).Error("Failed to mint NFT")
					return err
				}
				receipt, err := bind.WaitMined(context.Background(), config.GethClient, tran)
				if err != nil {
					log.WithError(err).Error("Failed to wait for NFT minting transaction to be mined")
					return err
				}
				log.Infof("Gas Used: %d", receipt.GasUsed)

				// 更新表
				err = tx.Model(&donation).
					Updates(map[string]interface{}{
						"mint_level": threshold.level,
					}).Error
				if err != nil {
					return err
				}
				break
			}
		}
	}
	return nil
}

func updateCampaignWithdraw(record *abi.DonationManageWithdraw) error {
	// 实现活动状态更新逻辑
	db := config.GetDB().WithContext(context.Background())
	return db.Transaction(func(tx *gorm.DB) error {
		var existing model.CampaignModel

		// 查询现有记录（带行锁）
		err := tx.
			Clauses(clause.Locking{Strength: "UPDATE"}).
			Where("campaign_id = ?", record.Id.String()).
			First(&existing).Error

		// db.Transaction回调方法中需要返回error，若为nil，则提交事务，否则rollback
		switch {
		// err包括记录没有查找到的情况
		case err != nil:
			return err

		default:

			// 更新记录
			return tx.Model(&existing).
				Updates(map[string]any{
					"is_withdraw": model.Withdraw,
				}).Error
		}
	})
}
