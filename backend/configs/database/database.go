package database

import (
	"fmt"

	"github.com/RaflyDioniswaraPramono/my-portofolio/configs"
	"github.com/RaflyDioniswaraPramono/my-portofolio/utils"
	"gorm.io/driver/mysql"
	"gorm.io/gorm"
	"gorm.io/gorm/logger"
)

var DB *gorm.DB

func DBInit() {
	username := configs.LoadConfig("development", "DB_USERNAME")
	password := configs.LoadConfig("development", "DB_PASSWORD")
	hostname := configs.LoadConfig("development", "DB_HOSTNAME")
	port := configs.LoadConfig("development", "DB_PORT")
	name := configs.LoadConfig("development", "DB_NAME")

	dsn := fmt.Sprintf("%v:%v@tcp(%v:%v)/%v", username, password, hostname, port, name)

	db, err := gorm.Open(mysql.Open(dsn), &gorm.Config{
		Logger: logger.Default,
	})

	if err != nil {
		panic("Failed connecting to database!")
	}

	DB = db

	utils.SuccessCLIMessage("Database connected! \n \n")
}
