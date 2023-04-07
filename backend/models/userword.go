package models

type UserWord struct {
	UserId  string `json:"user_id" validate:"required" `
	WordId  string `json:"word_id" validate:"required"`
	Learned *bool  `json:"learned" validate:"required"`
}
