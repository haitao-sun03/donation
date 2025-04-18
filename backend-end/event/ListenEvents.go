package event

import (
	"context"
	"fmt"
	"io/ioutil"
	"math/big"
	"path/filepath"
	"runtime"
	"strings"
	"sync"
	"time"

	"github.com/ethereum/go-ethereum"
	"github.com/ethereum/go-ethereum/accounts/abi"
	"github.com/ethereum/go-ethereum/accounts/abi/bind"
	"github.com/ethereum/go-ethereum/common"
	"github.com/ethereum/go-ethereum/core/types"
	"github.com/ethereum/go-ethereum/crypto"
	"github.com/ethereum/go-ethereum/ethclient"
	_abi "github.com/haitao-sun03/donation/backend-end/abi"
	"github.com/haitao-sun03/donation/backend-end/config"
	"github.com/haitao-sun03/donation/backend-end/utils"
	"github.com/haitao-sun03/golang-distributelock/distribute"
	log "github.com/sirupsen/logrus"
)

// 定义confirmationMode
const (
	Soft string = "soft"
	Hard string = "hard"
)

type EventDispatcher struct {
	Events      chan *_abi.LineaSepoliaVerifyDataFinalizedV3 // 接收事件的通道
	Subscribers map[*big.Int][]chan common.Hash              // 按 L2 区块号存储订阅者通道
	Mu          sync.RWMutex                                 // 保护 subscribers 的读写锁
}

var EventDispatchers *EventDispatcher

func (ed *EventDispatcher) Start(ctx context.Context) {
	opts := &bind.WatchOpts{Context: ctx}

	sub, err := config.LineaSepoliaVerify.WatchDataFinalizedV3(opts, ed.Events, nil, nil, nil)
	if err != nil {
		log.Fatalf("WatchDataFinalizedV3 error: %v", err)
	}

	go func() {
		defer sub.Unsubscribe()
		for event := range ed.Events { // 使用结构体字段
			ed.dispatch(event)
		}
	}()
}

// 分发事件给订阅者（逻辑不变）
func (ed *EventDispatcher) dispatch(event *_abi.LineaSepoliaVerifyDataFinalizedV3) {
	ed.Mu.RLock()
	defer ed.Mu.RUnlock()
	log.Infoln("dispatch event,contains L2 blocks from ", event.StartBlockNumber.Uint64(), " to ", event.EndBlockNumber.Uint64())
	for l2BlockNum := event.StartBlockNumber; l2BlockNum.Cmp(event.EndBlockNumber) <= 0; l2BlockNum.Add(l2BlockNum, big.NewInt(1)) {
		if chans, ok := ed.Subscribers[l2BlockNum]; ok {
			log.Infof("dispatch l2BlockNum: %v L1 Verfify hash: %v", l2BlockNum.Uint64(), event.Raw.TxHash)
			for _, ch := range chans {
				ch <- event.Raw.TxHash
			}
			delete(ed.Subscribers, l2BlockNum)
		}
	}
}

// 注册订阅（逻辑不变）
func (ed *EventDispatcher) Subscribe(l2BlockNum *big.Int) <-chan common.Hash {
	ed.Mu.Lock()
	defer ed.Mu.Unlock()

	ch := make(chan common.Hash, 1)
	ed.Subscribers[l2BlockNum] = append(ed.Subscribers[l2BlockNum], ch)
	log.Infof("l2BlockNum %v subscribe L1 Verfify hash", l2BlockNum.Uint64())
	return ch
}

func ListenEvents(ctx context.Context) {
	DonationContractAddress := common.HexToAddress(config.Config.Geth.DonationContract.Address)
	NftContractAddress := common.HexToAddress(config.Config.Geth.NftContract.Address)
	switch config.Config.Geth.DeployNetwork {

	case "l2-linea":
		if config.Config.Geth.ConfimationMode == Hard {
			// 初始化EventDispatchers，并开始WatchDataFinalizedV3
			log.Infoln("ConfimationMode(hard),Init EventDispatchers and Start WatchDataFinalizedV3")
			EventDispatchers = &EventDispatcher{
				Events:      make(chan *_abi.LineaSepoliaVerifyDataFinalizedV3, 100),
				Subscribers: make(map[*big.Int][]chan common.Hash),
				Mu:          sync.RWMutex{},
			}
			EventDispatchers.Start(ctx)
		} else {
			log.Infoln("ConfimationMode(soft)")
		}

	}
	ListenAllContractEvents(ctx, config.GethWsClient, []common.Address{DonationContractAddress, NftContractAddress})
}

// 监听所有合约事件
func ListenAllContractEvents(ctx context.Context, client *ethclient.Client, contractAddrs []common.Address) {
	query := ethereum.FilterQuery{
		Addresses: contractAddrs,
	}

	// 未确认事件队列（BlockingQueue）
	unconfirmedQueue := make(chan types.Log, 1000)
	confirmedQueue := make(chan types.Log, 1000)

	sub, err := client.SubscribeFilterLogs(ctx, query, unconfirmedQueue)
	if err != nil {
		log.Fatalf("Failed to subscribe to logs: %v", err)
	}

	// 启动确认检查协程
	go checkBlockConfirmations(ctx, client, unconfirmedQueue, confirmedQueue)

	// 处理已确认事件
	go processConfirmedLogs(ctx, sub, unconfirmedQueue, confirmedQueue)

}

