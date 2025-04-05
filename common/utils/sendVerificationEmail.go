package utils

import (
	"net/smtp"
	"os"
)

func SendEmail(email string) error {
	from := os.Getenv("SMTP_FROM")
	password := os.Getenv("SMTP_PASSWORD")
	smtpHost := os.Getenv("SMTP_HOST")
	smtpPort := os.Getenv("SMTP_PORT")

	// 构建邮件内容
	to := []string{email}
	subject := "邮箱注册验证"
	body := "感谢您注册，请点击以下链接完成验证：[验证链接]"
	message := []byte("To: " + email + "\r\n" +
		"From: " + from + "\r\n" +
		"Subject: " + subject + "\r\n" +
		"\r\n" +
		body)

	// 认证
	auth := smtp.PlainAuth("", from, password, smtpHost)

	// 发送邮件
	err := smtp.SendMail(smtpHost+":"+smtpPort, auth, from, to, message)
	return err
}
