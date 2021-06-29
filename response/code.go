//go:generate stringer -type ErrCode -linecomment

package response

type ErrCode int

const (
	Success               ErrCode = 0     // 操作成功
	Error                 ErrCode = -1    // 操作失败
	RequestParamError     ErrCode = 10000 // 请求参数错误
	CaptchaGenerateFailed ErrCode = 10001 // 会员已存在

	MemberHasExist           ErrCode = 20001 // 会员创建失败
	MemberCreateFailed       ErrCode = 20002 //
	MemberAccountPasswordErr ErrCode = 20003 //
)
