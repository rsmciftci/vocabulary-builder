package controllers

import (
	"net/http"
	"vocabulary-builder/dto"
	"vocabulary-builder/models"
	"vocabulary-builder/services"

	"github.com/gin-gonic/gin"
)

type UserController struct {
	UserService services.UserService
}

func NewUserController(UserService services.UserService) UserController {
	return UserController{
		UserService: UserService,
	}
}

func (userController *UserController) SaveUser(ctx *gin.Context) {
	var user models.User
	if err := ctx.ShouldBindJSON(&user); err != nil {
		ctx.JSON(http.StatusBadRequest, gin.H{"Message": err.Error()})
		return
	}

	err := userController.UserService.SaveUser(&user)
	if err != nil {
		ctx.JSON(http.StatusBadGateway, gin.H{"Message": err.Error()})
		return
	}
	ctx.JSON(http.StatusOK, gin.H{"Message": "Success"})

}

func (userController *UserController) DeleteByEmail(ctx *gin.Context) {
	var email string = ctx.Param("email")
	err := userController.UserService.DeleteByEmail(&email)
	if err != nil {
		ctx.JSON(http.StatusBadGateway, gin.H{"Message": err.Error()})
		return
	}
	ctx.JSON(http.StatusOK, gin.H{"Message": "Success"})

}

func (userController *UserController) UpdateUser(ctx *gin.Context) {
	var user models.User
	if err := ctx.ShouldBindJSON(&user); err != nil {
		ctx.JSON(http.StatusBadRequest, gin.H{"Message": err.Error()})
		return
	}

	err := userController.UserService.UpdateUser(&user)
	if err != nil {
		ctx.JSON(http.StatusBadGateway, gin.H{"Message": err.Error()})
		return
	}
	ctx.JSON(http.StatusOK, gin.H{"Message": "Success"})

}

func (userController *UserController) FindUserByEmailAndPassword(ctx *gin.Context) {
	var emailAndPassword dto.EmailAndPassword
	if err := ctx.ShouldBindJSON(&emailAndPassword); err != nil {
		ctx.JSON(http.StatusBadRequest, gin.H{"Message": err.Error()})
	}
	err := userController.UserService.FindUserByEmailAndPassword(&emailAndPassword)
	if err != nil {
		ctx.JSON(http.StatusBadGateway, gin.H{"Message": err.Error()})
		return
	}

	ctx.JSON(http.StatusOK, gin.H{"Message": "Success"})
	// TODO: sifrealani olmadan user"i don
}

func (userController *UserController) RegisterUserRoutes(routerGroup *gin.RouterGroup) {
	userRoute := routerGroup.Group("/user")
	userRoute.POST("", userController.SaveUser)
	userRoute.DELETE("/:email", userController.DeleteByEmail)
	userRoute.PUT("", userController.UpdateUser)
	userRoute.GET("", userController.FindUserByEmailAndPassword)
}
