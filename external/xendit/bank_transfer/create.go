package bank_transfer

import (
	"context"
	"log"

	entity "github.com/MuhAndriJP/personal-practice.git/entity/xendit"
	"github.com/MuhAndriJP/personal-practice.git/helper"
	"github.com/labstack/echo/v4"
)

type CreateInvoice struct{}

func (e *CreateInvoice) CreateEWalletCharge(c echo.Context) error {
	ctx := c.Request().Context()
	if ctx == nil {
		ctx = context.Background()
	}

	r := new(entity.CreateInvoiceRequest)
	if err := c.Bind(r); err != nil {
		return err
	}

	log.Println("Create Invoice Request", r)
	res, err := NewXenditBankTransfer().CreateInvoice(ctx, r)
	if err != nil {
		log.Println("Error Create Invoice", err)
		return c.JSON(helper.HTTPStatusFromCode(helper.InvalidArgument), &helper.Response{
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

func NewCreateInvoice() *CreateInvoice {
	return &CreateInvoice{}
}
