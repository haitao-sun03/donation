package common

import (
	"net/http"

	"github.com/gin-gonic/gin"
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
