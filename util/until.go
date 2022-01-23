package util

import (
	"math/rand"
	"time"
)

/**
生成随机字符串
*/
func RandomString(num int) string {
	var letters = []byte("asdfghjklzxcvbnmqwertyuiopASDFGHJKLQWERTYUIOPZXCVBNM")
	result := make([]byte, num)
	rand.Seed(time.Now().Unix())
	for i := range result {
		result[i] = letters[rand.Intn(len(letters))]
	}
	return string(result)
}

