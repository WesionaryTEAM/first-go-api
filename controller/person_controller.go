package controller

import (
	"cloud-upload/models"
	"net/http"

	"github.com/gin-gonic/gin"
)

//Adding a person
func CreatePerson(c *gin.Context) {
	// db := c.MustGet("db").(*gorm.DB)
	var person models.Person
	c.BindJSON(&person)
	err := models.AddPerson(&person)
	if err != nil {
		c.AbortWithStatus(http.StatusNotFound)
	} else {
		c.JSON(http.StatusOK, person)
	}
}
