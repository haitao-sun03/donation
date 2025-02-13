package event

import (
	"context"
	"fmt"
	"io/ioutil"
	"math/big"
	"strings"
	"time"

	"github.com/ethereum/go-ethereum"
	"github.com/ethereum/go-ethereum/accounts/abi"
	"github.com/ethereum/go-ethereum/common"
	"github.com/ethereum/go-ethereum/core/types"
	"github.com/ethereum/go-ethereum/crypto"
	"github.com/ethereum/go-ethereum/ethclient"
	"github.com/haitao-sun03/go/config"
	"github.com/haitao-sun03/go/model"
	log "github.com/sirupsen/logrus"
	"gorm.io/gorm"
	"gorm.io/gorm/clause"
)

func ListenEvents() {
	DonationContractAddress := common.HexToAddress(config.Config.Geth.DonationContract.Address)
	ListenContractEvents(context.Background(), config.GethWsClient, DonationContractAddress)
}

// 监听所有合约事件
func ListenContractEvents(ctx context.Context, client *ethclient.Client, contractAddr common.Address) {
	query := ethereum.FilterQuery{
		Addresses: []common.Address{contractAddr},
	}

	logs := make(chan types.Log)
	sub, err := client.SubscribeFilterLogs(ctx, query, logs)
	if err != nil {
		log.Fatalf("Failed to subscribe to logs: %v", err)
	}

	for {
		select {
		case err := <-sub.Err():
			log.Errorf("Subscription error: %v", err)
		case <-ctx.Done():
			sub.Unsubscribe()
			log.Info("Event listener stopped")
			return
		case vLog := <-logs:
			go handleLog(vLog)
		}
	}
}

func handleLog(vLog types.Log) {
	handler, exists := eventHandlers[vLog.Topics[0]]
	if !exists {
		log.Warnf("Unhandled event signature: %s", vLog.Topics[0].Hex())
		return
	}

	data, err := handler.Parse(vLog)
	if err != nil {
		log.Errorf("Failed to parse event: %v", err)
		return
	}

	if err := handler.Handle(data); err != nil {
		log.Errorf("Failed to handle event: %v", err)
	}
}

// 事件处理接口
type EventHandler interface {
	Parse(log types.Log) (interface{}, error)
	Handle(data interface{}) error
}

// 事件类型枚举
const (
	DonateEvent           = "Donate"
	StartCampaignEvent    = "StartCampaign"
	CancelCampaignEvent   = "CancelCampaign"
	CompleteCampaignEvent = "CompleteCampaign"
	WithdrawEvent         = "Withdraw"
	RefundEvent           = "Refund"
)

// 事件处理器注册表
var eventHandlers = map[common.Hash]EventHandler{
	crypto.Keccak256Hash([]byte("Donate(uint256,address,uint256)")):                                     &DonationHandler{},
	crypto.Keccak256Hash([]byte("StartCampaign(uint256,address,string,uint256,uint256,uint256,uint8)")): &CampaignHandler{},
	crypto.Keccak256Hash([]byte("CancellCampaign(uint256,address,uint256)")):                            &CampaignHandler{},
	crypto.Keccak256Hash([]byte("CompletedCampaign(uint256,address,uint256)")):                          &CampaignHandler{},
	crypto.Keccak256Hash([]byte("Withdraw(uint256,address,uint256,uint256)")):                           &CampaignHandler{},
	crypto.Keccak256Hash([]byte("Refund(uint256,address,uint256,uint256)")):                             &DonationHandler{},
}

// 基础事件结构
type BaseEvent struct {
	EventType string
	BlockTime uint64
	TxHash    common.Hash
}

// DonationHandler 处理捐赠相关事件
type DonationHandler struct {
	EventHandler
}

