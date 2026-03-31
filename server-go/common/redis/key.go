package redis

import (
	"fmt"

	"github.com/yuanji6666/gopherAI/config"
)

// GenerateCaptcha 生成带格式的验证码键
func GenerateCaptcha(email string) string{
	return fmt.Sprintf(config.DefaultRedisKeyConfig.CaptchaPrefix, email)
}