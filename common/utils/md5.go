package utils

import (
	"crypto/md5"
	"fmt"
)

//加密

func Md5(str string) string {
	// 判断验证码是否正确

	h := md5.New()
	h.Write([]byte(str))
	return fmt.Sprintf("%x", h.Sum(nil))
}

func AesDeCrypt(str string) string {

	return str
}