const CONFIRMED_BLOCK_NUM = 6

// 检查区块确认数
func checkBlockConfirmations(ctx context.Context, client *ethclient.Client, in chan types.Log, out chan<- types.Log) {
	for vLog := range in {
		switch config.Config.Geth.DeployNetwork {
		case "l1":
			if config.Config.Geth.ConfimationMode == Hard {
				// 获取当前最新区块号
				latestBlock, err := client.BlockNumber(ctx)
				if err != nil {
					log.Errorf("Failed to get block number: %v", err)
					continue
				}
				// 等待CONFIRMED_BLOCK个区块确认
				if latestBlock-vLog.BlockNumber >= CONFIRMED_BLOCK_NUM {
					log.Info("enter confirme")
					out <- vLog
				} else {
					// 未达到确认数，重新放回队列， 30秒后重试
					log.Info("enter not confirme")
					go func(log types.Log) {
						time.Sleep(30 * time.Second)
						in <- log
					}(vLog)
				}
			} else {
				out <- vLog
			}

		case "l2-linea":
			if config.Config.Geth.ConfimationMode == Hard {
				// 需等待L1验证 + CONFIRMED_BLOCK_NUM个区块确认
				if err := waitL1Confirmation(ctx, vLog); err != nil {
					log.Errorf("waitL1Confirmation failed: %v", err)
					continue
				}
			}

			out <- vLog
		}

	}
}

// 等待L1验证并确认
func waitL1Confirmation(ctx context.Context, l2Log types.Log) error {
	// 获取该L2事件在L1上的验证交易
	l1TxHash, err := getL1VerificationTx(l2Log.TxHash)
	if err != nil {
		log.Errorf("getL1VerificationTx error: %v", err)
		return err
	}

	// 监听L1的验证事件并等待确认
	receipt, _ := config.GethWsClientVerifier.TransactionReceipt(ctx, l1TxHash)
	l1VerfyTxBlockNum, err := utils.BigIntToUint64(receipt.BlockNumber)
	if err != nil {
		log.Errorf("BigIntToUint64 error: %v", err)
		return err
	}

	// 等待CONFIRMED_BLOCK_NUM个L1区块确认
	for {
		currentL1Block, _ := config.GethWsClientVerifier.BlockNumber(ctx)
		if currentL1Block-l1VerfyTxBlockNum >= CONFIRMED_BLOCK_NUM {
			return nil
		}
		time.Sleep(12 * time.Second) // 按平均出块时间轮询
	}
}

func getL1VerificationTx(l2TxHash common.Hash) (common.Hash, error) {
	receipt, err := config.GethWsClient.TransactionReceipt(context.Background(), l2TxHash)
	if err != nil {
		log.Errorf("GethWsClientVerifier.TransactionReceipt error: %v", err)
		return common.Hash{}, fmt.Errorf("GethWsClientVerifier.TransactionReceipt: %v", err)
	}
	l2BlockNum := receipt.BlockNumber
	l1HashChan := EventDispatchers.Subscribe(l2BlockNum)
	ticker := time.NewTicker(1 * time.Hour)
	defer ticker.Stop()
	// 每10minutes重试一次，直到收到l1TxHash
	retryCount := 0
	maxRetries := 10
	for retryCount < maxRetries {
		retryCount++
		select {
		case l1TxHash := <-l1HashChan:
			log.Infof("l2 %v(%v) receive l1Hash:%v", l2TxHash.Hex(), l2BlockNum.Uint64(), l1TxHash.Hex())
			return l1TxHash, nil
		case <-ticker.C:
			log.Infof("l2 %v(%v) not recieve l1Hash,enter the next %v 1 hour", l2TxHash.Hex(), l2BlockNum.Uint64(), retryCount)
			continue
		}
	}
	return common.Hash{}, fmt.Errorf("wait DataFinilizeV3 verify l1Hash over 10 hours")
}

// 处理已确认事件
func processConfirmedLogs(ctx context.Context, sub ethereum.Subscription, unconfirmedQueue chan types.Log, confirmedQueue <-chan types.Log) {
	// 工作池：限制并发处理协程数量
	const maxWorkers = 10
	workerChan := make(chan types.Log, maxWorkers)

	var wg sync.WaitGroup

	// 避免开启大量的goroutine
	for range maxWorkers {
		go func() {
			for vLog := range workerChan {
				handleLogWithLock(ctx, vLog)
				wg.Done()
			}
		}()
	}

	for {
		select {
		case err := <-sub.Err():
			log.Errorf("Subscription error: %v", err)
		case <-ctx.Done():
			sub.Unsubscribe()
			// 清空未确认事件通道残留
			for range unconfirmedQueue {
			}
			// 此时再关闭workerChan，避免workerChan关闭后，还将vLog放入workerChan导致panic
			close(workerChan)
			wg.Wait()
			log.Info("Event listener stopped")
			return
		case vLog := <-confirmedQueue:
			wg.Add(1)
			workerChan <- vLog
		}
	}
}

