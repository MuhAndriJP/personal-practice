package mysql

import (
	"log"
	"os"

	"github.com/MuhAndriJP/personal-practice.git/entity"
	xendit "github.com/MuhAndriJP/personal-practice.git/entity/xendit"
	"gorm.io/driver/mysql"
	"gorm.io/gorm"
)

var DB *gorm.DB

func InitDB() {
	usernameAndPassword := os.Getenv("MYSQL_USER") + ":" + os.Getenv("MYSQL_PASSWORD")
	hostName := "tcp(" + os.Getenv("MYSQL_HOST") + ":" + os.Getenv("MYSQL_PORT") + ")"
	urlConnection := usernameAndPassword + "@" + hostName + "/" + os.Getenv("MYSQL_DATABASE") + "?charset=utf8&parseTime=true&loc=UTC"
	if os.Getenv("APP_ENV") == "PROD" {
		urlConnection += "&tls=true"
	}

	var e error
	DB, e = gorm.Open(mysql.Open(urlConnection), &gorm.Config{})
	if e != nil {
		panic(e)
	}

	log.Println("Connect MySQL ", urlConnection)
	InitMigrate()
}

func InitMigrate() {
	// DB.Migrator().DropTable(&entity.Users{})
	DB.AutoMigrate(&entity.Users{})
	DB.AutoMigrate(&xendit.EWalletPayment{})
}
