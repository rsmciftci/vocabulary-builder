package services

import (
	"vocabulary-builder/dto"
	"vocabulary-builder/models"
)

type WordService interface {
	SaveWord(*models.Word) error
	DeleteWordById(string) error
	UpdateWord(*dto.Word) error
	FindWord(*string) (*models.Word, error)
}