func handleLogWithLock(ctx context.Context, vLog types.Log) {
	eventKey := fmt.Sprintf("event:%s:%d", vLog.TxHash.Hex(), vLog.Index)
	lock := distribute.NewDistributedLock(config.RedisClient, eventKey, fmt.Sprintf("instance-%d", time.Now().UnixNano()), 10*time.Second)
	childCtx, cancel := context.WithCancel(ctx)
	defer cancel()
	acquired, err := lock.TryLock(childCtx)
	if err != nil || !acquired {
		log.Infof("Failed to acquire lock for event %s: %v", eventKey, err)
		return
	}

	log.Infof("acquire lock for event %s", eventKey)

	defer lock.Unlock(childCtx)

	// 每5s续期一次
	go func() {
		ticker := time.NewTicker(5 * time.Second)
		defer ticker.Stop()
		for {
			select {
			case <-childCtx.Done():
				return
			case <-ticker.C:
				retryCount := 0
				maxRetries := 3
				for retryCount < maxRetries {
					if err := lock.Renew(childCtx); err != nil {
						log.Errorf("Failed to renew lock for %s (retry %d/%d): %v", lock.Key, retryCount+1, maxRetries, err)
						retryCount++
						if retryCount == maxRetries {
							log.Errorf("Max retries reached, canceling renew operation")
							cancel()
							return
						}
						time.Sleep(1 * time.Second)
						continue
					}
					break
				}
			}
		}
	}()

	handleLog(ctx, vLog)
}

// 事件处理接口
type EventHandler interface {
	Parse(log types.Log) (interface{}, error)
	Handle(data interface{}) error
}

func handleLog(ctx context.Context, vLog types.Log) {
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

// 事件类型枚举
const (
	DonateEvent            = "Donate"
	StartCampaignEvent     = "StartCampaign"
	CancelledCampaignEvent = "CancellCampaign"
	CompletedCampaignEvent = "CompletedCampaign"
	ActiveCampaignEvent    = "ActiveCampaign"
	WithdrawEvent          = "Withdraw"
	RefundEvent            = "Refund"
	NftDeployedEvent       = "NFTDeployed"
)

var (
	eventHandlers map[common.Hash]EventHandler
	once          sync.Once
)

// 事件处理器注册表
func initHandlers() map[common.Hash]EventHandler {
	return map[common.Hash]EventHandler{
		crypto.Keccak256Hash([]byte("Donate(uint256,address,uint256)")):                                                               NewDonationHandler(nil),
		crypto.Keccak256Hash([]byte("StartCampaign(uint256,address,string,uint256,uint256,uint256,uint8,uint8,string,string,uint8)")): &CampaignHandler{},
		crypto.Keccak256Hash([]byte("CancellCampaign(uint256,address,uint256)")):                                                      &CampaignHandler{},
		crypto.Keccak256Hash([]byte("CompletedCampaign(uint256,address,uint256)")):                                                    &CampaignHandler{},
		crypto.Keccak256Hash([]byte("Withdraw(uint256,address,uint256,uint256)")):                                                     &CampaignHandler{},
		crypto.Keccak256Hash([]byte("Refund(uint256,address,uint256,uint256)")):                                                       NewDonationHandler(nil),
		crypto.Keccak256Hash([]byte("NFTDeployed(tuple[])")):                                                                          &NFTHandler{},
		crypto.Keccak256Hash([]byte("NFTDeployed((string,string)[])")):                                                                &NFTHandler{},
		crypto.Keccak256Hash([]byte("ActiveCampaign(uint256,address,uint256)")):                                                       &CampaignHandler{},
	}
}

// GetEventHandlers 返回事件处理器注册表（单例）
func GetEventHandlers() map[common.Hash]EventHandler {
	once.Do(func() {
		config.Init(false) // 确保数据库初始化
		eventHandlers = initHandlers()
	})
	return eventHandlers
}

// 基础事件结构
type BaseEvent struct {
	EventType string
	BlockTime uint64
	TxHash    common.Hash
}

const (
	CampaignRelContract = "CampaignRelContract"
	NftContract         = "NftContract"
)

// 获取合约完整ABI
func getContractABI(contract string) abi.ABI {
	_, filename, _, _ := runtime.Caller(0)
	projectRoot := filepath.Dir(filepath.Dir(filename))
	var abiPath string
	if contract == "CampaignRelContract" {
		abiPath = config.Config.Geth.DonationContract.AbiPath // 从配置文件中读取路径
	} else {
		abiPath = config.Config.Geth.NftContract.AbiPath
	}

	abiBytes, err := ioutil.ReadFile(filepath.Join(projectRoot, abiPath))
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
