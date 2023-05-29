package model

import (
	"context"
	"go.mongodb.org/mongo-driver/bson"
	"go.mongodb.org/mongo-driver/bson/primitive"
	"go.mongodb.org/mongo-driver/mongo"
	"time"
)

type AccessDayModel struct {
	Mongo *mongo.Database
}

func NewAccessDayModel(mongo *mongo.Database) *AccessDayModel {
	return &AccessDayModel{Mongo: mongo}
}

type AccessDay struct {
	ID        primitive.ObjectID `bson:"_id,omitempty"`
	AccessNum int64              `bson:"access_num"`
	CreatedAt time.Time          `bson:"created_at"`
	UpdatedAt time.Time          `bson:"updated_at"`
}

func (m *AccessDayModel) CreateTodayAccessNum(ctx context.Context, order *AccessDay) (*string, error) {
	rst, err := m.Mongo.Collection("access_day").InsertOne(ctx, order)
	if err != nil {
		return nil, err
	}
	order.ID = rst.InsertedID.(primitive.ObjectID)
	id := order.ID.Hex()
	return &id, nil
}

func (m *AccessDayModel) UpdateTodayAccessNum(ctx context.Context, order *AccessDay) error {
	_, err := m.Mongo.Collection("access_day").ReplaceOne(ctx, bson.M{"_id": order.ID}, order)
	if err != nil {
		return err
	}
	return nil
}

func (m *AccessDayModel) GetTodayAccessNum(ctx context.Context) (*AccessDay, error) {
	now := time.Now()
	startOfDay := time.Date(now.Year(), now.Month(), now.Day(), 0, 0, 0, 0, now.Location())
	filter := bson.M{
		"created_at": bson.M{
			"$gte": startOfDay,
			"$lt":  startOfDay.Add(24 * time.Hour),
		},
	}

	rst := m.Mongo.Collection("access_day").FindOne(ctx, filter)
	if err := rst.Err(); err != nil {
		return nil, err
	}
	order := new(AccessDay)
	err := rst.Decode(order)
	if err != nil {
		return nil, err
	}
	return order, nil
}

func (m *AccessDayModel) GetAccessTotal(ctx context.Context) (int64, error) {
	collection := m.Mongo.Collection("access_day")

	pipeline := mongo.Pipeline{
		{{"$group", bson.D{{"_id", nil}, {"total", bson.D{{"$sum", "$access_num"}}}}}},
	}

	// 4. 执行聚合操作
	cursor, err := collection.Aggregate(ctx, pipeline)
	if err != nil {
		return 0, err
	}
	defer cursor.Close(ctx)

	var result struct {
		Total int `bson:"total"`
	}

	// 5. 输出聚合结果
	if cursor.Next(ctx) {

		err := cursor.Decode(&result)
		if err != nil {
			return 0, err
		}

	}

	return int64(result.Total), nil
}
