package main

import (
	"fmt"
	"todo-app/Config"
	"todo-app/Models"
	"todo-app/Routes"

	"github.com/jinzhu/gorm"
)

var err error

func main() {
	//Connection with database
	Config.DB, err = gorm.Open("mysql", Config.DbURL(Config.BuildDBConfig()))

	if err != nil {
		fmt.Println("status:", err)
	}

	defer Config.DB.Close()

	//migration running
	Config.DB.AutoMigrate(&Models.Todo{})

	//setup Routes
	r := Routes.SetupRouter()

	//running
	r.Run()
}
