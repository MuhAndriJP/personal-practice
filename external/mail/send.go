package mail

import (
	"context"

	"gopkg.in/gomail.v2"
)

type IGoMail interface {
	Send(context.Context, *Mail) error
}

type GoMail struct {
}

func NewGoMail() IGoMail {
	return GoMail{}
}

func (gm GoMail) Send(ctx context.Context, req *Mail) (err error) {
	gomailer := gomail.NewMessage()
	gomailer.SetHeader(FROM, SENDER_NAME)
	gomailer.SetHeader(TO, req.Recipient)
	gomailer.SetHeader(SUBJECT, SUBJECT_EMAIL)
	gomailer.SetBody(BODY_FORMAT, req.Body)
	// gomailer.Attach("/filename.jpg")

	auth := LoginAuth(CONFIG_USERNAME, CONFIG_PASSWORD)

	dialer := gomail.NewDialer(CONFIG_SMTP_HOST, CONFIG_SMTP_PORT, CONFIG_USERNAME, CONFIG_PASSWORD)
	dialer.Auth = auth

	return dialer.DialAndSend(gomailer)
}
