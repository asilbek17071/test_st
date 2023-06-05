package routes

import (
	"github.com/asilbek17071/test_st/controllers"
	"github.com/asilbek17071/test_st/middleware"
	"github.com/gin-gonic/gin"
)

type PhoneRouteController struct {
	phoneController controllers.PhoneController
}

func NewRoutePhoneController(phoneController controllers.PhoneController) PhoneRouteController {
	return PhoneRouteController{phoneController}
}

func (pc *PhoneRouteController) PhoneRoute(rg *gin.RouterGroup) {

	router := rg.Group("phones")
	router.Use(middleware.DeserializeUser())
	router.POST("/", pc.phoneController.CreatePhone)
	router.PUT("/:phoneId", pc.phoneController.UpdatePhone)
	router.GET("/", pc.phoneController.FindPhoneById)
	router.DELETE("/:phoneId", pc.phoneController.DeletePhone)
}
