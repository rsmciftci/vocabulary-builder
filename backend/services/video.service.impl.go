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

type VideoServiceImpl struct {
	videoCollection *mongo.Collection
	ctx             context.Context
}

func NewVideoService(videoCollection *mongo.Collection, ctx context.Context) VideoService {
	return &VideoServiceImpl{
		videoCollection: videoCollection,
		ctx:             ctx,
	}
}

func (vs *VideoServiceImpl) SaveVideo(video *models.Video) error {
	_, err := vs.videoCollection.InsertOne(vs.ctx, video)
	return err
}
func (vs *VideoServiceImpl) DeleteVideo(id string) error {
	hexId := hex.EncodeToString([]byte(id))
	objectId, _ := primitive.ObjectIDFromHex(hexId)

	filter := bson.D{primitive.E{Key: "_id", Value: objectId}}
	_, err := vs.videoCollection.DeleteOne(vs.ctx, filter)
	return err
}
func (vs *VideoServiceImpl) UpdateVideo(videoDto dto.VideoDto) error {
	hexId := hex.EncodeToString([]byte(videoDto.ObjectId))
	objectId, _ := primitive.ObjectIDFromHex(hexId)
	filter := bson.D{primitive.E{Key: "_id", Value: objectId}}
	update := bson.D{
		primitive.E{Key: "name", Value: videoDto.Name},
		primitive.E{Key: "text", Value: videoDto.Text},
		primitive.E{Key: "video_type", Value: videoDto.VideoType},
		primitive.E{Key: "season", Value: videoDto.Season},
		primitive.E{Key: "episode", Value: videoDto.Episode},
	}
	result, _ := vs.videoCollection.UpdateOne(vs.ctx, filter, update)
	if result.UpsertedCount == 0 {
		return errors.New("video not updated")
	}

	return nil
}

func (vs *VideoServiceImpl) GetAllVideos() ([]*models.Video, error) {
	var videos []*models.Video
	filter := bson.D{{}}
	cursor, _ := vs.videoCollection.Find(vs.ctx, filter)

	for cursor.Next(vs.ctx) {
		var video models.Video
		err := cursor.Decode(&video)
		if err != nil {
			return nil, err
		}
		videos = append(videos, &video)
	}

	if err := cursor.Err(); err != nil {
		return nil, err
	}

	cursor.Close(vs.ctx)

	if len(videos) == 0 {
		return nil, errors.New("video not found")
	}

	return videos, nil
}

func (vs *VideoServiceImpl) FindAVideo(videoDto *dto.VideoDto) (*models.Video, error) {
	var video *models.Video
	// TODO: farkli kosullar icin farkli filtreler olustur
	// film ise season ve episodes olmamali
	// tvs ise season ve episode bilgisi kesinlikle olmali
	filter := bson.D{

		primitive.E{Key: "name", Value: videoDto.Name},
		primitive.E{Key: "text", Value: videoDto.Text},
		primitive.E{Key: "video_type", Value: videoDto.VideoType},
		primitive.E{Key: "season", Value: videoDto.Season},
		primitive.E{Key: "episode", Value: videoDto.Episode}}

	err := vs.videoCollection.FindOne(vs.ctx, filter).Decode(&video)

	return video, err

}
