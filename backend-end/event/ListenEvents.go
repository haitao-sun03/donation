package event

import (
	"context"
	"fmt"
	"io/ioutil"
	"path/filepath"
	"runtime"
	"strings"
	"sync"
	"time"

	"github.com/ethereum/go-ethereum"
	"github.com/ethereum/go-ethereum/accounts/abi"
	"github.com/ethereum/go-ethereum/common"
	"github.com/ethereum/go-ethereum/core/types"
	"github.com/ethereum/go-ethereum/crypto"
	"github.com/ethereum/go-ethereum/ethclient"
	"github.com/haitao-sun03/donation/backend-end/config"
	"github.com/haitao-sun03/golang-distributelock/distribute"
	log "github.com/sirupsen/logrus"
)

func ListenEvents(ctx context.Context) {
	DonationContractAddress := common.HexToAddress(config.Config.Geth.DonationContract.Address)
	NftContractAddress := common.HexToAddress(config.Config.Geth.NftContract.Address)
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
	go processConfirmedLogs(ctx, sub, confirmedQueue)

}

const CONFIRMED_BLOCK_NUM = 12

// 检查区块确认数
func checkBlockConfirmations(ctx context.Context, client *ethclient.Client, in chan types.Log, out chan<- types.Log) {
	// 获取当前最新区块号
	latestBlock, _ := client.BlockNumber(ctx)
	for vLog := range in {
		// 等待CONFIRMED_BLOCK个区块确认
		if latestBlock-vLog.BlockNumber >= CONFIRMED_BLOCK_NUM {
			out <- vLog
		} else {
			// 未达到确认数，重新放回队列（需实现延迟重试逻辑）
			go func(log types.Log) {
				time.Sleep(30 * time.Second) // 30秒后重试
				in <- log
			}(vLog)
		}
	}
}

// 处理已确认事件
func processConfirmedLogs(ctx context.Context, sub ethereum.Subscription, confirmedQueue <-chan types.Log) {
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
			// 清空残留数据
			for range confirmedQueue {
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
