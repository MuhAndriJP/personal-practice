package handler

import (
	"context"
	"encoding/json"
	"log"

	"github.com/MuhAndriJP/gateway-service.git/entity"
	pb "github.com/MuhAndriJP/gateway-service.git/grpc/user"
	"github.com/MuhAndriJP/gateway-service.git/grpc/user/client"
	"github.com/MuhAndriJP/gateway-service.git/helper"
	"github.com/labstack/echo/v4"
)

type RegisterUser struct {
	client client.Client
}

func (u *RegisterUser) Handle(c echo.Context) (err error) {
	ctx := c.Request().Context()
	if ctx == nil {
		ctx = context.Background()
	}

	r := new(entity.Users)
	if err := c.Bind(r); err != nil {
		return err
	}

	req := pb.RegisterUserRequest{}

	bytes, _ := json.Marshal(&r)
	_ = json.Unmarshal(bytes, &req)

	log.Println("Register User Request", &req)
	_, err = u.client.RegisterUser(ctx, &req)
	if err != nil {
		log.Println("[ERROR] Register User", err)
		return
	}

	resp := &helper.Response{
		Code:    helper.SuccessCreated,
		Message: helper.StatusMessage[helper.SuccessCreated],
		Data: map[string]interface{}{
			"message": helper.StatusMessage[helper.SuccessCreated],
		},
	}

	return c.JSON(helper.HTTPStatusFromCode(helper.Success), resp)
}

func NewRegisterUser() *RegisterUser {
	return &RegisterUser{}
}
