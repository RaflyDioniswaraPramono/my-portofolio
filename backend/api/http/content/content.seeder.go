package content

import (
	"encoding/json"
	"log"

	"github.com/RaflyDioniswaraPramono/my-portofolio/helper"
	"gorm.io/gorm"
)

func ContentSeeder(db *gorm.DB) {
	file := helper.LoadJSONFile("content")

	var contents []Content

	err := json.Unmarshal(file, &contents)

	if err != nil {
		log.Fatalf("Error: %v", err)
	}

	for _, content := range contents {
		db.FirstOrCreate(&content, Content{ContentText: content.ContentText})
	}
}
