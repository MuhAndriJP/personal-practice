package mail

import (
	"context"
	"os"

	"gopkg.in/gomail.v2"
)

type IGoMail interface {
	Send(context.Context, *MailRequest) error
}

type GoMail struct {
}

func NewGoMail() IGoMail {
	return GoMail{}
}

func (gm GoMail) Send(ctx context.Context, req *MailRequest) (err error) {
	gomailer := gomail.NewMessage()
	gomailer.SetHeader(FROM, req.SenderName+" <"+os.Getenv("MAIL_SENDER")+">")
	gomailer.SetHeader(TO, req.To)
	gomailer.SetHeader(SUBJECT, req.Subject)
	gomailer.SetBody(BODY_FORMAT, req.Body)
	if req.Attach != "" {
		gomailer.Attach(req.Attach)
	}

	dialer := gomail.NewDialer(CONFIG_SMTP_HOST, CONFIG_SMTP_PORT, CONFIG_EMAIL, CONFIG_PASSWORD)

	return dialer.DialAndSend(gomailer)
}
