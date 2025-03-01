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
	"github.com/haitao-sun03/go/config"
	"github.com/haitao-sun03/go/model"
	"github.com/haitao-sun03/go/utils"
	log "github.com/sirupsen/logrus"
	"gorm.io/gorm"
	"gorm.io/gorm/clause"
)

// CampaignHandler 处理捐赠相关事件
type CampaignHandler struct {
	EventHandler
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
	switch vLog.Topics[0] {
	case crypto.Keccak256Hash([]byte("StartCampaign(uint256,address,string,uint256,uint256,uint256,uint8,uint8,string,string,uint8)")):
		return parseStartCampaignEvent(vLog)
	case crypto.Keccak256Hash([]byte("CancellCampaign(uint256,address,uint256)")):
		return parseCancelCampaignEvent(vLog)
	case crypto.Keccak256Hash([]byte("CompletedCampaign(uint256,address,uint256)")):
		return parseCompleteCampaignEvent(vLog)
	case crypto.Keccak256Hash([]byte("Withdraw(uint256,address,uint256,uint256)")):
		return parseWithdrawEvent(vLog)
	case crypto.Keccak256Hash([]byte("ActiveCampaign(uint256,address,uint256)")):
		return parseActiveCampaignEvent(vLog)
	default:
		return nil, fmt.Errorf("unsupported campaign event")
	}
}

func (h *CampaignHandler) Handle(data any) error {
	switch event := data.(type) {
	case *StartCampaignRecord:
		return saveCampaignRecordAndUpdateAuth(event)
	case *CampaignStatusRecord:
		return updateCampaignStatus(event)
	case *CampaignIsWithdrawRecord:
		return updateCampaignWithdraw(event)

	default:
		return fmt.Errorf("invalid campaign event data")
	}
}

func parseStartCampaignEvent(vLog types.Log) (*StartCampaignRecord, error) {
	contractAbi := getContractABI(CampaignRelContract)

	var event struct {
		Id             *big.Int
		Starter        common.Address
		Title          string
		Goal           *big.Int
		StartTime      *big.Int
		EndTime        *big.Int
		Nature         uint8
		Beneficiary    uint8
		Purpose        string
		ExpectedImpact string
		Status         uint8
	}

	if err := contractAbi.UnpackIntoInterface(&event, StartCampaignEvent, vLog.Data); err != nil {
		return nil, fmt.Errorf("failed to unpack StartCampaign event: %v", err)
	}

	campaignId := new(big.Int).SetBytes(vLog.Topics[1].Bytes())

	header, err := getBlockHeader(vLog.BlockHash)
	if err != nil {
		return nil, err
	}

	return &StartCampaignRecord{
		CampaignID:     campaignId,
		Starter:        event.Starter,
		Title:          event.Title,
		Goal:           event.Goal,
		StartTime:      event.StartTime,
		EndTime:        event.EndTime,
		Nature:         event.Nature,
		Beneficiary:    event.Beneficiary,
		Purpose:        event.Purpose,
		ExpectedImpact: event.ExpectedImpact,
		Status:         event.Status,
		BaseEvent: BaseEvent{
			EventType: StartCampaignEvent,
			BlockTime: header.Time,
			TxHash:    vLog.TxHash,
		},
	}, nil
}

