package controller

import (
	"cloud-upload/models"
	"net/http"

	"github.com/gin-gonic/gin"
	"gorm.io/gorm"
)

//Adding a person
func CreatePerson(c *gin.Context) {
	db := c.MustGet("db").(*gorm.DB)
	var person models.Person
	c.BindJSON(&person)
	_, err := models.Save(db)
	if err != nil {
		c.AbortWithStatus(http.StatusNotFound)
	} else {
		c.JSON(http.StatusOK, person)
	}
}
