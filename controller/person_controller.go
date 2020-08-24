package controller

import (
	"cloud-upload/imageUpload"
	"cloud-upload/models"
	"fmt"
	"log"
	"net/http"

	"github.com/gin-gonic/gin"
)

//Adding a person
func CreatePerson(c *gin.Context) {
	// db := c.MustGet("db").(*gorm.DB)
	var person models.Person

	// ParseMultipartForm parses a request body as multipart/form-data
	_ = c.Request.ParseMultipartForm(32 << 20)
	person = models.Person{}
	var photo = ""

	errForm := c.Request.ParseForm()
	log.Println(errForm)
	//name := c.Request.Form.Get("name")
	photoPath := c.Request.Form.Get("photo")

	//id := c.Param("id")

	file, _, err := c.Request.FormFile("photo")

	switch err {
	case nil:
		defer file.Close()

		photo = imageUpload.ImageUpload(photoPath)
		if photo == "" {
			fmt.Println("error getting url of image")
			return
		}
	default:
		photo = "picture could not be uploaded"
		log.Println(err)
		return
	}

	c.BindJSON(&person)
	err = models.AddPerson(&person)
	if err != nil {
		c.AbortWithStatus(http.StatusNotFound)
	} else {
		c.JSON(http.StatusOK, person)
	}
}
