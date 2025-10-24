package main

import (
	"log"

	"github.com/biryanim/wb_tech_calendar/internal/config"
)

const (
	envFilePath = ".env"
)

func main() {
	err := config.Load(envFilePath)
	if err != nil {
		log.Fatalf("load config: %v", err)
	}
}
