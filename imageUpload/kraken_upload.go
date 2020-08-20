package imageUpload

import (
	"log"
	"os"

	"github.com/joho/godotenv"
	"github.com/kraken-io/kraken-go"
)

func ImageUpload() {
	err := godotenv.Load(".env")
	if err != nil {
		log.Fatalf("Error loading .env file")
		os.Exit(1)
	}
	api_key := os.Getenv("API_Key")
	api_secret := os.Getenv("API_Secret")

	kr, err := kraken.New(api_key, api_secret)

	if err != nil {
		log.Fatal(err)
	}

	params := map[string]interface{}{
		"wait": true,
	}

	imgPath := `C:\Users\HP\Downloads\my_picture.jpg`

	data, err := kr.Upload(params, imgPath)

	if err != nil {
		log.Fatal(err)
	}

	if data["success"] != true {
		log.Println("Failed, error message", data["message"])
	} else {
		log.Println("Success, Optimized image URL: ", data["kraked_url"])
	}

}
