package main

import (
	"first-api-go/Config"
	"first-api-go/Models"
	"first-api-go/Routes"
	"fmt"

	"github.com/jinzhu/gorm"
)

var err error

func main() {
	Config.DB, err = gorm.Open("mysql",
		Config.DbURL(Config.BuildDBConfig()))

	if err != nil {
		fmt.Println("Status: ", err)
	}

	defer Config.DB.Close()
	Config.DB.AutoMigrate(&Models.User{})

	r := Routes.SetupRouter()
	r.Run()
}
