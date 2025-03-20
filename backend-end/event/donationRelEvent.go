package event

import (
	"context"
	"errors"
	"fmt"
	"math/big"

	"github.com/ethereum/go-ethereum/common"
	"github.com/ethereum/go-ethereum/core/types"
	"github.com/ethereum/go-ethereum/crypto"
	"github.com/haitao-sun03/donation/backend-end/abi"
	"github.com/haitao-sun03/donation/backend-end/config"
	"github.com/haitao-sun03/donation/backend-end/db"
	"github.com/haitao-sun03/donation/backend-end/model"
	"gorm.io/gorm"
	"gorm.io/gorm/clause"
)

// DonationHandler 处理捐赠相关事件
type DonationHandler struct {
	DB db.DBInterface // 可选注入，默认为 config.GetDB()
}

// NewDonationHandler 创建 DonationHandler 实例，带默认 DB
func NewDonationHandler(_db db.DBInterface) *DonationHandler {
	if _db == nil {
		_db = &db.GormDB{InnerDB: config.GetDB()}
	}
	return &DonationHandler{DB: _db}
}

// 捐赠事件结构体
type DonationRecord struct {
	CampaignID *big.Int
	Donor      common.Address
	Amount     *big.Int
	BlockTime  uint64
	BaseEvent
}

// 退款事件结构体
type RefundRecord struct {
	CampaignID *big.Int
	Caller     common.Address
	Timestamp  *big.Int
	IsRefund   uint8
	BaseEvent
}

func (h *DonationHandler) Parse(vLog types.Log) (interface{}, error) {
	donationManage := config.DonationManageContract.DonationManageFilterer
	switch vLog.Topics[0] {
	case crypto.Keccak256Hash([]byte("Donate(uint256,address,uint256)")):
		return donationManage.ParseDonate(vLog)
	case crypto.Keccak256Hash([]byte("Refund(uint256,address,uint256,uint256)")):
		return donationManage.ParseRefund(vLog)
	default:
		return nil, fmt.Errorf("unsupported event type")
	}
}

func (h *DonationHandler) Handle(data interface{}) error {

	switch event := data.(type) {
	case *abi.DonationManageDonate:
		return h.saveDonationRecord(event)
	case *abi.DonationManageRefund:
		return updateDonationRefund(event)

	default:
		return fmt.Errorf("invalid donation event data")
	}

}

func (h *DonationHandler) saveDonationRecord(donationRecord *abi.DonationManageDonate) error {
	// 获取带上下文和超时的DB实例
	db := h.DB.WithContext(context.Background())

	// 开启事务（确保原子性）
	return db.Transaction(func(tx *gorm.DB) error {
		// 1 更新活动总捐赠额
		var campaign model.CampaignModel
		err := tx.Clauses(clause.Locking{Strength: "UPDATE"}).
			Where("campaign_id = ?", donationRecord.Id.String()).
			First(&campaign).Error
		if err != nil {
			return err
		}

		// 执行大数运算
		oldTotal, _ := new(big.Int).SetString(campaign.TotalDonated, 10)
		newTotal := new(big.Int).Add(oldTotal, donationRecord.Amount)

		// 更新活动的总捐赠额
		err = tx.Model(&campaign).
			Where("campaign_id = ?", campaign.CampaignID).
			Updates(map[string]interface{}{
				"total_donated": newTotal.String(),
			}).Error
		if err != nil {
			return err
		}

		// 2 处理捐赠记录
		var existing model.DonationModel

		// 查询现有记录（带行锁）
		err = tx.
			Where("campaign_id = ?", donationRecord.Id.String()).
			Where("donor = ?", donationRecord.Donater.Hex()).
			First(&existing).Error

		// db.Transaction回调方法中需要返回error，若为nil，则提交事务，否则rollback
		if !donationRecord.Id.IsInt64() {
			return fmt.Errorf("CampaignID %s exceeds int64 range", donationRecord.Id.String())
		}
		campaignIdInt64 := donationRecord.Id.Int64()
		// 插入 user_activity_roles
		var userActivityRoleModel model.UserActivityRoleModel
		result := tx.
			Where("address = ?", donationRecord.Donater.Hex()).
			Where("campaign_id = ?", campaignIdInt64).
			Where("role = ?", model.DonorRole).
			First(&userActivityRoleModel)
		if errors.Is(result.Error, gorm.ErrRecordNotFound) || result.RowsAffected == 0 {
			if err := tx.Exec(
				"INSERT INTO user_activity_roles (address, campaign_id, role) VALUES (?, ?, ?)",
				donationRecord.Donater.Hex(),
				campaignIdInt64,
				model.DonorRole,
			).Error; err != nil {
				return err
			}
		}

		switch {
		case err == gorm.ErrRecordNotFound:
			// 该用户第一次捐赠，创建新的捐赠记录
			return tx.Create(&model.DonationModel{
				CampaignID: donationRecord.Id.String(),
				Donor:      donationRecord.Donater.Hex(),
				Amount:     donationRecord.Amount.String(),
				// BlockTime:  time.Unix(int64(donationRecord.BlockTime), 0),
			}).Error

		case err != nil:
			return err

		default:
			oldAmount, _ := new(big.Int).SetString(existing.Amount, 10)
			newAmount := new(big.Int).Add(oldAmount, donationRecord.Amount)

			// 更新记录
			return tx.Model(&existing).
				Where("campaign_id = ?", existing.CampaignID).
				Where("donor = ?", existing.Donor).
				Updates(map[string]interface{}{
					"amount": newAmount.String(),
					// "block_time": time.Unix(int64(donationRecord.BlockTime), 0),
				}).Error
		}

	})
}

func updateDonationRefund(record *abi.DonationManageRefund) error {
	// 实现活动状态更新逻辑
	db := config.GetDB().WithContext(context.Background())
	return db.Transaction(func(tx *gorm.DB) error {
		var existing model.DonationModel

		// 查询现有记录（带行锁）
		err := tx.
			Clauses(clause.Locking{Strength: "UPDATE"}).
			Where("campaign_id = ?", record.Id.String()).
			Where("donor = ?", record.Refunder.Hex()).
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
					"is_refund": model.Refund,
				}).Error
		}
	})
}
