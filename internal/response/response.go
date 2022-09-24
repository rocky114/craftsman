package response

const Invalid = -1
const OK = iota

const (
	ParameterErr = "参数错误"
	UnknownErr   = "服务出错了"
)

type Result struct {
	Code    int    `json:"code"`
	Message string `json:"message"`
	Data    any    `json:"data"`
}
