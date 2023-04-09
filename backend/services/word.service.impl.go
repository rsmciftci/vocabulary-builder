package services

import (
	"context"
	"encoding/hex"
	"errors"
	"vocabulary-builder/dto"
	"vocabulary-builder/models"

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
	_, err := w.wordColection.InsertOne(w.ctx, word)
	if err != nil {
		return errors.New("word not saved")
	}
	return nil

}

func (w *WordServiceImpl) DeleteWordById(id string) error {
	hexId := hex.EncodeToString([]byte(id))
	objectId, _ := primitive.ObjectIDFromHex(hexId)
	filter := bson.D{primitive.E{Key: "_id", Value: objectId}}
	_, err := w.wordColection.DeleteOne(w.ctx, filter)
	return err
}

func (w *WordServiceImpl) UpdateWord(word *dto.Word) error {
	hexId := hex.EncodeToString([]byte(word.ObjectId))
	objectId, _ := primitive.ObjectIDFromHex(hexId)
	updated := bson.D{
		primitive.E{Key: "word", Value: word.Word},
		primitive.E{Key: "meaning", Value: word.Meaning},
		primitive.E{Key: "word_type", Value: word.WordType},
		primitive.E{Key: "example_sentences", Value: word.ExampleSentences},
	}

	filter := bson.D{primitive.E{Key: "object_id", Value: objectId}}
	result, _ := w.wordColection.UpdateOne(w.ctx, filter, updated)
	if result.UpsertedCount == 0 {
		return errors.New("word not updated")
	}
	return nil
}

func (w *WordServiceImpl) FindWord(word *string) (*models.Word, error) {
	var wordModel *models.Word
	filter := bson.D{primitive.E{Key: "word", Value: word}}
	err := w.wordColection.FindOne(w.ctx, filter).Decode(&wordModel)
	return wordModel, err

}
