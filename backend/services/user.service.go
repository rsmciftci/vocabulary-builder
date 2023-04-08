package services

import (
	"vocabulary-builder/dto"
	"vocabulary-builder/models"
)

type UserService interface {
	SaveUser(*models.User) error
	DeleteByEmail(*string) error
	UpdateUser(*models.User) error
	FindUserByEmailAndPassword(*dto.EmailAndPassword) error
}
