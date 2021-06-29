package response

import (
	"github.com/gin-gonic/gin"
)

type GinContext struct {
	C *gin.Context
}

type Result struct {
	Code ErrCode     `json:"code"`
	Msg  interface{} `json:"msg"`
	Data interface{} `json:"data"`
}

func (g *GinContext) Response(httpCode int, errCode ErrCode, data interface{}) {
	if data == nil {
		data = []struct{}{}
	}

	g.C.JSON(httpCode, Result{
		Code: errCode,
		Msg:  GetMsg(errCode),
		Data: data,
	})
}

func GetMsg(code ErrCode) string {
	if msg, ok := message[code]; ok {
		return msg
	}

	return message[Success]
}
