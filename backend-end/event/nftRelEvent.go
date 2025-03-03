package event

import (
	"context"
	"fmt"

	"github.com/ethereum/go-ethereum/core/types"
	"github.com/ethereum/go-ethereum/crypto"
	"github.com/haitao-sun03/donation/backend-end/config"
	"github.com/haitao-sun03/donation/backend-end/model"
)

// NftHandler 处理Nft相关
type NFTHandler struct {
	EventHandler
}

// NftDeployedRecord 处理Nft部署事件
type NftDeployedRecord struct {
	LevelMetaDataURI []LevelMetaDataURI
	BaseEvent
}

// 定义一个结构体来存储键值对
type LevelMetaDataURI struct {
	Level string `abi:"level"`
	Uri   string `abi:"uri"`
}

func (nft *NFTHandler) Parse(vLog types.Log) (interface{}, error) {
	switch vLog.Topics[0] {
	case crypto.Keccak256Hash([]byte("NFTDeployed(tuple[])")):
		return parseNftDeployedEvent(vLog)
	case crypto.Keccak256Hash([]byte("NFTDeployed((string,string)[])")):
		return parseNftDeployedEvent(vLog)
	case crypto.Keccak256Hash([]byte("Refund(uint256,address,uint256,uint256)")):
		return parseRefundEvent(vLog)
	default:
		return nil, fmt.Errorf("unsupported event type")
	}
}

func (nft *NFTHandler) Handle(data interface{}) error {

	switch event := data.(type) {
	case *NftDeployedRecord:
		return saveNftMeta(event)
	case *RefundRecord:
		return updateDonationRefund(event)

	default:
		return fmt.Errorf("invalid donation event data")
	}

}

func parseNftDeployedEvent(vLog types.Log) (*NftDeployedRecord, error) {
	contractAbi := getContractABI(NftContract)

	var event struct {
		LevelMetaDataURI []LevelMetaDataURI
	}

	if err := contractAbi.UnpackIntoInterface(&event, NftDeployedEvent, vLog.Data); err != nil {
		return nil, fmt.Errorf("failed to unpack NftDeployed event: %v", err)
	}

	header, err := getBlockHeader(vLog.BlockHash)
	if err != nil {
		return nil, err
	}

	return &NftDeployedRecord{
		LevelMetaDataURI: event.LevelMetaDataURI,
		BaseEvent: BaseEvent{
			EventType: NftDeployedEvent,
			BlockTime: header.Time,
			TxHash:    vLog.TxHash,
		},
	}, nil
}

func saveNftMeta(nftDeployedRecord *NftDeployedRecord) error {
	db := config.GetDB().WithContext(context.Background())
	for _, nftMetadata := range nftDeployedRecord.LevelMetaDataURI {
		err := db.Create(&model.NFTMetaDataModel{
			Level: nftMetadata.Level,
			URI:   nftMetadata.Uri,
		}).Error
		if err != nil {
			return err
		}
	}
	return nil
}
