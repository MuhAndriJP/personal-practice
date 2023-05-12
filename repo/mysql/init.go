package mysql

import (
	"log"
	"os"

	"github.com/MuhAndriJP/personal-practice.git/entity"
	"gorm.io/driver/mysql"
	"gorm.io/gorm"
)

var DB *gorm.DB

func InitDB() {
	usernameAndPassword := os.Getenv("db_user") + ":" + os.Getenv("db_password")
	hostName := "tcp(" + os.Getenv("db_host") + ":" + os.Getenv("db_port") + ")"
	urlConnection := usernameAndPassword + "@" + hostName + "/" + os.Getenv("db_database") + "?charset=utf8&parseTime=true&loc=UTC"
	if os.Getenv("APP_ENV") == "prod" {
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
}
