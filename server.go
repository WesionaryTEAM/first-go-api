package main

import (
	"go-jwt/controller"
	"go-jwt/service"
	"net/http"

	"github.com/gin-gonic/gin"
)

var (
	videoService    service.VideoService       = service.New()
	videoController controller.VideoController = controller.New(videoService)
)

func main() {

	var loginService service.LoginService = service.NewLoginService()
	var jwtService service.JWTService = service.NewJWTService()
	var loginController controller.LoginController = controller.LoginHandler(loginService, jwtService)

	server := gin.New()

	server.POST("/login", func(ctx *gin.Context) {
		token := loginController.Login(ctx)
		if token != "" {
			ctx.JSON(http.StatusOK, gin.H{
				"token": token,
			})
		} else {
			ctx.JSON(http.StatusUnauthorized, nil)
		}
	})

	// server.Use(gin.Recovery(), middlewares.Logger(), middlewares.BasicAuth())

	// server.GET("/videos", func(c *gin.Context) {
	// 	c.JSON(200, videoController.FindAll())
	// })

	// server.POST("/videos", func(c *gin.Context) {
	// 	err := videoController.Save(c)
	// 	if err != nil {
	// 		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
	// 	} else {
	// 		c.JSON(http.StatusOK, gin.H{"message": "Video Input is valid"})
	// 	}

	// })

	server.Run(":8080")
}
