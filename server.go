package main

import (
	"go-jwt/controller"
	"go-jwt/middlewares"
	"go-jwt/service"
	"io"
	"net/http"
	"os"

	"github.com/gin-gonic/gin"
)

var (
	videoService    service.VideoService       = service.New()
	videoController controller.VideoController = controller.New(videoService)
)

func setupLogOutput() {
	f, _ := os.Create("gin.log")
	gin.DefaultWriter = io.MultiWriter(f, os.Stdout)
}

func main() {

	setupLogOutput()

	var loginService service.LoginService = service.NewLoginService()
	var jwtService service.JWTService = service.NewJWTService()
	var loginController controller.LoginController = controller.LoginHandler(loginService, jwtService)

	server := gin.New()

	server.Use(gin.Recovery(), gin.Logger())

	//Login Endpoint: Authentication + Token Creation
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

	//JWT Authorization Middleware applies to "/api" only
	apiRoutes := server.Group("/api", middlewares.AuthorizeJWT())
	{
		apiRoutes.GET("/videos", func(ctx *gin.Context) {
			ctx.JSON(200, videoController.FindAll())
		})

		apiRoutes.POST("/videos", func(ctx *gin.Context) {
			err := videoController.Save(ctx)
			if err != nil {
				ctx.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
			} else {
				ctx.JSON(http.StatusOK, gin.H{"message": "Video Input is Valid"})
			}
		})
	}

	//Basic Auth

	//server.Use(gin.Recovery(), middlewares.Logger(), middlewares.BasicAuth())

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
