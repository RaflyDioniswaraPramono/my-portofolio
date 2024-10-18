package helper

import (
	"fmt"
	"log"
	"os"
)

func LoadJSONFile(fileName string) []byte {
	filePath := fmt.Sprintf("./configs/seeds/%v-seeds.json", fileName)

	file, err := os.ReadFile(filePath)

	if err != nil {
		log.Fatalf("Error: %v", err)
	}

	return file
}
