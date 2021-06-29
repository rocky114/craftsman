package response

var message = map[ErrCode]string{
	Success:                  "操作成功",
	Error:                    "操作失败",
	RequestParamError:        "请求参数错误",
	MemberHasExist:           "会员已存在",
	MemberCreateFailed:       "会员创建失败",
	MemberAccountPasswordErr: "账号密码错误",
	CaptchaGenerateFailed:    "验证码生成失败",
}
