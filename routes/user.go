package routes

import (
	"github.com/asilbek17071/test_st/controllers"
	"github.com/asilbek17071/test_st/middleware"
	"github.com/gin-gonic/gin"
)

type UserRouteController struct {
	userController controllers.UserController
}

func NewRouteUserController(userController controllers.UserController) UserRouteController {
	return UserRouteController{userController}
}

func (uc *UserRouteController) UserRoute(rg *gin.RouterGroup) {

	rg.GET("/:name", middleware.DeserializeUser(), uc.userController.FindUserByName)
}
