package notify

import (
	"bytes"
	"fmt"
	"github.com/lstack-org/go-web-framework/pkg/notify/templates"
	"html/template"
	"strings"
	"time"

	"gopkg.in/gomail.v2"
)

var (
	mOpts mailOpts
	_     Interface = &mail{}
)

func InitMailNotify(mailHost, mailUser, mailPass, mailTo string, mailPort int) {
	mOpts.mailHost = mailHost
	mOpts.mailPort = mailPort
	mOpts.mailUser = mailUser
	mOpts.mailPass = mailPass
	mOpts.mailTo = mailTo
}

func NewMail(subject string) Interface {
	return &mail{subject: subject, mailOpts: mOpts}
}

type mailOpts struct {
	mailHost string
	mailPort int
	mailUser string
	mailPass string
	mailTo   string
}

type mail struct {
	mailOpts
	subject string
}

func (m *mail) Validate() bool {
	if m.mailHost == "" {
		return false
	}

	if m.mailPort == 0 {
		return false
	}

	if m.mailUser == "" {
		return false
	}

	if m.mailPass == "" {
		return false
	}

	if m.mailTo == "" {
		return false
	}

	return true
}

func (m *mail) Send(body string) error {
	message := gomail.NewMessage()

	//设置发件人
	message.SetHeader("From", m.mailUser)

	//设置发送给多个用户
	mailArrTo := strings.Split(m.mailTo, ",")
	message.SetHeader("To", mailArrTo...)

	//设置邮件主题
	message.SetHeader("Subject", m.subject)

	//设置邮件正文
	message.SetBody("text/html", body)

	d := gomail.NewDialer(m.mailHost, m.mailPort, m.mailUser, m.mailPass)

	return d.DialAndSend(message)
}

// NewPanicHTMLEmail 发送系统异常邮件 html
func NewPanicHTMLEmail(method, host, uri, stack string) (subject string, body string, err error) {
	mailData := &struct {
		URL   string
		Stack string
		Year  int
	}{
		URL:   fmt.Sprintf("%s %s%s", method, host, uri),
		Stack: stack,
		Year:  time.Now().Year(),
	}

	mailTplContent, err := getEmailHTMLContent(templates.PanicMail, mailData)
	return fmt.Sprintf("[系统异常]-%s", uri), mailTplContent, err
}

// getEmailHTMLContent 获取邮件模板
func getEmailHTMLContent(mailTpl string, mailData interface{}) (string, error) {
	tpl, err := template.New("email tpl").Parse(mailTpl)
	if err != nil {
		return "", err
	}
	buffer := new(bytes.Buffer)
	err = tpl.Execute(buffer, mailData)
	if err != nil {
		return "", err
	}
	return buffer.String(), nil
}
