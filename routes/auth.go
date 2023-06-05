package routes

import (
	"github.com/asilbek17071/test_st/controllers"
	"github.com/gin-gonic/gin"
)

type AuthRouteController struct {
	authController controllers.AuthController
}

func NewAuthRouteController(authController controllers.AuthController) AuthRouteController {
	return AuthRouteController{authController}
}

func (rc *AuthRouteController) AuthRoute(rg *gin.RouterGroup) {

	rg.POST("/register", rc.authController.SignUpUser)
	rg.POST("/auth", rc.authController.SignInUser)
	// router.GET("/refresh", rc.authController.RefreshAccessToken)
	// router.GET("/logout", middleware.DeserializeUser(), rc.authController.LogoutUser)
}
