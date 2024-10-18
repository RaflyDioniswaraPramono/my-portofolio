package role

import (
	"encoding/json"
	"log"

	"github.com/RaflyDioniswaraPramono/my-portofolio/helper"
	"gorm.io/gorm"
)

func RoleSeeder(db *gorm.DB) {
	file := helper.LoadJSONFile("role")

	var roles []Role

	err := json.Unmarshal(file, &roles)

	if err != nil {
		log.Fatalf("Error: %v", err)
	}

	for _, role := range roles {
		db.FirstOrCreate(&role, Role{RoleName: role.RoleName})
	}
}
