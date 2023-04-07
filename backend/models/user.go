package models

import (
	"go.mongodb.org/mongo-driver/bson/primitive"
)

type User struct {
	Name        *string             `json:"name" validate:"required"`
	Surname     *string             `json:"surname" validate:"required"`
	DateOfBirth *primitive.DateTime `json:"date_of_birth" validate:"required,datetime"`
	Gender      *string             `json:"gender" validate:"required"`
	Email       *string             `json:"email" validate:"email,required"` // TODO: unique olmali, kayit sirasinda sorgula
	Password    *string             `json:"password" `
}
