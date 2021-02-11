package response

const (
	Success               = 0
	Error                 = -1
	RequestParamError     = 10000
	CaptchaGenerateFailed = 10001

	MemberHasExist           = 20001
	MemberCreateFailed       = 20002
	MemberAccountPasswordErr = 20003
)