func saveCampaignRecordAndUpdateAuth(StartCampaignRecord *StartCampaignRecord) error {
	// 获取带上下文和超时的DB实例
	db := config.GetDB().WithContext(context.Background())
	return db.Transaction(func(tx *gorm.DB) error {
		if err := tx.Create(&model.CampaignModel{
			CampaignID:     StartCampaignRecord.CampaignID.String(),
			Title:          StartCampaignRecord.Title,
			Goal:           StartCampaignRecord.Goal.String(),                   // *big.Int to string
			StartTime:      time.Unix(StartCampaignRecord.StartTime.Int64(), 0), // *big.Int to time.Time
			EndTime:        time.Unix(StartCampaignRecord.EndTime.Int64(), 0),   // *big.Int to time.Time
			Status:         StartCampaignRecord.Status,
			Starter:        StartCampaignRecord.Starter.Hex(), // common.Address to string
			Nature:         StartCampaignRecord.Nature,
			Beneficiary:    StartCampaignRecord.Beneficiary,
			Purpose:        StartCampaignRecord.Purpose,
			ExpectedImpact: StartCampaignRecord.ExpectedImpact,
		}).Error; err != nil {
			return err
		}

		if !StartCampaignRecord.CampaignID.IsInt64() {
			return fmt.Errorf("CampaignID %s exceeds int64 range", StartCampaignRecord.CampaignID.String())
		}
		campaignIdInt64 := StartCampaignRecord.CampaignID.Int64()

		if res := tx.Exec("INSERT INTO user_activity_roles (address, campaign_id, role) VALUES (?, ?, ?)",
			StartCampaignRecord.Starter.Hex(),
			campaignIdInt64,
			model.StarterRole); res.Error != nil {
			return res.Error
		}

		return nil
	})

}

// CampaignStatusRecord 活动状态变更事件
type CampaignStatusRecord struct {
	CampaignID *big.Int
	Caller     common.Address
	Timestamp  *big.Int
	Status     uint8
	IsWithdraw uint8
	BaseEvent
}

// CampaignIsWithdrawRecord 活动取款
type CampaignIsWithdrawRecord struct {
	CampaignID *big.Int
	Caller     common.Address
	Timestamp  *big.Int
	IsWithdraw uint8
	BaseEvent
}

func parseActiveCampaignEvent(vLog types.Log) (*CampaignStatusRecord, error) {
	contractAbi := getContractABI(CampaignRelContract)

	var event struct {
		Id     *big.Int
		Caller common.Address
		Time   *big.Int
	}

	if err := contractAbi.UnpackIntoInterface(&event, ActiveCampaignEvent, vLog.Data); err != nil {
		return nil, fmt.Errorf("failed to unpack ActiveCampaignEvent event: %v", err)
	}

	campaignId := new(big.Int).SetBytes(vLog.Topics[1].Bytes())

	header, err := getBlockHeader(vLog.BlockHash)
	if err != nil {
		return nil, err
	}

	return &CampaignStatusRecord{
		CampaignID: campaignId,
		Caller:     event.Caller,
		Timestamp:  event.Time,
		Status:     model.CampaignStatusActive,
		BaseEvent: BaseEvent{
			EventType: ActiveCampaignEvent,
			BlockTime: header.Time,
			TxHash:    vLog.TxHash,
		},
	}, nil
}

func parseCancelCampaignEvent(vLog types.Log) (*CampaignStatusRecord, error) {
	contractAbi := getContractABI(CampaignRelContract)

	var event struct {
		Id     *big.Int
		Caller common.Address
		Time   *big.Int
	}

	if err := contractAbi.UnpackIntoInterface(&event, CancelledCampaignEvent, vLog.Data); err != nil {
		return nil, fmt.Errorf("failed to unpack CancellCampaign event: %v", err)
	}

	campaignId := new(big.Int).SetBytes(vLog.Topics[1].Bytes())

	header, err := getBlockHeader(vLog.BlockHash)
	if err != nil {
		return nil, err
	}

	return &CampaignStatusRecord{
		CampaignID: campaignId,
		Caller:     event.Caller,
		Timestamp:  event.Time,
		Status:     model.CampaignStatusCancelled,
		BaseEvent: BaseEvent{
			EventType: CancelledCampaignEvent,
			BlockTime: header.Time,
			TxHash:    vLog.TxHash,
		},
	}, nil
}

func parseCompleteCampaignEvent(vLog types.Log) (*CampaignStatusRecord, error) {
	contractAbi := getContractABI(CampaignRelContract)

	var event struct {
		Id     *big.Int
		Caller common.Address
		Time   *big.Int
	}

	if err := contractAbi.UnpackIntoInterface(&event, CompletedCampaignEvent, vLog.Data); err != nil {
		return nil, fmt.Errorf("failed to unpack CompletedCampaign event: %v", err)
	}

	campaignId := new(big.Int).SetBytes(vLog.Topics[1].Bytes())

	header, err := getBlockHeader(vLog.BlockHash)
	if err != nil {
		return nil, err
	}

	return &CampaignStatusRecord{
		CampaignID: campaignId,
		Caller:     event.Caller,
		Timestamp:  event.Time,
		Status:     model.CampaignStatusCompleted,
		BaseEvent: BaseEvent{
			EventType: CompletedCampaignEvent,
			BlockTime: header.Time,
			TxHash:    vLog.TxHash,
		},
	}, nil
}

