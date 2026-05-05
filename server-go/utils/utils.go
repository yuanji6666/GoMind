// Package utils 工具函数包
package utils

import (
	"crypto/md5"
	"encoding/hex"
	"math/rand"
	"strconv"
	"time"
)

// GetRandomNumbers 调用rand标准库生成num位随机数
func GetRandomNumbers(num int) string {
	r := rand.New(rand.NewSource(time.Now().UnixNano()))

	code := ""
	for i := 0; i < num; i++ {
		// 0~9随机数
		digit := r.Intn(10)
		code += strconv.Itoa(digit)
	}
	return code
}

// MD5 返回src串的MD5加密
func MD5(src string) string {
	m := md5.New()
	m.Write([]byte(src))
	return hex.EncodeToString(m.Sum(nil))
}
