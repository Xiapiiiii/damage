package main

import (
	"context"
	"damage/handler"
	"damage/svc"
	"github.com/gin-gonic/gin"
	"go.mongodb.org/mongo-driver/mongo"
	"go.mongodb.org/mongo-driver/mongo/options"
	"log"
	"time"
)

func main() {
	client, err := connectToDB()
	if err != nil {
		log.Fatal(err)
	}

	r := gin.Default()

	r.Use(func(c *gin.Context) {
		c.Set("dbClient", client)
		c.Next()
	})

	r.Use(func(c *gin.Context) {
		ctx := svc.NewServiceContext(c)
		c.Set("ctx", ctx)
		c.Next()
	})

	//秒伤数据接口
	r.POST("/form", handler.CalculatedDamage)
	r.POST("api/form", handler.CalculatedDamage)

	//获取最佳搭配接口
	r.POST("/collocation", handler.GetCollocation)

	//获取总访问量、当日访问、总用户量
	r.POST("/visits", handler.GetVisits)
	//r.POST("api/visits", handler.GetVisits)

	//获取排行榜
	r.POST("/ranking/list", handler.GetRankingList)
	r.POST("api/ranking/list", handler.GetRankingList)
	//修改排行榜记录
	r.POST("/ranking/update", handler.UpdateRankingRecord)

	r.Run(":8080")
}

func connectToDB() (*mongo.Client, error) {
	ctx, cancel := context.WithTimeout(context.Background(), 10*time.Second)
	defer cancel()

	client, err := mongo.Connect(ctx, options.Client().ApplyURI("mongodb://localhost:27017"))
	log.Println("mongo client success")
	if err != nil {
		return nil, err
	}

	return client, nil
}
