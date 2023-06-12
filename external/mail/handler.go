package mail

import (
	"context"
	"log"

	"github.com/MuhAndriJP/personal-practice.git/helper"
	"github.com/labstack/echo/v4"
)

type Mail struct{}

func (m *Mail) Handle(c echo.Context) error {
	ctx := c.Request().Context()
	if ctx == nil {
		ctx = context.Background()
	}

	file, err := c.FormFile("attach")
	if err != nil {
		return err
	}

	r := new(MailRequest)
	if err := c.Bind(r); err != nil {
		return err
	}
	r.SenderName = c.FormValue("sender_name")
	r.To = c.FormValue("to")
	r.CC = c.FormValue("cc")
	r.Subject = c.FormValue("subject")
	r.Body = c.FormValue("body")
	r.Attach = file.Filename

	log.Println("Send Email Request", r)
	errSend := NewGoMail().Send(ctx, r)
	if errSend != nil {
		log.Println("Error Send Email", err)
		return c.JSON(helper.HTTPStatusFromCode(helper.InvalidArgument), &helper.Response{
			Code:    helper.InvalidArgument,
			Message: helper.StatusMessage[helper.InvalidArgument],
		})
	}

	return c.JSON(helper.HTTPStatusFromCode(helper.Success), &helper.Response{
		Code:    helper.SuccessCreated,
		Message: helper.StatusMessage[helper.SuccessCreated],
	})
}

func NewMail() *Mail {
	return &Mail{}
}
