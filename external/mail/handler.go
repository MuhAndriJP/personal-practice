package mail

import (
	"context"
	"log"

	"github.com/labstack/echo/v4"
)

func Handle(c echo.Context) error {
	ctx := c.Request().Context()
	if ctx == nil {
		ctx = context.Background()
	}

	r := new(Mail)
	if err := c.Bind(r); err != nil {
		return err
	}

	log.Println("Send Email Request", r)
	err := NewGoMail().Send(ctx, r)
	if err != nil {
		log.Println("Error Send Email", err)
		return c.JSON(400, map[string]interface{}{
			"error": "Failed send email",
		})
	}

	return c.JSON(201, map[string]interface{}{
		"data": "Succes send email",
	})
}
