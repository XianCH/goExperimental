package response

import (
	"net/http"

	"github.com/gin-gonic/gin"
)

const (
	SUCCESS = 0
	ERROR   = 1
)

// Response 结构体表示 HTTP 响应的数据结构
type Response struct {
	Code int         `json:"code"`   // 响应状态码
	Msg  string      `json:"status"` // 响应状态描述
	Data interface{} `json:"data"`   // 响应数据
}

func ResultJson(ctx *gin.Context, code int, msg string, data any) {
	ctx.JSON(http.StatusOK, Response{
		Code: code,
		Msg:  msg,
		Data: data,
	})
}

func OK(ctx *gin.Context) {
	ResultJson(ctx, SUCCESS, "SUCCESS", nil)
}

func OkWithData(ctx *gin.Context, data any) {
	ResultJson(ctx, SUCCESS, "SUCCESS", data)
}

func OkWithMsg(ctx *gin.Context, msg string) {
	ResultJson(ctx, SUCCESS, msg, nil)
}

func Error(ctx *gin.Context) {
	ResultJson(ctx, ERROR, "Error", nil)
}

func ErrorWithData(ctx *gin.Context, data any) {
	ResultJson(ctx, ERROR, "ERROR", data)
}

func ErrorWithMsg(ctx *gin.Context, msg string) {
	ResultJson(ctx, ERROR, msg, nil)
}
