package user

import (
	"encoding/json"
	"log"

	"github.com/RaflyDioniswaraPramono/my-portofolio/helper"
	"gorm.io/gorm"
)

func UserSeeder(db *gorm.DB) {
	file := helper.LoadJSONFile("user")

	var users []User

	err := json.Unmarshal(file, &users)

	if err != nil {
		log.Fatalf("Error: %v", err)
	}

	for _, user := range users {
		encryptedPassword := helper.EncryptPassword([]byte(user.Password))

		user.Password = string(encryptedPassword)

		db.FirstOrCreate(&user, User{Username: user.Username, Email: user.Email})
	}
}
