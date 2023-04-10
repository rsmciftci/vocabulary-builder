package models

type User struct {
	UUID        string
	Name        *string `binding:"required"`
	Surname     *string `binding:"required"`
	DateOfBirth *string `binding:"required"` // TODO: validate current time or past
	Gender      *string `binding:"required"` //TODO: validate gender
	Email       *string `binding:"required,email"`
	Password    *string `binding:"required,min=6"` //  TODO: password md5 olarak tutulmali
}
