package user

import (
	"context"
	"log"

	"github.com/MuhAndriJP/personal-practice.git/action/user"
	"github.com/MuhAndriJP/personal-practice.git/entity"
	"github.com/MuhAndriJP/personal-practice.git/helper"
	"github.com/labstack/echo/v4"
)

type UserRegister struct {
}

func (u *UserRegister) Handle(c echo.Context) (err error) {
	ctx := c.Request().Context()
	if ctx == nil {
		ctx = context.Background()
	}

	req := entity.Users{}
	if err = c.Bind(&req); err != nil {
		return
	}

	log.Println("Register User Request", req)
	err = user.NewUserRegister().Handle(ctx, req)
	if err != nil {
		log.Println("[ERROR] User Register:", err)
		return c.JSON(helper.HTTPStatusFromCode(helper.InvalidArgument), &helper.Response{
			Code:    helper.InvalidArgument,
			Message: helper.StatusMessage[helper.InvalidArgument],
			Data: map[string]interface{}{
				"error": err.Error(),
			},
		})
	}

	return c.JSON(helper.HTTPStatusFromCode(helper.Success), &helper.Response{
		Code:    helper.SuccessCreated,
		Message: helper.StatusMessage[helper.SuccessCreated],
	})
}

func NewUserRegister() *UserRegister {
	return &UserRegister{}
}
