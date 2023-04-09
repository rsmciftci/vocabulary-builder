package services

import (
	"vocabulary-builder/dto"
	"vocabulary-builder/models"
)

type UserWordService interface {
	SaveUserWord(*models.UserWord) error
	UpdateUserWord(*dto.UserWordDto) error
	GetAllByUser(*string) ([]*models.UserWord, error)
	DeleteUserWord(*string) error
}
