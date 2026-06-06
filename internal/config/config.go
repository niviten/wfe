package config

import (
	"log"
	"os"

	"github.com/joho/godotenv"
)

const PREFIX = "WFE_"

func Init() {
	if err := godotenv.Load(); err != nil {
		log.Fatalln("Error while loading env ", err)
	}
}

func GetConfig(name string) string {
	return os.Getenv(PREFIX + name)
}
