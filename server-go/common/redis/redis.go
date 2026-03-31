// Package redis 与Redis缓存数据库交互
// 处理邮箱验证码的存储和验证
package redis

import (
	"context"
	"strconv"
	"strings"
	"time"

	"github.com/go-redis/redis/v8"
	"github.com/yuanji6666/gopherAI/config"
)

// Rdb 全局redis客户端实例
var Rdb *redis.Client

var ctx = context.Background()

// InitRedis 初始化Redis
func InitRedis(){
	conf := config.GetConfig()
	host := conf.RedisHost
	port := conf.RedisPort
	password := conf.RedisPassword
	db := conf.RedisDB

	addr := host + ":" + strconv.Itoa(port)

	Rdb = redis.NewClient(&redis.Options{
		Addr: addr,
		Password: password,
		DB: db,
	})
}

// SetCaptchaForEmail 生成验证码并存Redis
func SetCaptchaForEmail(email, captcha string) error{
	key := GenerateCaptcha(email)
	expire := 2*time.Minute
	return  Rdb.Set(ctx, key, captcha, expire).Err()
}

// CheckCaptchaForEmail 检查验证码
func CheckCaptchaForEmail(email, userInputCaptcha string) (bool , error){
	key := GenerateCaptcha(email)

	storedCaptcha, err := Rdb.Get(ctx, key).Result()

	if err != nil {
		if err == redis.Nil {
			return false, nil
		}

		return false, nil
	}

	if strings.EqualFold(storedCaptcha, userInputCaptcha) {
		Rdb.Del(ctx, key)
		return true, nil
	}

	return false, nil
}



