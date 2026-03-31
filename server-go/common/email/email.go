//Package email 调用gomail库发送邮件
package email

import (
	"fmt"

	"github.com/yuanji6666/gopherAI/config"
	"gopkg.in/gomail.v2"
)

const (
	CodeMsg     = "GopherAI验证码如下(验证码仅限于2分钟有效): "
	UserNameMsg = "GopherAI的账号如下，请保留好，后续可以用账号/邮箱登录 "
)

//SendCaptcha 以“msg code”为正文向email发送邮件
//使用SMTP（simple mail transfer protocol）协议
//使用SMTP标准587端口（支持TLS加密）
func SendCaptcha(email, code, msg string) error {
	m := gomail.NewMessage()

	m.SetHeader("From", config.GetConfig().Email)
	m.SetHeader("To", email)
	m.SetHeader("Subject", "来自GopherAI的消息")
	m.SetBody("text/plain", msg+" "+code)

	d := gomail.NewDialer("smtp.qq.com", 587, config.GetConfig().Email, config.GetConfig().Authcode)

	if err := d.DialAndSend(m); err != nil {
		fmt.Printf("Dial and send Err : %s", err)
		return err
	}

	fmt.Println("send mail success!")
	return nil
}
