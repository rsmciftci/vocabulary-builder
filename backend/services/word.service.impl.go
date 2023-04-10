package services

import (
	"context"
	"errors"
	"vocabulary-builder/dto"
	"vocabulary-builder/models"

	"github.com/google/uuid"
	"go.mongodb.org/mongo-driver/bson"
	"go.mongodb.org/mongo-driver/bson/primitive"
	"go.mongodb.org/mongo-driver/mongo"
)

type WordServiceImpl struct {
	wordColection *mongo.Collection
	ctx           context.Context
}

func NewWordService(wordCollection *mongo.Collection, ctx context.Context) WordService {
	return &WordServiceImpl{
		wordColection: wordCollection,
		ctx:           ctx,
	}
}

func (w *WordServiceImpl) SaveWord(word *models.Word) error {
	word.UUID = uuid.New().String()
	_, err := w.wordColection.InsertOne(w.ctx, word)
	if err != nil {
		return errors.New("word not saved")
	}
	return nil

}

func (w *WordServiceImpl) DeleteWordById(uuid string) error {
	filter := bson.D{primitive.E{Key: "uuid", Value: uuid}}
	result, _ := w.wordColection.DeleteOne(w.ctx, filter)
	if result.DeletedCount == 0 {
		return errors.New("word not deleted")
	}
	return nil
}

func (w *WordServiceImpl) UpdateWord(word *dto.Word) error {

	updated := bson.D{
		primitive.E{Key: "uuid", Value: word.UUID},
		primitive.E{Key: "word", Value: word.Word},
		primitive.E{Key: "meaning", Value: word.Meaning},
		primitive.E{Key: "word_type", Value: word.WordType},
		primitive.E{Key: "example_sentences", Value: word.ExampleSentences},
	}

	filter := bson.D{primitive.E{Key: "uuid", Value: word.UUID}}
	result, _ := w.wordColection.ReplaceOne(w.ctx, filter, updated)
	if result.MatchedCount == 0 {
		return errors.New("uuid didn't match with any words")
	}
	return nil
}

func (w *WordServiceImpl) FindWord(word *string) (*models.Word, error) {
	var wordModel *models.Word
	filter := bson.D{primitive.E{Key: "word", Value: word}}
	err := w.wordColection.FindOne(w.ctx, filter).Decode(&wordModel)
	return wordModel, err

}
