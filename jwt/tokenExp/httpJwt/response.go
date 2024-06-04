package httpjwt

import (
	"encoding/json"
	"net/http"

	"github.com/gin-gonic/gin"
)

const (
	SuccessCode ResponseCode = 200
	ErrorCode   ResponseCode = 400
)

type ResponseCode int

type ResponseEntity struct {
	Code ResponseCode `json:"code"`
	Msg  string       `json:"Msg"`
	Data interface{}  `json:"Data"`
}

func ResponseError(c *gin.Context, code ResponseCode, err error) {
	rep := &ResponseEntity{
		Code: ErrorCode,
		Data: "",
		Msg:  err.Error(),
	}
	c.JSON(http.StatusOK, rep)
	response, _ := json.Marshal(rep)
	c.Set("response", string(response))
}

func ResponseSuccess(c *gin.Context, data interface{}) {
	rep := &ResponseEntity{
		Code: SuccessCode,
		Data: data,
		Msg:  "success!",
	}
	c.JSON(http.StatusOK, rep)
	response, _ := json.Marshal(rep)

	c.Set("response", string(response))
}
