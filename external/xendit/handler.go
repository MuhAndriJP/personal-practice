package xendit

import (
	"context"
	"log"

	"github.com/labstack/echo/v4"
)

func CreateEWalletCharge(c echo.Context) error {
	ctx := c.Request().Context()
	if ctx == nil {
		ctx = context.Background()
	}

	r := new(CreateEWalletChargeRequest)
	if err := c.Bind(r); err != nil {
		return err
	}

	log.Println("Create EWallet Charge Request", r)
	res, err := NewXendit().CreateEWalletCharge(ctx, r)
	if err != nil {
		log.Println("Error Create EWallet Charge", err)
		return err
	}

	return c.JSON(201, map[string]interface{}{
		"data": res,
	})
}

func GetEWalletChargeStatus(c echo.Context) error {
	ctx := c.Request().Context()
	if ctx == nil {
		ctx = context.Background()
	}

	r := new(GetEWalletChargeStatusRequest)
	if err := c.Bind(r); err != nil {
		return err
	}

	chargeID := c.Param("id")
	r.ChargeID = chargeID

	log.Println("Get EWallet Charge Status Request", r)
	res, err := NewXendit().GetEWalletChargeStatus(ctx, r)
	if err != nil {
		log.Println("Error Create EWallet Charge", err)
		return err
	}

	return c.JSON(201, map[string]interface{}{
		"data": res,
	})
}

func CreateEWalletCallback(c echo.Context) error {

	ctx := c.Request().Context()
	if ctx == nil {
		ctx = context.Background()
	}

	r := new(CreateEWalletCallbackRequest)
	if err := c.Bind(r); err != nil {
		return err
	}

	log.Println("Create Callback Request", r)
	err := NewXendit().CreateEWalletCallback(ctx)
	if err != nil {
		log.Println("Error Create EWallet Callback", err)
		return err
	}

	return c.JSON(200, map[string]interface{}{
		"data": "success",
	})
}
