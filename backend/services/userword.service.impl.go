package services

import (
	"context"
	"errors"
	"vocabulary-builder/models"

	"github.com/google/uuid"
	"go.mongodb.org/mongo-driver/bson"
	"go.mongodb.org/mongo-driver/bson/primitive"
	"go.mongodb.org/mongo-driver/mongo"
)

type UserWordServiceImpl struct {
	UserWordCollection *mongo.Collection
	ctx                context.Context
}

func NewUserWordService(UserWordCollection *mongo.Collection, ctx context.Context) UserWordService {

	return &UserWordServiceImpl{
		UserWordCollection: UserWordCollection,
		ctx:                ctx,
	}
}

func (uws *UserWordServiceImpl) SaveUserWord(userword *models.UserWord) error {
	userword.UUID = uuid.New().String()
	_, err := uws.UserWordCollection.InsertOne(uws.ctx, userword)

	return err
}
func (uws *UserWordServiceImpl) UpdateUserWord(userword *models.UserWord) error {

	filter := bson.D{primitive.E{Key: "uuid", Value: userword.UUID}}

	update := bson.D{
		primitive.E{Key: "uuid", Value: userword.UUID},
		primitive.E{Key: "learned", Value: userword.Learned},
		primitive.E{Key: "userUUID", Value: userword.UserUUID},
		primitive.E{Key: "wordUUID", Value: userword.WordUUID},
	}
	result, _ := uws.UserWordCollection.ReplaceOne(uws.ctx, filter, update)
	if result.ModifiedCount == 0 {
		return errors.New("userword not updated")
	}
	return nil
}
func (uws *UserWordServiceImpl) GetAllByUser(uuid *string) ([]*models.UserWord, error) {

	filter := bson.D{primitive.E{Key: "uuid", Value: uuid}}
	cursor, err := uws.UserWordCollection.Find(uws.ctx, filter)

	if err != nil {
		return nil, err
	}

	var userWordList []*models.UserWord

	for cursor.Next(uws.ctx) {
		var userWord models.UserWord
		err := cursor.Decode(&userWord)
		if err != nil {
			return nil, err
		}
		userWordList = append(userWordList, &userWord)
	}
	if err := cursor.Err(); err != nil {
		return nil, err
	}
	cursor.Close(uws.ctx)

	if len(userWordList) == 0 {
		return nil, errors.New("userword not found")
	}

	return userWordList, nil
}

func (uws *UserWordServiceImpl) DeleteUserWord(uuid *string) error {
	filter := bson.D{primitive.E{Key: "uuid", Value: uuid}}
	result, _ := uws.UserWordCollection.DeleteOne(uws.ctx, filter)
	if result.DeletedCount == 0 {
		return errors.New("userword not deleted")
	}
	return nil
}
