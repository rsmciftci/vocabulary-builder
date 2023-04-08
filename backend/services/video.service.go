package services

import (
	"vocabulary-builder/dto"
	"vocabulary-builder/models"
)

type VideoService interface {
	SaveVideo(*models.Video) error
	DeleteVideo(string) error
	UpdateVideo(dto.VideoDto) error
	GetAllVideos() ([]*models.Video, error)
	FindAVideo(*dto.VideoDto) (*models.Video, error) // TODO: name, type, season, episode required as input
}
