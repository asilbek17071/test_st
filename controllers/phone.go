package controllers

import (
	"net/http"
	"strings"
	"time"

	"github.com/asilbek17071/test_st/models"
	"github.com/asilbek17071/test_st/utils"
	"github.com/gin-gonic/gin"
	"gorm.io/gorm"
)

type PhoneController struct {
	DB *gorm.DB
}

func NewPhoneController(DB *gorm.DB) PhoneController {
	return PhoneController{DB}
}

func (pc *PhoneController) CreatePhone(ctx *gin.Context) {
	currentUser := ctx.MustGet("currentUser").(models.User)
	var payload *models.CreatePhoneRequest

	if err := ctx.ShouldBindJSON(&payload); err != nil {
		ctx.JSON(http.StatusBadRequest, err.Error())
		return
	}

	now := time.Now()
	newPhone := models.Phone{
		Phone:       payload.Phone,
		Description: payload.Description,
		IsMobile:    payload.IsMobile,
		User:        currentUser.ID,
		CreatedAt:   now,
		UpdatedAt:   now,
	}

	result := pc.DB.Create(&newPhone)
	if result.Error != nil {
		if strings.Contains(result.Error.Error(), "duplicate key") {
			ctx.JSON(http.StatusConflict, gin.H{"status": "fail", "message": "Phone with that title already exists"})
			return
		}
		ctx.JSON(http.StatusBadGateway, gin.H{"status": "error", "message": result.Error.Error()})
		return
	}

	ctx.JSON(http.StatusCreated, gin.H{"status": "success", "data": newPhone})
}

func (pc *PhoneController) FindPhoneById(ctx *gin.Context) {
	queryParams := ctx.Request.URL.Query()

	params, errStr := utils.ParseQueryParams(queryParams)
	if errStr != nil {
		ctx.JSON(http.StatusNotFound, gin.H{"status": "fail", "message": "failed to parse query params json"})
		return
	}

	var phone models.PhoneResponse
	result := pc.DB.Table("phones").First(&phone, "phone = ?", params.QQQ)
	if result.Error != nil {
		ctx.JSON(http.StatusNotFound, gin.H{"status": "fail", "message": "No phone with that title exists"})
		return
	}

	ctx.JSON(http.StatusOK, gin.H{"status": "success", "data": phone})
}

func (pc *PhoneController) UpdatePhone(ctx *gin.Context) {
	phoneId := ctx.Param("phoneId")
	currentUser := ctx.MustGet("currentUser").(models.User)

	var payload *models.UpdatePhone
	if err := ctx.ShouldBindJSON(&payload); err != nil {
		ctx.JSON(http.StatusBadGateway, gin.H{"status": "fail", "message": err.Error()})
		return
	}
	var updatedPhone models.Phone
	result := pc.DB.First(&updatedPhone, "id = ?", phoneId)
	if result.Error != nil {
		ctx.JSON(http.StatusNotFound, gin.H{"status": "fail", "message": "No phone with that title exists"})
		return
	}
	now := time.Now()
	phoneToUpdate := models.Phone{
		Phone:       payload.Phone,
		Description: payload.Description,
		IsMobile:    payload.IsMobile,
		User:        currentUser.ID,
		CreatedAt:   updatedPhone.CreatedAt,
		UpdatedAt:   now,
	}

	pc.DB.Model(&updatedPhone).Updates(phoneToUpdate)

	ctx.JSON(http.StatusOK, gin.H{"status": "success", "data": updatedPhone})
}

func (pc *PhoneController) DeletePhone(ctx *gin.Context) {
	phoneId := ctx.Param("phoneId")

	result := pc.DB.Delete(&models.Phone{}, "id = ?", phoneId)

	if result.Error != nil {
		ctx.JSON(http.StatusNotFound, gin.H{"status": "fail", "message": "No phone with that title exists"})
		return
	}

	ctx.JSON(http.StatusNoContent, nil)
}