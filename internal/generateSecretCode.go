package internal

import (
	"fmt"
	"math/rand"
	"time"
)

var letterBytes = "0123456789abcdefghijklmnopqrstuvwxyzABCDEFGHIJKLMNOPQRSTUVWXYZ"

const (
	letterIdxBits = 10
	letterIdxMask = 1<<letterIdxBits - 1
	letterIdxMax  = 63 / letterIdxBits
)

func init() {
	rand.Seed(time.Now().UnixNano())
}

// 生成随机密码
func GenerateSecretCode(specialCharacter string, length int) {
	special_char := "~`!@#$%^&*()_+=-,.?/|{}[]"
	if specialCharacter == "" {
		letterBytes += special_char
	} else if specialCharacter == "false" {

	} else {
		letterBytes += specialCharacter
	}
	randomCode := RandStringBytesMaskImpr(length)
	fmt.Printf("生成的的随机密码为:\n\n%s\n\n请您妥善保管，祝您工作愉快^o^\n", randomCode)
}

func RandStringBytesMaskImpr(n int) string {
	b := make([]byte, n)
	for i, cache, remain := n-1, rand.Int63(), letterIdxMax; i >= 0; {
		if remain == 0 {
			cache, remain = rand.Int63(), letterIdxMax
		}
		if idx := int(cache & letterIdxMask); idx < len(letterBytes) {
			b[i] = letterBytes[idx]
			i--
		}
		cache >>= letterIdxBits
		remain--
	}

	return string(b)
}
