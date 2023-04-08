package dto

type VideoDto struct {
	ObjectId  string  `json:"object_id"`
	Name      *string `json:"name" validate:"required" `
	Text      *string `json:"text" validate:"required"` // TODO: string mi byte array mi?
	VideoType *string `json:"video_type" validate:"required"`
	Season    *int    `json:"season"`
	Episode   *int    `json:"episode"`
}
