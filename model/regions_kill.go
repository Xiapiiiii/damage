package model

import (
	"context"
	"damage/define"
	"damage/types"
	"go.mongodb.org/mongo-driver/bson"
	"go.mongodb.org/mongo-driver/bson/primitive"
	"go.mongodb.org/mongo-driver/mongo"
	"go.mongodb.org/mongo-driver/mongo/options"
	"time"
)

type RegionsKillModel struct {
	Mongo *mongo.Database
}

func NewRegionsKillModel(mongo *mongo.Database) *RegionsKillModel {
	return &RegionsKillModel{Mongo: mongo}
}

type RegionsKill struct {
	ID               primitive.ObjectID `bson:"_id,omitempty"`
	Name             string             `bson:"name"`              //名字
	CoordX           int64              `bson:"coord_x"`           //x坐标
	CoordY           int64              `bson:"coord_y"`           //y坐标
	TriggerCondition string             `bson:"trigger_condition"` //触发条件
	Location         int64              `bson:"location" `         //地点
	CreatedAt        time.Time          `bson:"created_at"`
	UpdatedAt        time.Time          `bson:"updated_at"`
}

func (r *RegionsKillModel) CreateRegionsKill(ctx context.Context, order *RegionsKill) (*string, error) {
	rst, err := r.Mongo.Collection("regions_kill").InsertOne(ctx, order)
	if err != nil {
		return nil, err
	}
	order.ID = rst.InsertedID.(primitive.ObjectID)
	id := order.ID.Hex()
	return &id, nil
}

func (r *RegionsKillModel) UpdateRegionsKill(ctx context.Context, order *RegionsKill) error {
	_, err := r.Mongo.Collection("regions_kill").ReplaceOne(ctx, bson.M{"_id": order.ID}, order)
	if err != nil {
		return err
	}
	return nil
}

func (r *RegionsKillModel) GetRegionsKill(ctx context.Context, ID primitive.ObjectID) (*RegionsKill, error) {
	rst := r.Mongo.Collection("regions_kill").FindOne(ctx, bson.M{"_id": ID})
	if err := rst.Err(); err != nil {
		return nil, err
	}
	order := new(RegionsKill)
	err := rst.Decode(order)
	if err != nil {
		return nil, err
	}
	return order, nil
}

func (r *RegionsKillModel) GetRegionsKillByName(ctx context.Context, name string) (*RegionsKill, error) {
	rst := r.Mongo.Collection("regions_kill").FindOne(ctx, bson.M{"name": name})
	if err := rst.Err(); err != nil {
		return nil, err
	}
	order := new(RegionsKill)
	err := rst.Decode(order)
	if err != nil {
		return nil, err
	}
	return order, nil
}

func (r *RegionsKillModel) GetRegionsKillList(ctx context.Context, cmd *types.GetMobileRegionsKillListReq) ([]*RegionsKill, error) {
	result := make([]*RegionsKill, 0)
	coll := r.Mongo.Collection("regions_kill")
	opt := new(options.FindOptions)
	opt.SetLimit(500)

	sortMap := bson.D{{"location", 1}}
	filter := bson.M{}

	if cmd.Name != "" {
		filter["name"] = bson.M{"$regex": cmd.Name}
	}

	if cmd.Location != "" {
		filter["location"] = define.LocationS[cmd.Location]
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
		item := new(RegionsKill)
		err := cursor.Decode(item)
		if err != nil {
			return nil, err
		}
		result = append(result, item)
	}
	return result, nil

}
