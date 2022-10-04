package response

const (
	successed = 0
	failed    = -1
)

type Result struct {
	Code int    `json:"code"`
	Msg  string `json:"msg"`
	Data any    `json:"data"`
}

func NewSuccessed(data interface{}) *Result {
	return &Result{Code: successed, Msg: "ok", Data: data}
}

func NewFailed(msg string) *Result {
	return &Result{Code: failed, Msg: msg, Data: []struct{}{}}
}
