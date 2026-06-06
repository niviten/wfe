package main

import (
	"log"

	"github.com/niviten/wfe/internal/config"
	"github.com/niviten/wfe/internal/db"
)

func main() {
	config.Init()
	db.Init()

	log.Println("___done")
}
