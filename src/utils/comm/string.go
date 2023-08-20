package comm

import (
	"math/rand"
	"time"
)

var baseStr = "0123456789aAbBcCdDeEfFgGhHiIjJkKlLmMnNoOpPqQrRsStTuUvVwWxXyYzZ"
var traceStr = "ABCDEFGHIJKLMNOPQRSTUVWXYZ"

// TraceID 生成8位随机日志ID
func TraceID() string {
	return RandStr(9, true)
}

// SaltKey 生成一个16位的随机盐值
func SaltKey() string {
	return RandStr(16, false)
}

// RandStr 生成指定位数的随机字符
func RandStr(length int, f bool) string {
	bytes := []byte(baseStr)
	if f {
		bytes = []byte(traceStr)
	}
	result := make([]byte, length)
	rand.Seed(time.Now().UnixNano() + int64(rand.Intn(10000)))
	for i := 0; i < length; i++ {
		result[i] = bytes[rand.Intn(len(bytes))]
	}
	return string(result)
}
