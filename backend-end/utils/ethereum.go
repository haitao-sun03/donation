package utils

import (
	"context"
	"crypto/rand"
	"encoding/base64"
	"errors"
	"fmt"
	"math/big"
	"strings"

	"github.com/ethereum/go-ethereum"
	"github.com/ethereum/go-ethereum/accounts/abi/bind"
	"github.com/ethereum/go-ethereum/common"
	"github.com/ethereum/go-ethereum/crypto"
	"github.com/haitao-sun03/donation/backend-end/config"
	log "github.com/sirupsen/logrus"
)

// GetTransactOpts 返回配置好的TransactOpts，gasLimit为0时使用默认值
func GetTransactOpts(gasLimit uint64, to common.Address, value *big.Int, data []byte) (*bind.TransactOpts, error) {
	privateKeyStr := config.Config.Geth.Nft.PrivateKey
	log.Debugf("Private key length: %d", len(privateKeyStr))
	privateKey, err := crypto.HexToECDSA(strings.TrimSpace(privateKeyStr))
	if err != nil {
		return nil, fmt.Errorf("failed to parse private key: %v, key length: %d", err, len(privateKeyStr))
	}

	chainID := big.NewInt(config.Config.Geth.Nft.ChainID)
	auth, err := bind.NewKeyedTransactorWithChainID(privateKey, chainID)
	if err != nil {
		return nil, err
	}

	callMsg := ethereum.CallMsg{
		From:  auth.From,
		To:    &to,
		Value: value,
		Data:  data,
	}
	estimatedGas, err := config.GethClient.EstimateGas(context.Background(), callMsg)
	if err != nil {
		log.Warnf("EstimateGas failed: %v, using default gasLimit: 300000", err)
		estimatedGas = 300000 // 兜底默认值
	}

	// 添加 20% 余量
	gasLimitWithMargin := uint64(float64(estimatedGas) * 1.2)
	if gasLimit > 0 {
		auth.GasLimit = gasLimit // 优先使用传入值
	} else {
		auth.GasLimit = gasLimitWithMargin
	}

	// 获取当前区块
	header, err := config.GethClient.HeaderByNumber(context.Background(), nil)
	if err != nil {
		return nil, err
	}

	// 设置 EIP-1559 费用
	tip, err := config.GethClient.SuggestGasTipCap(context.Background())
	if err != nil {
		return nil, err
	}

	baseFee := header.BaseFee
	if baseFee == nil {
		baseFee = big.NewInt(0)
	}

	maxPriorityFeePerGas := new(big.Int).Set(tip)
	maxFeePerGas := new(big.Int).Add(
		maxPriorityFeePerGas,
		new(big.Int).Mul(baseFee, big.NewInt(2)), // 通常设置为基础费用的2倍
	)

	auth.GasFeeCap = maxFeePerGas
	auth.GasTipCap = maxPriorityFeePerGas

	// 设置交易超时
	auth.Context = context.Background()
	auth.NoSend = false // 允许直接发送交易

	return auth, nil
}

func GenerateNonce() (string, error) {
	nonceBytes := make([]byte, 16)  // 生成 16 字节随机数
	_, err := rand.Read(nonceBytes) // 填充随机字节
	if err != nil {
		return "", err
	}
	nonceString := base64.URLEncoding.EncodeToString(nonceBytes) // 转为字符串
	return nonceString, nil
}

// 签名验证工具函数
func VerifySignature(address, signatureHex, nonce string) (bool, error) {
	// 构造预期消息
	signParam := "please login,confirm the signature that do not cost gas\n nonce: " + nonce
	message := fmt.Sprintf("\x19Ethereum Signed Message:\n%d%s", len(signParam), signParam)
	messageHash := crypto.Keccak256Hash([]byte(message))

	// 解码签名
	signature := common.Hex2Bytes(signatureHex[2:]) // 去掉0x前缀
	if len(signature) != 65 {
		return false, errors.New("signature length is not 65")
	}

	// 恢复公钥
	signature[64] -= 27 // 调整恢复标识位
	pubKey, err := crypto.SigToPub(messageHash.Bytes(), signature)
	if err != nil {
		return false, err
	}

	// 验证地址匹配
	recoveredAddr := crypto.PubkeyToAddress(*pubKey)
	return common.HexToAddress(address) == recoveredAddr, nil
}
