package ewallet

import (
	"context"
	"log"

	entity "github.com/MuhAndriJP/personal-practice.git/entity/xendit"
	"github.com/MuhAndriJP/personal-practice.git/helper"
	"github.com/labstack/echo/v4"
)

type EWalletCharge struct{}

func (e *EWalletCharge) CreateEWalletCharge(c echo.Context) error {
	ctx := c.Request().Context()
	if ctx == nil {
		ctx = context.Background()
	}

	r := new(entity.CreateEWalletChargeRequest)
	if err := c.Bind(r); err != nil {
		return err
	}

	log.Println("Create EWallet Charge Request", r)
	res, err := NewXendit().CreateEWalletCharge(ctx, r)
	if err != nil {
		log.Println("Error Create EWallet Charge", err)
		return c.JSON(helper.HTTPStatusFromCode(helper.Success), &helper.Response{
			Code:    helper.InvalidArgument,
			Message: helper.StatusMessage[helper.InvalidArgument],
		})
	}

	return c.JSON(helper.HTTPStatusFromCode(helper.Success), &helper.Response{
		Code:    helper.SuccessCreated,
		Message: helper.StatusMessage[helper.SuccessCreated],
		Data: map[string]interface{}{
			"data": res,
		},
	})
}

func NewEwalletCharge() *EWalletCharge {
	return &EWalletCharge{}
}
