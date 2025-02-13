package common

import (
	"context"
	"crypto/ecdsa"
	"math/big"
	"net/http"

	"github.com/ethereum/go-ethereum/accounts/abi/bind"
	"github.com/ethereum/go-ethereum/crypto"
	"github.com/ethereum/go-ethereum/ethclient"
	"github.com/gin-gonic/gin"
	"github.com/haitao-sun03/go/config"
	log "github.com/sirupsen/logrus"
	"gorm.io/gorm"
)

const (
	MessageSuccess = "success"
	MessageFailure = "failure"
	// 你可以定义更多的消息常量
)

type Result struct {
	Code  int   `json:"code"`
	Msg   any   `json:"msg"`
	Data  any   `json:"data"`
	Count int64 `json:"count"`
}

func Success(ctx *gin.Context, code int, msg any, data any, count int64) {
	res := Result{
		Code:  code,
		Msg:   msg,
		Data:  data,
		Count: count,
	}
	ctx.JSON(http.StatusOK, res)
}

func Fail(ctx *gin.Context, code int, msg any) {
	res := Result{
		Code: code,
		Msg:  msg,
	}
	ctx.JSON(http.StatusOK, res)
}

func PaginationFunc(db *gorm.DB, pageNum int, pageSize int) {

	// 计算偏移量
	offset := (pageNum - 1) * pageSize
	db.Offset(offset).Limit(pageSize)
}

// func InitAccountPathInController() {
// 	AccountPath = fmt.Sprint(config.Config.Geth.KeystorePath, "/UTC--2024-11-28T03-26-22.871269100Z--6c0db8c49190b517b949429b9dea1c2b32143bd2")
// }

func ConstructTransactionOpts(gethClient *ethclient.Client, privateKey *ecdsa.PrivateKey, ctx *gin.Context, gasLimit, gasTipCap, GasFeeCap int64) (*bind.TransactOpts, bool) {
	// 获取链ID
	chainID, err := gethClient.ChainID(context.Background())
	if err != nil {
		log.WithError(err).Error("get chain id error")
	}
	opts, err := bind.NewKeyedTransactorWithChainID(privateKey, chainID)
	if err != nil {
		log.WithError(err).Error("NewKeyedTransactorWithChainID error")
		Fail(ctx, http.StatusInternalServerError, err)
		return nil, true
	}
	publicKey := privateKey.PublicKey
	address := crypto.PubkeyToAddress(publicKey)
	nonce, err := config.GethClient.PendingNonceAt(context.Background(), address)
	if err != nil {
		log.WithError(err).Error("get from account nonce fail")
		Fail(ctx, http.StatusInternalServerError, err)
		return nil, true
	}
	opts.Nonce = new(big.Int).SetUint64(nonce)

	opts.GasLimit = uint64(gasLimit)
	opts.GasTipCap = big.NewInt(gasTipCap)
	opts.GasFeeCap = big.NewInt(GasFeeCap)
	return opts, false
}
