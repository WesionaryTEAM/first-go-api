package routes

import (
	"cloud-upload/controller"

	"github.com/gin-gonic/gin"
)

func SetupRouter() *gin.Engine {
	r := gin.Default()
	api := r.Group("/api")
	{
		v1.POST("upload", controller.CreatePerson)
	}
}
