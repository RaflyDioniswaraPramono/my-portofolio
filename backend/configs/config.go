package configs

import (
	"fmt"
	"log"
	"os"

	"github.com/joho/godotenv"
)

func LoadConfig(stage string, key string) any {
	filePath := fmt.Sprintf("./configs/%v.env", stage)

	err := godotenv.Load(filePath)

	if err != nil {
		log.Fatalf("Erro: %v", err)
	}

	configValue := os.Getenv(key)

	return configValue
}
