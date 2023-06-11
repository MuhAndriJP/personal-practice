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

	r := new(MailRequest)
	if err := c.Bind(r); err != nil {
		return err
	}

	log.Println("Send Email Request", r)
	err := NewGoMail().Send(ctx, r)
	if err != nil {
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
