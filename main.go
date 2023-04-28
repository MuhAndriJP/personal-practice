package main

import (
	"os"

	"github.com/MuhAndriJP/gateway-service.git/middleware"
	"github.com/MuhAndriJP/gateway-service.git/routes"
)

func main() {
	e := routes.New()
	middleware.LogMiddlewares(e)

	e.Logger.Fatal(e.Start(":" + os.Getenv("PORT")))
}
