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

type AdventureModel struct {
	Mongo *mongo.Database
}

func NewAdventureModel(mongo *mongo.Database) *AdventureModel {
	return &AdventureModel{Mongo: mongo}
}

type Adventure struct {
	ID               primitive.ObjectID `bson:"_id,omitempty"`
	Name             string             `bson:"name"`              //名字
	CoordX           int64              `bson:"coord_x"`           //x坐标
	CoordY           int64              `bson:"coord_y"`           //y坐标
	TriggerNpc       string             `bson:"trigger_npc"`       //触发npc
	TriggerCondition string             `bson:"trigger_condition"` //触发条件
	Location         int64              `bson:"location" `         //地点
	Quality          int64              `bson:"quality" `          //品质
	Award            string             `bson:"award" `            //奖励
	CreatedAt        time.Time          `bson:"created_at"`
	UpdatedAt        time.Time          `bson:"updated_at"`
}

func (r *AdventureModel) CreateAdventure(ctx context.Context, order *Adventure) (*string, error) {
	rst, err := r.Mongo.Collection("adventure").InsertOne(ctx, order)
	if err != nil {
		return nil, err
	}
	order.ID = rst.InsertedID.(primitive.ObjectID)
	id := order.ID.Hex()
	return &id, nil
}

func (r *AdventureModel) UpdateAdventure(ctx context.Context, order *Adventure) error {
	_, err := r.Mongo.Collection("adventure").ReplaceOne(ctx, bson.M{"_id": order.ID}, order)
	if err != nil {
		return err
	}
	return nil
}

func (r *AdventureModel) GetAdventure(ctx context.Context, ID primitive.ObjectID) (*Adventure, error) {
	rst := r.Mongo.Collection("adventure").FindOne(ctx, bson.M{"_id": ID})
	if err := rst.Err(); err != nil {
		return nil, err
	}
	order := new(Adventure)
	err := rst.Decode(order)
	if err != nil {
		return nil, err
	}
	return order, nil
}

func (r *AdventureModel) GetAdventureList(ctx context.Context, cmd *types.GetMobileAdventureListReq) ([]*Adventure, error) {
	result := make([]*Adventure, 0)
	coll := r.Mongo.Collection("adventure")
	opt := new(options.FindOptions)
	opt.SetLimit(10)

	sortMap := bson.D{{"quality", 1}}
	filter := bson.M{}

	if cmd.Name != "" {
		filter["name"] = bson.M{"$regex": cmd.Name}
	}

	if cmd.Quality != 0 {
		filter["quality"] = cmd.Quality
	}

	if cmd.Location != 0 {
		filter["location"] = cmd.Location
	}

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
		item := new(Adventure)
		err := cursor.Decode(item)
		if err != nil {
			return nil, err
		}
		result = append(result, item)
	}
	return result, nil

}
