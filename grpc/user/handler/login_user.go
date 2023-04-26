package handler

import (
	"context"
	"encoding/json"

	"github.com/MuhAndriJP/gateway-service.git/entity"
	pb "github.com/MuhAndriJP/gateway-service.git/grpc/user"
	"github.com/MuhAndriJP/gateway-service.git/grpc/user/client"
	"github.com/MuhAndriJP/gateway-service.git/helper"
	"github.com/labstack/echo/v4"
)

type LoginUser struct {
	client client.Client
}

func (u *LoginUser) Handle(c echo.Context) (err error) {
	ctx := c.Request().Context()
	if ctx == nil {
		ctx = context.Background()
	}

	r := new(entity.Users)
	if err := c.Bind(r); err != nil {
		return err
	}

	req := pb.LoginUserRequest{}

	bytes, _ := json.Marshal(&r)
	_ = json.Unmarshal(bytes, &req)

	res, err := u.client.LoginUser(ctx, &req)
	if err != nil {
		return
	}

	resp := &helper.Response{
		Code:    helper.SuccessCreated,
		Message: helper.StatusMessage[helper.SuccessCreated],
		Data: map[string]interface{}{
			"message": helper.StatusMessage[helper.Success],
			"data":    res,
		},
	}

	return c.JSON(helper.HTTPStatusFromCode(helper.Success), resp)
}

func NewLoginUser() *LoginUser {
	return &LoginUser{}
}
