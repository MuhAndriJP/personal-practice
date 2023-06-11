package ewallet

import (
	"context"
	"log"

	entity "github.com/MuhAndriJP/personal-practice.git/entity/xendit"
	"github.com/MuhAndriJP/personal-practice.git/helper"
	"github.com/labstack/echo/v4"
)

type EWalletCallback struct{}

func (e *EWalletCallback) CreateEWalletCallback(c echo.Context) error {

	ctx := c.Request().Context()
	if ctx == nil {
		ctx = context.Background()
	}

	r := new(entity.CreateEWalletCallbackRequest)
	if err := c.Bind(r); err != nil {
		return err
	}

	log.Println("Create Callback Request", r)
	err := NewXendit().CreateEWalletCallback(ctx, r)
	if err != nil {
		log.Println("Error Create EWallet Callback", err)
		return c.JSON(helper.HTTPStatusFromCode(helper.Success), &helper.Response{
			Code:    helper.InvalidArgument,
			Message: helper.StatusMessage[helper.InvalidArgument],
		})
	}

	return c.JSON(helper.HTTPStatusFromCode(helper.Success), &helper.Response{
		Code:    helper.SuccessCreated,
		Message: helper.StatusMessage[helper.SuccessCreated],
	})
}

func NewEWalletCallback() *EWalletCallback {
	return &EWalletCallback{}
}