func parseWithdrawEvent(vLog types.Log) (*CampaignIsWithdrawRecord, error) {
	contractAbi := getContractABI(CampaignRelContract)

	var event struct {
		Id             *big.Int
		Withdrawer     common.Address
		Time           *big.Int
		WithdrawAmount *big.Int
	}

	if err := contractAbi.UnpackIntoInterface(&event, WithdrawEvent, vLog.Data); err != nil {
		return nil, fmt.Errorf("failed to unpack Withdraw event: %v", err)
	}

	campaignId := new(big.Int).SetBytes(vLog.Topics[1].Bytes())

	header, err := getBlockHeader(vLog.BlockHash)
	if err != nil {
		return nil, err
	}

	return &CampaignIsWithdrawRecord{
		CampaignID: campaignId,
		Caller:     event.Withdrawer,
		Timestamp:  event.Time,
		IsWithdraw: model.Withdraw,
		BaseEvent: BaseEvent{
			EventType: WithdrawEvent,
			BlockTime: header.Time,
			TxHash:    vLog.TxHash,
		},
	}, nil
}

func updateCampaignStatus(record *CampaignStatusRecord) error {
	// 实现活动状态更新逻辑
	db := config.GetDB().WithContext(context.Background())
	return db.Transaction(func(tx *gorm.DB) error {
		var existing model.CampaignModel

		// 查询现有记录（带行锁）
		err := tx.
			Clauses(clause.Locking{Strength: "UPDATE"}).
			Where("campaign_id = ?", record.CampaignID.String()).
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
					"status": record.Status,
				}).Error; err != nil {
				return err
			}
		}
		// 如果活动状态为已完成，则检查是否需要铸造NFT
		if record.Status == model.CampaignStatusCompleted {
			err = checkMintNFTOfTheCampaign(tx, record, existing)
			if err != nil {
				return err
			}
		}
		return nil
	})
}

func checkMintNFTOfTheCampaign(tx *gorm.DB, record *CampaignStatusRecord, campaign model.CampaignModel) error {
	// 查询该活动的所有donations
	var donations []model.DonationModel
	if err := tx.Where("campaign_id = ?", record.CampaignID.String()).Find(&donations).Error; err != nil {
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

				auth, err := utils.GetTransactOpts(2000000)
				if err != nil {
					log.WithError(err).Error("Failed to get transact opts")
					return err
				}
				tran, err := config.NFTContract.SafeMint(auth, common.HexToAddress(donation.Donor), threshold.level)
				if err != nil {
					log.WithError(err).Error("Failed to mint NFT")
					return err
				}
				if _, err = bind.WaitMined(context.Background(), config.GethClient, tran); err != nil {
					log.WithError(err).Error("Failed to wait for NFT minting transaction to be mined")
					return err
				}

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

func updateCampaignWithdraw(record *CampaignIsWithdrawRecord) error {
	// 实现活动状态更新逻辑
	db := config.GetDB().WithContext(context.Background())
	return db.Transaction(func(tx *gorm.DB) error {
		var existing model.CampaignModel

		// 查询现有记录（带行锁）
		err := tx.
			Clauses(clause.Locking{Strength: "UPDATE"}).
			Where("campaign_id = ?", record.CampaignID.String()).
			First(&existing).Error

		// db.Transaction回调方法中需要返回error，若为nil，则提交事务，否则rollback
		switch {
		// err包括记录没有查找到的情况
		case err != nil:
			return err

		default:

			// 更新记录
			return tx.Model(&existing).
				Updates(map[string]interface{}{
					"is_withdraw": record.IsWithdraw,
				}).Error
		}
	})
}
