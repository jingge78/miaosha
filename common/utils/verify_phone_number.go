package utils

import "regexp"

// ValidatePhoneNumber 验证手机号是否有效
func ValidatePhoneNumber(phone string) bool {
	pattern := `^1[3-9]\d{9}$`
	match, _ := regexp.MatchString(pattern, phone)
	return match
}
