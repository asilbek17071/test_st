package controllers

import (
	"net/http"

	"github.com/asilbek17071/test_st/models"
	"github.com/gin-gonic/gin"
	"gorm.io/gorm"
)

type UserController struct {
	DB *gorm.DB
}

func NewUserController(DB *gorm.DB) UserController {
	return UserController{DB}
}

func (pc *UserController) FindUserByName(ctx *gin.Context) {
	userName := ctx.Param("name")

	var user models.UserResponseByName
	result := pc.DB.Table("users").First(&user, "name = ?", userName)
	if result.Error != nil {
		ctx.JSON(http.StatusNotFound, gin.H{"status": "fail", "message": "No user with that title exists"})
		return
	}

	ctx.JSON(http.StatusOK, gin.H{"status": "success", "data": user})
}
