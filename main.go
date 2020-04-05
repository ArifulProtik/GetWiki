package main

import (
	"getwiki/cli"
	"log"

	"github.com/joho/godotenv"
)

func main() {
	// Getting All The Bootable Configs from the env file
	err := godotenv.Load()
	if err != nil {
		log.Fatal("Error loading .env file")
	}
	cli.Execute()
}
