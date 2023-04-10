package models

type UserWord struct {
	UUID     string
	UserUUID string `json:"user_uuid" validate:"required" `
	WordUUID string `json:"word_uuid" validate:"required"`
	Learned  *bool  `json:"learned" validate:"required"`
}
