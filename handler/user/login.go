package user

import (
	"context"
	"log"

	"github.com/MuhAndriJP/personal-practice.git/action/user"
	"github.com/MuhAndriJP/personal-practice.git/entity"
	"github.com/MuhAndriJP/personal-practice.git/helper"
	"github.com/labstack/echo/v4"
)

type UserLogin struct {
}

func (u *UserLogin) Handle(c echo.Context) (err error) {
	ctx := c.Request().Context()
	if ctx == nil {
		ctx = context.Background()
	}

	req := new(entity.Users)
	if err = c.Bind(req); err != nil {
		return
	}

	log.Println("Register User Request", &req)
	res, err := user.NewUserLogin().Handle(ctx, req)
	if err != nil {
		return
	}

	resp := &helper.Response{
		Code:    helper.SuccessCreated,
		Message: helper.StatusMessage[helper.SuccessCreated],
		Data: map[string]interface{}{
			"data": res,
		},
	}

	return c.JSON(helper.HTTPStatusFromCode(helper.Success), resp)
}

func NewUserLogin() *UserLogin {
	return &UserLogin{}
}
