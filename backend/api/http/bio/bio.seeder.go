package bio

import (
	"encoding/json"
	"log"

	"github.com/RaflyDioniswaraPramono/my-portofolio/helper"
	"gorm.io/gorm"
)

func BioSeeder(db *gorm.DB) {
	file := helper.LoadJSONFile("bio")

	var bios []Bio

	err := json.Unmarshal(file, &bios)

	if err != nil {
		log.Fatalf("Error: %v", err)
	}

	for _, bio := range bios {
		db.FirstOrCreate(&bio, Bio{UserId: bio.UserId})
	}
}
