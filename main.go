package main

import (
	"cloud-upload/config"
	"fmt"
	"log"

	"github.com/joho/godotenv"
)

func main() {

	//imageUpload.ImageUpload()

	var err error

	err = godotenv.Load(".env")
	if err != nil {
		log.Fatalf("Error getting env, %v", err)
	} else {
		fmt.Println("We are getting the env values")
	}

	config.InitializeDatabase()
}
