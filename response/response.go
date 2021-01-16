package response

import (
	"github.com/gin-gonic/gin"
)

type GinContext struct {
	C *gin.Context
}

type Result struct {
	Code int         `json:"code"`
	Msg  interface{} `json:"msg"`
	Data interface{} `json:"data"`
}

func (g *GinContext) Response(httpCode, errCode int, data interface{}) {
	g.C.JSON(httpCode, Result{
		Code: errCode,
		Msg:  GetMsg(errCode),
		Data: data,
	})
}

func GetMsg(code int) string {
	if msg, ok := message[code]; ok {
		return msg
	}

	return message[Success]
}
