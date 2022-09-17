package response

const Invalid = -1
const OK = iota

type Result struct {
	Code    int    `json:"code"`
	Message string `json:"message"`
	Data    any    `json:"data"`
}
