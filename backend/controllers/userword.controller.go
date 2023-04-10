package controllers

import (
	"net/http"
	"vocabulary-builder/models"
	"vocabulary-builder/services"

	"github.com/gin-gonic/gin"
)

type UserWordController struct {
	UserWordService services.UserWordService
}

func NewUserWordController(UserWordService services.UserWordService) UserWordController {
	return UserWordController{
		UserWordService: UserWordService,
	}
}

func (uwc *UserWordController) SaveUserWord(ctx *gin.Context) {
	var userword *models.UserWord
	if err := ctx.ShouldBindJSON(&userword); err != nil {
		ctx.JSON(http.StatusBadRequest, err.Error())
		return
	}
	err := uwc.UserWordService.SaveUserWord(userword)
	if err != nil {
		ctx.JSON(http.StatusBadGateway, gin.H{"Message": err.Error()})
		return
	}
	ctx.JSON(http.StatusOK, gin.H{"Message": "Success"})

}

func (uwc *UserWordController) UpdateUserWord(ctx *gin.Context) {
	var userword *models.UserWord
	if err := ctx.ShouldBindJSON(&userword); err != nil {
		ctx.JSON(http.StatusBadRequest, gin.H{"Message": err.Error()})
		return
	}
	err := uwc.UserWordService.UpdateUserWord(userword)
	if err != nil {
		ctx.JSON(http.StatusBadGateway, gin.H{"Message": err.Error()})
		return
	}
	ctx.JSON(http.StatusOK, gin.H{"Message": "Success"})
}

func (uwc *UserWordController) GetAllByUser(ctx *gin.Context) {
	uuid := ctx.Param("uuid")
	userWordList, err := uwc.UserWordService.GetAllByUser(&uuid)
	if err != nil {
		ctx.JSON(http.StatusBadGateway, gin.H{"Message": err.Error()})
		return
	}
	ctx.JSON(http.StatusOK, userWordList)
}

func (uwc *UserWordController) DeleteUserWord(ctx *gin.Context) {
	uuid := ctx.Param("uuid")
	err := uwc.UserWordService.DeleteUserWord(&uuid)
	if err != nil {
		ctx.JSON(http.StatusBadGateway, gin.H{"Message": err.Error()})
		return
	}
	ctx.JSON(http.StatusOK, gin.H{"Message": "Success"})

}

func (uwc *UserWordController) RegisterUserRoutes(routerGroup *gin.RouterGroup) {
	userRoute := routerGroup.Group("/userword")
	userRoute.POST("", uwc.SaveUserWord)
	userRoute.PUT("", uwc.UpdateUserWord)
	userRoute.GET("/:uuid", uwc.GetAllByUser)
	userRoute.DELETE("/:uuid", uwc.DeleteUserWord)
}
