package services

import (
	"vocabulary-builder/dto"
	"vocabulary-builder/models"
)

type UserService interface {
	SaveUser(*models.User) error
	DeleteByUUID(*string) error
	UpdateUser(*models.User) error
	FindUserByEmailAndPassword(*dto.EmailAndPassword) error
}
