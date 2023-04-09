package models

type User struct {
	Name        *string `binding:"required"`
	Surname     *string `binding:"required"`
	DateOfBirth *string `binding:"required"` // TODO: current time or past
	Gender      *string `binding:"required"` //TODO: "male" or "female"
	Email       *string `bson:"_id" binding:"required,email"`
	Password    *string `binding:"required,min=6"` //  TODO: password md5 olarak tutulmali
}
