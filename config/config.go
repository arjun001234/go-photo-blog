package config

import (
	"os"

	"github.com/joho/godotenv"
)

func LoadConfig() {
	if os.Args[len(os.Args)-1] == "env=prod" {
		godotenv.Load(".prod.env")
	} else {
		godotenv.Load(".dev.env")
	}
}
