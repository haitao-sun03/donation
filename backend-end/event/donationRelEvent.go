package event

import (
	"context"
	"fmt"
	"math/big"
	"time"

	"github.com/ethereum/go-ethereum/common"
	"github.com/ethereum/go-ethereum/core/types"
	"github.com/ethereum/go-ethereum/crypto"
	"github.com/haitao-sun03/go/config"
	"github.com/haitao-sun03/go/model"
	"gorm.io/gorm"
	"gorm.io/gorm/clause"
)

// DonationHandler 处理捐赠相关事件
type DonationHandler struct {
	EventHandler
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
	switch vLog.Topics[0] {
	case crypto.Keccak256Hash([]byte("Donate(uint256,address,uint256)")):
		return parseDonateEvent(vLog)
	case crypto.Keccak256Hash([]byte("Refund(uint256,address,uint256,uint256)")):
		return parseRefundEvent(vLog)
	default:
		return nil, fmt.Errorf("unsupported event type")
	}
}

func (h *DonationHandler) Handle(data interface{}) error {

	switch event := data.(type) {
	case *DonationRecord:
		return saveDonationRecord(event)
	case *RefundRecord:
		return updateDonationRefund(event)

	default:
		return fmt.Errorf("invalid donation event data")
	}

}

// parseDonateEvent 解析捐赠事件日志
func parseDonateEvent(vLog types.Log) (*DonationRecord, error) {
	contractAbi := getContractABI(CampaignRelContract)

	// 解包日志数据
	var event struct {
		CampaignId *big.Int
		Donater    common.Address
		Amount     *big.Int
	}

	if err := contractAbi.UnpackIntoInterface(&event, DonateEvent, vLog.Data); err != nil {
		return nil, fmt.Errorf("failed to unpack Donate event: %v", err)
	}

	campaignId := new(big.Int).SetBytes(vLog.Topics[1].Bytes())

	// 获取区块时间（需要额外查询）
	header, err := getBlockHeader(vLog.BlockHash)
	if err != nil {
		return nil, fmt.Errorf("failed to get block header: %v", err)
	}

	return &DonationRecord{
		CampaignID: campaignId,
		Donor:      event.Donater,
		Amount:     event.Amount,
		BlockTime:  header.Time,
		BaseEvent: BaseEvent{
			EventType: DonateEvent,
			BlockTime: header.Time,
			TxHash:    vLog.TxHash,
		},
	}, nil
}

// parseRefundEvent 退款事件日志
func parseRefundEvent(vLog types.Log) (*RefundRecord, error) {
	contractAbi := getContractABI(CampaignRelContract)

	// 解包日志数据
	var event struct {
		Refunder     common.Address
		Time         *big.Int
		RefundAmount *big.Int
	}

	if err := contractAbi.UnpackIntoInterface(&event, RefundEvent, vLog.Data); err != nil {
		return nil, fmt.Errorf("failed to unpack Refund event: %v", err)
	}

	campaignId := new(big.Int).SetBytes(vLog.Topics[1].Bytes())

	// 获取区块时间（需要额外查询）
	header, err := getBlockHeader(vLog.BlockHash)

	if err != nil {
		return nil, fmt.Errorf("failed to get block header: %v", err)
	}

	return &RefundRecord{
		CampaignID: campaignId,
		Caller:     event.Refunder,
		IsRefund:   model.Refund,
		BaseEvent: BaseEvent{
			EventType: RefundEvent,
			BlockTime: header.Time,
			TxHash:    vLog.TxHash,
		},
	}, nil
}

func saveDonationRecord(donationRecord *DonationRecord) error {
	// 获取带上下文和超时的DB实例
	db := config.GetDB().WithContext(context.Background())

	// 开启事务（确保原子性）
	return db.Transaction(func(tx *gorm.DB) error {
		// 1 更新活动总捐赠额
		var campaign model.CampaignModel
		err := tx.Clauses(clause.Locking{Strength: "UPDATE"}).
			Where("campaign_id = ?", donationRecord.CampaignID.String()).
			First(&campaign).Error
		if err != nil {
			return err
		}

		// 执行大数运算
		oldTotal, _ := new(big.Int).SetString(campaign.TotalDonated, 10)
		newTotal := new(big.Int).Add(oldTotal, donationRecord.Amount)

		// 更新活动的总捐赠额
		err = tx.Model(&campaign).
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
			Where("campaign_id = ?", donationRecord.CampaignID.String()).
			Where("donor = ?", donationRecord.Donor.Hex()).
			First(&existing).Error

		// db.Transaction回调方法中需要返回error，若为nil，则提交事务，否则rollback
		if !donationRecord.CampaignID.IsInt64() {
			return fmt.Errorf("CampaignID %s exceeds int64 range", donationRecord.CampaignID.String())
		}
		campaignIdInt64 := donationRecord.CampaignID.Int64()
		// 插入 user_activity_roles
		if err := tx.Exec(
			"INSERT INTO user_activity_roles (address, campaign_id, role) VALUES (?, ?, ?)",
			donationRecord.Donor.Hex(),
			campaignIdInt64,
			model.DonorRole,
		).Error; err != nil {
			return err
		}

		switch {
		case err == gorm.ErrRecordNotFound:
			// 该用户第一次捐赠，创建新的捐赠记录
			return tx.Create(&model.DonationModel{
				CampaignID: donationRecord.CampaignID.String(),
				Donor:      donationRecord.Donor.Hex(),
				Amount:     donationRecord.Amount.String(),
				BlockTime:  time.Unix(int64(donationRecord.BlockTime), 0),
			}).Error

		case err != nil:
			return err

		default:
			oldAmount, _ := new(big.Int).SetString(existing.Amount, 10)
			newAmount := new(big.Int).Add(oldAmount, donationRecord.Amount)

			// 更新记录
			return tx.Model(&existing).
				Updates(map[string]interface{}{
					"amount":     newAmount.String(),
					"block_time": time.Unix(int64(donationRecord.BlockTime), 0),
				}).Error
		}

	})
}

func updateDonationRefund(record *RefundRecord) error {
	// 实现活动状态更新逻辑
	db := config.GetDB().WithContext(context.Background())
	return db.Transaction(func(tx *gorm.DB) error {
		var existing model.DonationModel

		// 查询现有记录（带行锁）
		err := tx.
			Clauses(clause.Locking{Strength: "UPDATE"}).
			Where("campaign_id = ?", record.CampaignID.String()).
			Where("donor = ?", record.Caller.Hex()).
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