// CampaignHandler 处理捐赠相关事件
type CampaignHandler struct {
	EventHandler
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

func (h *CampaignHandler) Parse(vLog types.Log) (interface{}, error) {
	switch vLog.Topics[0] {
	case crypto.Keccak256Hash([]byte("StartCampaign(uint256,address,string,uint256,uint256,uint256,uint8)")):
		return parseStartCampaignEvent(vLog)
	case crypto.Keccak256Hash([]byte("CancellCampaign(uint256,address,uint256)")):
		return parseCancelCampaignEvent(vLog)
	case crypto.Keccak256Hash([]byte("CompletedCampaign(uint256,address,uint256)")):
		return parseCompleteCampaignEvent(vLog)
	case crypto.Keccak256Hash([]byte("Withdraw(uint256,address,uint256,uint256)")):
		return parseWithdrawEvent(vLog)
	default:
		return nil, fmt.Errorf("unsupported campaign event")
	}
}

func (h *CampaignHandler) Handle(data interface{}) error {
	switch event := data.(type) {
	case *StartCampaignRecord:
		return saveCampaignRecord(event)
	case *CampaignStatusRecord:
		return updateCampaignStatus(event)
	case *CampaignIsWithdrawRecord:
		return updateCampaignWithdraw(event)

	default:
		return fmt.Errorf("invalid campaign event data")
	}
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

// parseDonateEvent 解析捐赠事件日志
func parseDonateEvent(vLog types.Log) (*DonationRecord, error) {
	contractAbi := getContractABI()

	// 解包日志数据
	var event struct {
		CampaignId *big.Int
		Donor      common.Address
		Amount     *big.Int
	}

	if err := contractAbi.UnpackIntoInterface(&event, "Donate", vLog.Data); err != nil {
		return nil, fmt.Errorf("failed to unpack StartCampaign event: %v", err)
	}

	campaignId := new(big.Int).SetBytes(vLog.Topics[1].Bytes())

	// 获取区块时间（需要额外查询）
	header, err := config.GethClient.HeaderByHash(context.Background(), vLog.BlockHash)
	if err != nil {
		return nil, fmt.Errorf("failed to get block header: %v", err)
	}

	return &DonationRecord{
		CampaignID: campaignId,
		Donor:      event.Donor,
		Amount:     event.Amount,
		BlockTime:  header.Time,
	}, nil
}

// parseRefundEvent 退款事件日志
func parseRefundEvent(vLog types.Log) (*RefundRecord, error) {
	contractAbi := getContractABI()

	// 解包日志数据
	var event struct {
		Refunder common.Address
		Time     *big.Int
	}

	if err := contractAbi.UnpackIntoInterface(&event, "Refund", vLog.Data); err != nil {
		return nil, fmt.Errorf("failed to unpack Refund event: %v", err)
	}

	campaignId := new(big.Int).SetBytes(vLog.Topics[1].Bytes())

	// 获取区块时间（需要额外查询）
	header, err := config.GethClient.HeaderByHash(context.Background(), vLog.BlockHash)
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
		// 更新活动总捐赠额（只需要一次）
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

		// 处理捐赠记录
		var existing model.DonationModel

		// 查询现有记录（带行锁）
		err = tx.
			Where("campaign_id = ?", donationRecord.CampaignID.String()).
			Where("donor = ?", donationRecord.Donor.Hex()).
			First(&existing).Error

		// db.Transaction回调方法中需要返回error，若为nil，则提交事务，否则rollback
		switch {
		case err == gorm.ErrRecordNotFound:
			// 创建新记录
			return tx.Create(&model.DonationModel{
				CampaignID: donationRecord.CampaignID.String(),
				Donor:      donationRecord.Donor.Hex(),
				Amount:     donationRecord.Amount.String(),
				BlockTime:  time.Unix(int64(donationRecord.BlockTime), 0),
			}).Error

		case err != nil:
			return err

		default:
			// 金额累加（大数运算）
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

// StartCampaignRecord 活动创建事件
type StartCampaignRecord struct {
	CampaignID *big.Int
	Starter    common.Address
	Title      string
	Goal       *big.Int
	StartTime  *big.Int
	EndTime    *big.Int
	Status     uint8
	BaseEvent
}

func parseStartCampaignEvent(vLog types.Log) (*StartCampaignRecord, error) {
	contractAbi := getContractABI()

	var event struct {
		Id        *big.Int
		Starter   common.Address
		Title     string
		Goal      *big.Int
		StartTime *big.Int
		EndTime   *big.Int
		Status    uint8
	}

	if err := contractAbi.UnpackIntoInterface(&event, "StartCampaign", vLog.Data); err != nil {
		return nil, fmt.Errorf("failed to unpack StartCampaign event: %v", err)
	}

	campaignId := new(big.Int).SetBytes(vLog.Topics[1].Bytes())

	header, err := getBlockHeader(vLog.BlockHash)
	if err != nil {
		return nil, err
	}

	return &StartCampaignRecord{
		CampaignID: campaignId,
		Starter:    event.Starter,
		Title:      event.Title,
		Goal:       event.Goal,
		StartTime:  event.StartTime,
		EndTime:    event.EndTime,
		Status:     event.Status,
		BaseEvent: BaseEvent{
			EventType: StartCampaignEvent,
			BlockTime: header.Time,
			TxHash:    vLog.TxHash,
		},
	}, nil
}

func saveCampaignRecord(StartCampaignRecord *StartCampaignRecord) error {
	// 获取带上下文和超时的DB实例
	db := config.GetDB().WithContext(context.Background())
	if err := db.Create(&model.CampaignModel{
		CampaignID: StartCampaignRecord.CampaignID.String(),
		Title:      StartCampaignRecord.Title,
		Goal:       StartCampaignRecord.Goal.String(),                   // *big.Int to string
		StartTime:  time.Unix(StartCampaignRecord.StartTime.Int64(), 0), // *big.Int to time.Time
		EndTime:    time.Unix(StartCampaignRecord.EndTime.Int64(), 0),   // *big.Int to time.Time
		Status:     StartCampaignRecord.Status,
		Starter:    StartCampaignRecord.Starter.Hex(), // common.Address to string
	}).Error; err != nil {
		return err
	}
	return nil
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

func parseCancelCampaignEvent(vLog types.Log) (*CampaignStatusRecord, error) {
	contractAbi := getContractABI()

	var event struct {
		Id     *big.Int
		Caller common.Address
		Time   *big.Int
	}

	if err := contractAbi.UnpackIntoInterface(&event, "CancellCampaign", vLog.Data); err != nil {
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
			EventType: CancelCampaignEvent,
			BlockTime: header.Time,
			TxHash:    vLog.TxHash,
		},
	}, nil
}

func parseCompleteCampaignEvent(vLog types.Log) (*CampaignStatusRecord, error) {
	contractAbi := getContractABI()

	var event struct {
		Id     *big.Int
		Caller common.Address
		Time   *big.Int
	}

	if err := contractAbi.UnpackIntoInterface(&event, "CompleteCampaign", vLog.Data); err != nil {
		return nil, fmt.Errorf("failed to unpack CompleteCampaign event: %v", err)
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
			EventType: CompleteCampaignEvent,
			BlockTime: header.Time,
			TxHash:    vLog.TxHash,
		},
	}, nil
}

func parseWithdrawEvent(vLog types.Log) (*CampaignIsWithdrawRecord, error) {
	contractAbi := getContractABI()

	var event struct {
		Id         *big.Int
		Withdrawer common.Address
		Time       *big.Int
	}

	if err := contractAbi.UnpackIntoInterface(&event, "Withdraw", vLog.Data); err != nil {
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
			return tx.Model(&existing).
				Updates(map[string]interface{}{
					"status": record.Status,
				}).Error
		}
	})
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

// 获取合约完整ABI
func getContractABI() abi.ABI {
	abiPath := config.Config.Geth.DonationContract.AbiPath // 从配置文件中读取路径
	abiBytes, err := ioutil.ReadFile(abiPath)
	if err != nil {
		log.Fatalf("Failed to read ABI file: %v", err)
	}
	contractAbi, err := abi.JSON(strings.NewReader(string(abiBytes)))
	if err != nil {
		log.Fatalf("Failed to parse contract ABI: %v", err)
	}
	return contractAbi
}

// 获取区块头信息
func getBlockHeader(blockHash common.Hash) (*types.Header, error) {
	return config.GethClient.HeaderByHash(context.Background(), blockHash)
}
