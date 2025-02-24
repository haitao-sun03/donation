package event

import (
	"context"
	"io/ioutil"
	"strings"

	"github.com/ethereum/go-ethereum"
	"github.com/ethereum/go-ethereum/accounts/abi"
	"github.com/ethereum/go-ethereum/common"
	"github.com/ethereum/go-ethereum/core/types"
	"github.com/ethereum/go-ethereum/crypto"
	"github.com/ethereum/go-ethereum/ethclient"
	"github.com/haitao-sun03/go/config"
	log "github.com/sirupsen/logrus"
)

func ListenEvents() {
	DonationContractAddress := common.HexToAddress(config.Config.Geth.DonationContract.Address)
	NftContractAddress := common.HexToAddress(config.Config.Geth.NftContract.Address)
	ListenAllContractEvents(context.Background(), config.GethWsClient, []common.Address{DonationContractAddress, NftContractAddress})
}

// 监听所有合约事件
func ListenAllContractEvents(ctx context.Context, client *ethclient.Client, contractAddrs []common.Address) {
	query := ethereum.FilterQuery{
		Addresses: contractAddrs,
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

// 事件处理接口
type EventHandler interface {
	Parse(log types.Log) (interface{}, error)
	Handle(data interface{}) error
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

// 事件处理器注册表
var eventHandlers = map[common.Hash]EventHandler{
	crypto.Keccak256Hash([]byte("Donate(uint256,address,uint256)")):                                                               &DonationHandler{},
	crypto.Keccak256Hash([]byte("StartCampaign(uint256,address,string,uint256,uint256,uint256,uint8,uint8,string,string,uint8)")): &CampaignHandler{},
	crypto.Keccak256Hash([]byte("CancellCampaign(uint256,address,uint256)")):                                                      &CampaignHandler{},
	crypto.Keccak256Hash([]byte("CompletedCampaign(uint256,address,uint256)")):                                                    &CampaignHandler{},
	crypto.Keccak256Hash([]byte("Withdraw(uint256,address,uint256,uint256)")):                                                     &CampaignHandler{},
	crypto.Keccak256Hash([]byte("Refund(uint256,address,uint256,uint256)")):                                                       &DonationHandler{},
	crypto.Keccak256Hash([]byte("NFTDeployed(tuple[])")):                                                                          &NFTHandler{},
	crypto.Keccak256Hash([]byte("NFTDeployed((string,string)[])")):                                                                &NFTHandler{},
	crypto.Keccak256Hash([]byte("ActiveCampaign(uint256,address,uint256)")):                                                       &CampaignHandler{},
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
	var abiPath string
	if contract == "CampaignRelContract" {
		abiPath = config.Config.Geth.DonationContract.AbiPath // 从配置文件中读取路径
	} else {
		abiPath = config.Config.Geth.NftContract.AbiPath
	}

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
