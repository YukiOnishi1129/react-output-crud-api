package main

import (
	"fmt"
	"log"
	"os"

	"github.com/joho/godotenv"
)

func main() {
	log.Printf("Start backend service")
	isDev := os.Getenv("GO_ENV") != "production" && os.Getenv("GO_ENV") != "staging"
	if isDev {
		err := godotenv.Load()
		if err != nil {
			log.Fatalf("Error loading .env file")
			return
		}
	}

	fmt.Println("Hello, World!")
}