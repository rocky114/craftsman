package response

const (
	success = 0
	fail    = -1
)

type Result struct {
	Code int    `json:"code"`
	Msg  string `json:"msg"`
	Data any    `json:"data"`
}

func NewSuccess(data interface{}) *Result {
	return &Result{Code: success, Msg: "ok", Data: data}
}

func NewFail(msg string) *Result {
	return &Result{Code: fail, Msg: msg, Data: []struct{}{}}
}
