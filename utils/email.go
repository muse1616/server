package utils

import (
	"crypto/tls"
	"fmt"
	"log"
	"net"
	"net/smtp"
	"strings"
)

/**
此处邮箱发送全部使用smtp lTLS的465端口 注意云服务器25端口默认禁用 使用smtp的25端口会导致邮箱服务无法使用
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
	// linux发送
	err = SendMailUsingTLS(fmt.Sprintf("%s:%d", host, 465), auth, emailFrom, []string{emailTo}, msg)
	//err = smtp.SendMail("smtp.qq.com:25", auth, user, to, msg)
	return
}

func SendEmailVerificationForLogin(emailTo string, verification string) (err error) {
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
	body := fmt.Sprintf("【Code Playground】%s(账号登录提醒),请在5分钟内使用验证码登录。如非本人操作，请忽略。\r\n.", verification)

	msg := []byte("To: " + strings.Join(to, ",") + "\r\nFrom: " + nickname +
		"<" + user + ">\r\nSubject: " + subject + "\r\n" + contentType + "\r\n\r\n" + body)
	err = SendMailUsingTLS(fmt.Sprintf("%s:%d", host, 465), auth, emailFrom, []string{emailTo}, msg)
	//err = smtp.SendMail("smtp.qq.com:465", auth, user, to, msg)
	return
}

//return a smtp client
func Dial(addr string) (*smtp.Client, error) {
	conn, err := tls.Dial("tcp", addr, nil)
	if err != nil {
		log.Panicln("Dialing Error:", err)
		return nil, err
	}
	//分解主机端口字符串
	host, _, _ := net.SplitHostPort(addr)
	return smtp.NewClient(conn, host)
}

//参考net/smtp的func SendMail()
//使用net.Dial连接tls(ssl)端口时,smtp.NewClient()会卡住且不提示err
//len(to)>1时,to[1]开始提示是密送
func SendMailUsingTLS(addr string, auth smtp.Auth, from string,
	to []string, msg []byte) (err error) {

	//create smtp client
	c, err := Dial(addr)
	if err != nil {
		log.Println("Create smpt client error:", err)
		return err
	}
	defer c.Close()

	if auth != nil {
		if ok, _ := c.Extension("AUTH"); ok {
			if err = c.Auth(auth); err != nil {
				log.Println("Error during AUTH", err)
				return err
			}
		}
	}

	if err = c.Mail(from); err != nil {
		return err
	}

	for _, addr := range to {
		if err = c.Rcpt(addr); err != nil {
			return err
		}
	}

	w, err := c.Data()
	if err != nil {
		return err
	}

	_, err = w.Write(msg)
	if err != nil {
		return err
	}

	err = w.Close()
	if err != nil {
		return err
	}
	return c.Quit()
}
