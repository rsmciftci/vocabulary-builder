package controllers

import (
	"net/http"
	"vocabulary-builder/dto"
	"vocabulary-builder/models"
	"vocabulary-builder/services"

	"github.com/gin-gonic/gin"
)

type VideoController struct {
	videoService services.VideoService
}

func (c *VideoController) SaveVideo(ctx *gin.Context) {
	var video models.Video

	if err := ctx.ShouldBindJSON(&video); err != nil {
		ctx.JSON(http.StatusBadRequest, gin.H{"Message": err.Error()})
		return
	}

	err := c.videoService.SaveVideo(&video)

	if err != nil {
		ctx.JSON(http.StatusBadGateway, gin.H{"Message": err.Error()})
		return
	}
	ctx.JSON(http.StatusOK, gin.H{"Message": "Success"})
}

func (c *VideoController) DeleteVideo(ctx *gin.Context) {
	var id string = ctx.Param("id")
	err := c.videoService.DeleteVideo(id)
	if err != nil {
		ctx.JSON(http.StatusBadGateway, gin.H{"mMessage": err.Error()})
		return
	}
	ctx.JSON(http.StatusOK, gin.H{"Message": "Success"})

}

func (c *VideoController) UpdateVideo(ctx *gin.Context) {
	var video models.Video
	if err := ctx.ShouldBindJSON(&video); err != nil {
		ctx.JSON(http.StatusBadRequest, gin.H{"Message": err.Error()})
		return
	}
	err := c.videoService.UpdateVideo(&video)
	if err != nil {
		ctx.JSON(http.StatusBadGateway, gin.H{"Message": err.Error()})
		return
	}
	ctx.JSON(http.StatusOK, gin.H{"Message": "Success"})
}

func (c *VideoController) GetAllVideos(ctx *gin.Context) {
	videos, err := c.videoService.GetAllVideos()
	if err != nil {
		ctx.JSON(http.StatusBadGateway, gin.H{"Message": err.Error()})
	}
	ctx.JSON(http.StatusOK, videos)
}

func (c *VideoController) FindAVideo(ctx *gin.Context) {
	var videoDto dto.VideoDto
	if err := ctx.ShouldBindJSON(&videoDto); err != nil {
		ctx.JSON(http.StatusBadRequest, gin.H{"Message": err.Error()})
		return
	}
	video, err := c.videoService.FindAVideo(&videoDto)
	if err != nil {
		ctx.JSON(http.StatusNotFound, gin.H{"Message": err.Error()})
		return
	}
	ctx.JSON(http.StatusOK, video)
}

func (videoController *VideoController) RegisterVideoRoutes(routerGroup *gin.RouterGroup) {
	userRoute := routerGroup.Group("/video")
	userRoute.POST("", videoController.SaveVideo)
	userRoute.DELETE("/:id", videoController.DeleteVideo)
	userRoute.PATCH("", videoController.UpdateVideo)
	userRoute.GET("", videoController.GetAllVideos)
	userRoute.GET("/:name", videoController.FindByName)
}
