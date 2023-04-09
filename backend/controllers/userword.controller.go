package controllers

import (
	"net/http"
	"vocabulary-builder/dto"
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
	var userWordDto *dto.UserWordDto
	if err := ctx.ShouldBindJSON(userWordDto); err != nil {
		ctx.JSON(http.StatusBadRequest, gin.H{"Message": err.Error()})
		return
	}
	err := uwc.UserWordService.UpdateUserWord(userWordDto)
	if err != nil {
		ctx.JSON(http.StatusBadGateway, gin.H{"Message": err.Error()})
		return
	}
	ctx.JSON(http.StatusOK, gin.H{"Message": "Success"})
}

func (uwc *UserWordController) GetAllByUser(ctx *gin.Context) {
	userId := ctx.Param("userId")
	userWordList, err := uwc.UserWordService.GetAllByUser(&userId)
	if err != nil {
		ctx.JSON(http.StatusBadGateway, gin.H{"Message": err.Error()})
		return
	}
	ctx.JSON(http.StatusOK, userWordList)
}

func (uwc *UserWordController) DeleteUserWord(ctx *gin.Context) {
	wordId := ctx.Param("wordId")
	err := uwc.UserWordService.DeleteUserWord(&wordId)
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
	userRoute.GET("/:userId", uwc.GetAllByUser)
	userRoute.DELETE("/:wordId", uwc.DeleteUserWord)
}
