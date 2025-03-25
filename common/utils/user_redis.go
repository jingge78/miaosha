package utils

import (
	"miaosha-jjl/common/global"
	"time"
)

func SendSms(mobile, code string) error {
	return global.Rdb.Set(global.CTX, "sendSms"+mobile, code, time.Minute*5).Err()
}

func PasswordRecoverySendSms(mobile string) string {
	return global.Rdb.Get(global.CTX, "sendSms"+mobile).Val()
}
