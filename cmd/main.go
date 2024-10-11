package main

import (
	"github.com/NRKA/home-service/configs"
	"github.com/NRKA/home-service/internal/app"
	"log"
)

func main() {
	config, err := configs.FromEnv()
	if err != nil {
		log.Fatalf("Error loading .env file: %v", err)
	}

	app.Run(config)
}
