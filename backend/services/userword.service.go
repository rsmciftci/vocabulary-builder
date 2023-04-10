package services

import (
	"vocabulary-builder/models"
)

type UserWordService interface {
	SaveUserWord(*models.UserWord) error
	UpdateUserWord(*models.UserWord) error
	GetAllByUser(*string) ([]*models.UserWord, error)
	DeleteUserWord(*string) error
}
