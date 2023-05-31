package model

import (
	"context"
	"damage/types"
	"go.mongodb.org/mongo-driver/bson"
	"go.mongodb.org/mongo-driver/bson/primitive"
	"go.mongodb.org/mongo-driver/mongo"
	"go.mongodb.org/mongo-driver/mongo/options"
	"time"
)

type RankingRecordModel struct {
	Mongo *mongo.Database
}

func NewRankingRecordModel(mongo *mongo.Database) *RankingRecordModel {
	return &RankingRecordModel{Mongo: mongo}
}

type RankingRecord struct {
	ID          primitive.ObjectID `bson:"_id,omitempty"`
	RoleUID     int64              `bson:"role_uid"`
	RoleName    string             `bson:"role_name"`
	Occupation  int64              `bson:"occupation"`
	AreaService string             `bson:"area_service"`
	Damage      float64            `bson:"damage"`
	CreatedAt   time.Time          `bson:"created_at"`
	UpdatedAt   time.Time          `bson:"updated_at"`
}

func (r *RankingRecordModel) CreateRankingRecord(ctx context.Context, order *RankingRecord) (*string, error) {
	rst, err := r.Mongo.Collection("ranking").InsertOne(ctx, order)
	if err != nil {
		return nil, err
	}
	order.ID = rst.InsertedID.(primitive.ObjectID)
	id := order.ID.Hex()
	return &id, nil
}

func (r *RankingRecordModel) UpdateRankingRecord(ctx context.Context, order *RankingRecord) error {
	_, err := r.Mongo.Collection("ranking").ReplaceOne(ctx, bson.M{"_id": order.ID}, order)
	if err != nil {
		return err
	}
	return nil
}

func (r *RankingRecordModel) GetRankingRecord(ctx context.Context, ID primitive.ObjectID) (*RankingRecord, error) {
	rst := r.Mongo.Collection("ranking").FindOne(ctx, bson.M{"_id": ID})
	if err := rst.Err(); err != nil {
		return nil, err
	}
	order := new(RankingRecord)
	err := rst.Decode(order)
	if err != nil {
		return nil, err
	}
	return order, nil
}

func (r *RankingRecordModel) GetRankingRecordList(ctx context.Context, cmd *types.GetRankingReq) ([]*RankingRecord, error) {
	result := make([]*RankingRecord, 0, 10)
	coll := r.Mongo.Collection("ranking")
	opt := new(options.FindOptions)
	opt.SetLimit(10)

	sortMap := bson.D{{"damage", -1}}
	filter := bson.M{}

	_, err := coll.CountDocuments(ctx, filter)
	if err != nil {
		return nil, err
	}
	opt.SetSort(sortMap)
	cursor, err := coll.Find(ctx, filter, opt)
	if err != nil {
		return nil, err
	}
	for cursor.Next(ctx) {
		item := new(RankingRecord)
		err := cursor.Decode(item)
		if err != nil {
			return nil, err
		}
		result = append(result, item)
	}
	return result, nil

}
