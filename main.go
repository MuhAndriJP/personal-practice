package main

import (
	"os"

	"github.com/MuhAndriJP/personal-practice.git/repo/mysql"
	"github.com/MuhAndriJP/personal-practice.git/routes"
)

func main() {
	mysql.InitDB()
	e := routes.New()

	e.Logger.Fatal(e.Start(":" + os.Getenv("PORT")))
}
