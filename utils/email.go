package utils

import (
	"fmt"
	"net/smtp"
	"strings"
)

/**
注册时 邮箱发送验证码
*/

// 服务器邮箱
var emailFrom = "2735268738@qq.com"

// 邮箱smtp验证码
var password = "bzoixykruslsdebc"

// host
var host = "smtp.qq.com"

func SendEmailVerification(emailTo string, verification string) (err error) {
	// 授权
	auth := smtp.PlainAuth("", emailFrom, password, host)
	to := []string{emailTo}
	user := emailFrom
	// 昵称
	nickname := "Code Playground"
	// 标题
	subject := "【Code Playground】"
	// 编码格式
	contentType := "Content-Type: text/plain; charset=UTF-8"

	// 内容
	body := fmt.Sprintf("【Code Playground】%s(账号注册验证码),请在30分钟内完成注册。如非本人操作，请忽略。\r\n.", verification)

	msg := []byte("To: " + strings.Join(to, ",") + "\r\nFrom: " + nickname +
		"<" + user + ">\r\nSubject: " + subject + "\r\n" + contentType + "\r\n\r\n" + body)
	err = smtp.SendMail("smtp.qq.com:25", auth, user, to, msg)
	return
}
