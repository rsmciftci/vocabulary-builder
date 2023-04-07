package models

type Word struct {
	Word             *string   `json:"word" validate:"required" `
	Meaning          *string   `json:"meaning" validate:"required"`
	WordType         *string   `json:"word_type" validate:"required"`
	ExampleSentences []*string `json:"example_sentences"`
}
