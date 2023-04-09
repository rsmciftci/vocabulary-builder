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
	_, err := uws.UserWordCollection.InsertOne(uws.ctx, userword)

	return err
}
func (uws *UserWordServiceImpl) UpdateUserWord(userworddto *dto.UserWordDto) error {
	userWordId := userworddto.ObjectId
	hexId := hex.EncodeToString([]byte(userWordId))
	objectId, _ := primitive.ObjectIDFromHex(hexId)
	filter := bson.D{primitive.E{Key: "_id", Value: objectId}}

	update := bson.D{
		primitive.E{Key: "learned", Value: userworddto.Learned},
		primitive.E{Key: "userId", Value: userworddto.UserId},
		primitive.E{Key: "wordId", Value: userworddto.WordId},
	}
	result, _ := uws.UserWordCollection.UpdateOne(uws.ctx, filter, update)
	if result.UpsertedCount == 0 {
		return errors.New("userword not updated")
	}
	return nil
}
func (uws *UserWordServiceImpl) GetAllByUser(userId *string) ([]*models.UserWord, error) {

	filter := bson.D{primitive.E{Key: "user_id", Value: userId}}
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

func (uws *UserWordServiceImpl) DeleteUserWord(wordId *string) error {
	filter := bson.D{primitive.E{Key: "word_id", Value: wordId}}
	result, _ := uws.UserWordCollection.DeleteOne(uws.ctx, filter)
	if result.DeletedCount == 0 {
		return errors.New("userword not deleted")
	}
	return nil
}
