package captcha

import (
	"github.com/mojocn/base64Captcha"
)

func Get() (string, string, error) {
	return base64Captcha.NewCaptcha(base64Captcha.DefaultDriverDigit, base64Captcha.DefaultMemStore).Generate()
}

func Verify(captchaId, captchaVal string) bool {
	return base64Captcha.DefaultMemStore.Verify(captchaId, captchaVal, true)
}
