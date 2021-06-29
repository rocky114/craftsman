//go:generate stringer -type ErrCode -linecomment

package response

type ErrCode int

const (
	Success               ErrCode = 0
	Error                 ErrCode = -1
	RequestParamError     ErrCode = 10000
	CaptchaGenerateFailed ErrCode = 10001

	MemberHasExist           ErrCode = 20001
	MemberCreateFailed       ErrCode = 20002
	MemberAccountPasswordErr ErrCode = 20003
)
