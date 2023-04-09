package services

import (
	"context"
	"errors"
	"vocabulary-builder/dto"
	"vocabulary-builder/models"

	"go.mongodb.org/mongo-driver/bson"
	"go.mongodb.org/mongo-driver/bson/primitive"
	"go.mongodb.org/mongo-driver/mongo"
)

type UserServiceImpl struct {
	userCollection *mongo.Collection
	ctx            context.Context
}

func NewUserService(userCollection *mongo.Collection, ctx context.Context) UserService {
	return &UserServiceImpl{
		userCollection: userCollection,
		ctx:            ctx,
	}
}

func (UserService *UserServiceImpl) SaveUser(user *models.User) error {
	_, err := UserService.userCollection.InsertOne(UserService.ctx, user)
	return err
}
func (UserService *UserServiceImpl) DeleteByEmail(email *string) error {
	filter := bson.D{primitive.E{Key: "email", Value: email}}
	result, _ := UserService.userCollection.DeleteOne(UserService.ctx, filter)
	if result.DeletedCount == 1 {
		return errors.New("email didn't match with any user")
	}
	return nil
}
func (UserService *UserServiceImpl) UpdateUser(user *models.User) error {

	filter := bson.D{primitive.E{Key: "_id", Value: user.Email}}
	update := bson.D{
		primitive.E{Key: "name", Value: user.Name},
		primitive.E{Key: "surname", Value: user.Surname},
		primitive.E{Key: "DateOfBirth", Value: user.DateOfBirth},
		primitive.E{Key: "gender", Value: user.Gender},
		primitive.E{Key: "_id", Value: user.Email},
		primitive.E{Key: "password", Value: user.Password},
	}

	result, _ := UserService.userCollection.ReplaceOne(UserService.ctx, filter, update)

	if result.ModifiedCount == 0 {
		return errors.New("user not updated")
	}
	return nil
}

func (UserService *UserServiceImpl) FindUserByEmailAndPassword(emailAndPsswd *dto.EmailAndPassword) error {

	filter := bson.D{primitive.E{Key: "email", Value: emailAndPsswd.Email},
		primitive.E{Key: "password", Value: emailAndPsswd.Password}}
	err := UserService.userCollection.FindOne(UserService.ctx, filter).Err()
	return err
}
