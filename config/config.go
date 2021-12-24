package config

import (
	"os"
	"text/template"

	"github.com/joho/godotenv"
)

func LoadTemplates() *template.Template {
	return template.Must(template.ParseGlob("templates/*.gohtml"))
}

func LoadConfig() {
	if os.Args[len(os.Args)-1] == "env=prod" {
		godotenv.Load(".prod.env")
	} else {
		godotenv.Load(".dev.env")
	}
}
