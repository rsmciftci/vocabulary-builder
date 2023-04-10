package main

import (
	"context"
	"fmt"
	"log"
	"vocabulary-builder/controllers"
	"vocabulary-builder/services"

	"github.com/gin-gonic/gin"
	"go.mongodb.org/mongo-driver/mongo"
	"go.mongodb.org/mongo-driver/mongo/options"
	"go.mongodb.org/mongo-driver/mongo/readpref"
)

// Replace the placeholder with your Atlas connection string
const uri = "mongodb://localhost:27017"

var (
	server             *gin.Engine
	us                 services.UserService
	uc                 controllers.UserController
	uws                services.UserWordService
	uwc                controllers.UserWordController
	ws                 services.WordService
	wc                 controllers.WordController
	vs                 services.VideoService
	vc                 controllers.VideoController
	ctx                context.Context
	userCollection     *mongo.Collection
	userWordCollection *mongo.Collection
	wordCollection     *mongo.Collection
	videoCollection    *mongo.Collection

	mongoClient *mongo.Client
	err         error
)

func init() {
	ctx = context.TODO()

	mongoConnection := options.Client().ApplyURI(uri)
	mongoClient, err := mongo.Connect(ctx, mongoConnection)
	if err != nil {
		log.Fatal("error while connecting to mongo: ", err)
	}
	err = mongoClient.Ping(ctx, readpref.Primary())
	if err != nil {
		log.Fatal("error while trying to ping mongo: ", err)
	}

	fmt.Println("mongo connection established")

	userCollection = mongoClient.Database("vocabularybuilder").Collection("users")
	us = services.NewUserService(userCollection, ctx)
	uc = controllers.NewUserController(us)

	userWordCollection = mongoClient.Database("vocabularybuilder").Collection("userwords")
	uws = services.NewUserWordService(userWordCollection, ctx)
	uwc = controllers.NewUserWordController(uws)

	wordCollection = mongoClient.Database("vocabularybuilder").Collection("words")
	ws = services.NewWordService(wordCollection, ctx)
	wc = controllers.NewWordController(ws)

	videoCollection = mongoClient.Database("vocabularybuilder").Collection("videos")
	vs = services.NewVideoService(videoCollection, ctx)
	vc = controllers.NewVideoController(vs)

	server = gin.Default()
}

func main() {
	defer mongoClient.Disconnect(ctx)

	basepath := server.Group("/v1")
	uc.RegisterUserRoutes(basepath)
	uwc.RegisterUserRoutes(basepath)
	wc.RegisterUserRoutes(basepath)
	vc.RegisterVideoRoutes(basepath)

	log.Fatal(server.Run(":9090"))

}

// TODO: jwt
// TODO: input sanitise inputs against sql injection
