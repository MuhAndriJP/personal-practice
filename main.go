package main

import (
	"os"

	"github.com/MuhAndriJP/gateway-service.git/routes"
)

func main() {
	e := routes.New()

	e.Logger.Fatal(e.Start(":" + os.Getenv("PORT")))
}
