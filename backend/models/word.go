package models

type Word struct {
	UUID             string
	Word             *string   `json:"word" validate:"required" `
	Meaning          *string   `json:"meaning" validate:"required"`
	WordType         *string   `json:"word_type" bson:"word_type" validate:"required"`
	ExampleSentences []*string `json:"example_sentences" bson:"example_sentences"`
}
