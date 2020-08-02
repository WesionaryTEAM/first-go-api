package Controllers

import (
	"first-api-go/dto"
	"first-api-go/middlewares"
	"first-api-go/service"

	"github.com/gin-gonic/gin"
)

// type LoginController struct {
// 	Login(ctx *gin.Context) string
// }

type loginController struct {
	loginService service.LoginService
	jwtService   middlewares.JWTService
}

func LoginHandler(loginService service.LoginService,
	jwtService middlewares.JWTService) LoginController {
	return &LoginController{
		loginService: loginService,
		jwtService:   jwtService,
	}

}

func (controller *loginController) Login(ctx *gin.Context) string {
	var credential dto.LoginCredentials
	err := ctx.ShouldBind(&credential)
	if err != nil {
		return "no data found"
	}
	isUserAuthenticated := controller.loginService.LoginUser(credential.Email, credential.Password)
	if isUserAuthenticated {
		return controller.jwtService.GenerateToken(credential.Email, true)
	}
	return ""
}
