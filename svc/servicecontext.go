package svc

import (
	"context"
	"damage/model"
	"github.com/gin-gonic/gin"
	"go.mongodb.org/mongo-driver/mongo"
)

type ServiceContext struct {
	Mongo          *mongo.Database
	AccessDayModel *model.AccessDayModel
	RankingRecord  *model.RankingRecordModel
}

func NewServiceContext(c *gin.Context) *ServiceContext {
	ctx := context.Background()

	client := c.MustGet("dbClient").(*mongo.Client)
	err := client.Ping(ctx, nil)
	if err != nil {
		panic(err)
	}
	mongoDb := client.Database("expert")

	return &ServiceContext{
		Mongo:          mongoDb,
		AccessDayModel: model.NewAccessDayModel(mongoDb),
		RankingRecord:  model.NewRankingRecordModel(mongoDb),
	}
}
