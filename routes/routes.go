package routes

import (
	// "project_altabe4_1/external/google"

	"github.com/MuhAndriJP/gateway-service.git/grpc/user/handler"
	"github.com/labstack/echo/v4"
)

func New() *echo.Echo {

	e := echo.New()

	e.POST("/api/users/register", handler.NewRegisterUser().Handle)
	e.POST("/api/users/login", handler.NewLoginUser().Handle)

	// Google
	// e.GET("/api/google", google.NewGoogleAuth().HandleGoogleLogin)
	// e.GET("/api/UserOauth", google.NewGoogleAuth().HandleGoogleCallback)

	return e
}
