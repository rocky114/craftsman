package service

type Result struct {
	Code int         `json:"code"`
	Msg  string      `json:"msg"`
	Data interface{} `json:"data"`
}

func Success(data interface{}, msg ...string) *Result {
	var message string
	if len(msg) == 1 {
		message = msg[0]
	} else {
		message = "successfully"
	}

	return &Result{
		Code: 0,
		Msg:  message,
		Data: data,
	}
}

func Error(msg string, data ...interface{}) *Result {
	if len(data) == 0 {
		data = []interface{}{}
	}

	return &Result{
		Code: -1,
		Msg:  msg,
		Data: data,
	}
}
