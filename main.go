package main

import (
	"log"

	"github.com/joho/godotenv"
)

// Build tag is set up while compiling
var buildTag string

func init() {
	// load environment config
	if err := godotenv.Load(".env"); err != nil {
		log.Fatalf("Error loading .env file: %+v", err)
	}
}

func main() {

}
