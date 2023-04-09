package controllers

import (
	"net/http"
	"vocabulary-builder/dto"
	"vocabulary-builder/models"
	"vocabulary-builder/services"

	"github.com/gin-gonic/gin"
)

type WordController struct {
	WordService services.WordService
}

func NewWordController(WordService services.WordService) WordController {
	return WordController{
		WordService: WordService,
	}
}

func (c *WordController) SaveWord(ctx *gin.Context) {
	var word models.Word
	if err := ctx.ShouldBindJSON(&word); err != nil {
		ctx.JSON(http.StatusBadRequest, gin.H{"Message": err.Error()})
		return
	}
	err := c.WordService.SaveWord(&word)
	if err != nil {
		ctx.JSON(http.StatusBadGateway, gin.H{"Message": err.Error()})
		return
	}
	ctx.JSON(http.StatusOK, gin.H{"Message": "Success"})
}
func (c *WordController) DeleteWordById(ctx *gin.Context) {
	id := ctx.Param("id")
	err := c.WordService.DeleteWordById(id)
	if err != nil {
		ctx.JSON(http.StatusBadGateway, gin.H{"Message": err.Error()})
		return
	}
	ctx.JSON(http.StatusOK, gin.H{"Message": "Success"})
}

func (c *WordController) UpdateWord(ctx *gin.Context) {
	var word dto.Word
	if err := ctx.ShouldBindJSON(&word); err != nil {
		ctx.JSON(http.StatusBadRequest, gin.H{"Message": err.Error()})
		return
	}
	err := c.WordService.UpdateWord(&word)
	if err != nil {
		ctx.JSON(http.StatusBadGateway, gin.H{"Message": err.Error()})
		return
	}
	ctx.JSON(http.StatusOK, gin.H{"Message": "Success"})
}

func (c *WordController) FindWord(ctx *gin.Context) {
	word := ctx.Param("word")
	result, err := c.WordService.FindWord(&word)
	if err != nil {
		ctx.JSON(http.StatusBadGateway, gin.H{"Message": err.Error()})
		return
	}
	ctx.JSON(http.StatusOK, result)

}

func (c *WordController) RegisterUserRoutes(routerGroup *gin.RouterGroup) {
	userRoute := routerGroup.Group("/word")
	userRoute.POST("", c.SaveWord)
	userRoute.DELETE("/:id", c.DeleteWordById)
	userRoute.PUT("", c.UpdateWord)
	userRoute.GET("/:word", c.FindWord)
}
