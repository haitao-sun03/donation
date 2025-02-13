package utils

import (
	"bytes"
	"context"
	"crypto/ecdsa"
	"encoding/hex"
	"fmt"
	"io/ioutil"

	log "github.com/sirupsen/logrus"

	"github.com/ethereum/go-ethereum/accounts/keystore"
	"github.com/ethereum/go-ethereum/common"
	"github.com/ethereum/go-ethereum/core/types"
	"github.com/ethereum/go-ethereum/crypto"
	"github.com/ethereum/go-ethereum/rlp"
	"github.com/haitao-sun03/go/config"
)

func IsContract(addressStr string) bool {
	address := common.HexToAddress(addressStr)
	bytecode, err := config.GethClient.CodeAt(context.Background(), address, nil) // nil is lat
	if err != nil {
		log.Fatal(err)
	}
	return len(bytecode) > 0
}

func GetKeystorePK(keystorePath, password string) (*ecdsa.PrivateKey, error) {

	// 读取Keystore文件
	keyjson, err := ioutil.ReadFile(keystorePath)
	if err != nil {
		fmt.Println("Failed to read keystore file:", err)
		return nil, err
	}

	// 解析Keystore JSON数据
	key, err := keystore.DecryptKey(keyjson, password)
	if err != nil {
		log.WithError(err).Error("Failed to decrypt key")
		return nil, err
	}

	// 获取私钥的十六进制表示
	privateKeyBytes := key.PrivateKey.D.Bytes()
	privateKeyHex := fmt.Sprintf("%x", privateKeyBytes)
	// 通过私钥地址获取私钥
	privateKey, err := crypto.HexToECDSA(privateKeyHex)
	if err != nil {
		log.WithError(err).Error("get private key failed")
		return nil, err
	}
	return privateKey, nil
}

func Transection2HexStr(signedTx *types.Transaction) string {
	// 将交易编码为RLP字节数组
	rawTxBytes := new(bytes.Buffer)

	// 将signedTx的内容编码到rawTxBytes
	if err := rlp.Encode(rawTxBytes, &signedTx); err != nil {
		log.WithError(err).Error("signedTx to rawTxBytes failed")
		return ""
	}
	// 将字节数组转换为十六进制字符串
	rawTxHex := hex.EncodeToString(rawTxBytes.Bytes())
	return rawTxHex
}

func HexStr2Transection(rawTxHex string) *types.Transaction {
	// 将十六进制字符串解码为字节数组
	rawTxBytes, err := hex.DecodeString(rawTxHex)
	if err != nil {
		log.WithError(err).Error("DecodeString failed")
		return nil
	}

	// 使用RLP解码字节数组以恢复交易
	tx := new(types.Transaction)

	// 将rawTxBytes解码为types.Transaction，并写入tx指向的内存中
	if err := rlp.DecodeBytes(rawTxBytes, tx); err != nil {
		log.WithError(err).Error("rawTxBytes to tx failed")
		return nil
	}

	return tx
}
