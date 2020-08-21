package routes

import (
	"cloud-upload/controller"
	"net/http"

	"github.com/gin-gonic/gin"
)

func SetupRouter() *gin.Engine {
	r := gin.Default()

	//For testing server
	r.GET("/", func(c *gin.Context) {
		c.JSON(http.StatusOK, gin.H{"data": "Up and Running..."})
	})

	api := r.Group("/api")
	{
		api.POST("upload", controller.CreatePerson)
	}
	return r
}
