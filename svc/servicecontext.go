package svc

import (
	"context"
	"damage/conf"
	"damage/model"
	"github.com/gin-gonic/gin"
	"go.mongodb.org/mongo-driver/mongo"
	"gorm.io/driver/mysql"
	"gorm.io/gorm"
	"log"
	"time"
)

const (
	DefaultMaxOpenCount = 100
	DefaultMaxIdleCount = 50
)

type ServiceContext struct {
	Mongo          *mongo.Database
	AccessDayModel *model.AccessDayModel
	RankingRecord  *model.RankingRecordModel
	UserInfoModel  *model.UserInfoModel
}

func NewServiceContext(c *gin.Context) *ServiceContext {
	ctx := context.Background()

	client := c.MustGet("dbClient").(*mongo.Client)
	err := client.Ping(ctx, nil)
	if err != nil {
		panic(err)
	}
	mongoDb := client.Database("expert")

	log.Println("mongoDb create success ")
	var mysqlConf conf.MysqlConfig
	mysqlConf.Datasource = "root:123456@tcp(localhost:3306)/assistant?charset=utf8mb4&parseTime=True&loc=Local"
	db := initDB(mysqlConf)
	err = db.AutoMigrate(
		&model.UserInfo{})
	if err != nil {
		log.Println(err)
		panic(err)
	}
	log.Println("mysql create success ")
	return &ServiceContext{
		Mongo:          mongoDb,
		AccessDayModel: model.NewAccessDayModel(mongoDb),
		RankingRecord:  model.NewRankingRecordModel(mongoDb),
		UserInfoModel:  model.NewUserInfoModel(db),
	}
}

func initDB(mysqlConf conf.MysqlConfig) *gorm.DB {
	db, err := gorm.Open(mysql.Open(mysqlConf.Datasource), &gorm.Config{})
	if err != nil {
		log.Println("无法连接到数据库：", err)
		panic(err)
	}
	sqlDB, err := db.DB()
	if err != nil {
		log.Println("无法连接到数据库：", err)
		panic(err)
	}
	log.Println("成功连接到数据库！")
	if mysqlConf.ConnMaxOpen != 0 {
		sqlDB.SetMaxOpenConns(mysqlConf.ConnMaxOpen)
	} else {
		sqlDB.SetMaxOpenConns(DefaultMaxOpenCount)
	}
	if mysqlConf.ConnMaxIdle != 0 {
		sqlDB.SetMaxIdleConns(mysqlConf.ConnMaxIdle)
	} else {
		sqlDB.SetMaxIdleConns(DefaultMaxIdleCount)
	}
	if mysqlConf.ConnMaxLifeTime != 0 {
		sqlDB.SetConnMaxLifetime(time.Second * time.Duration(mysqlConf.ConnMaxLifeTime))
	} else {
		sqlDB.SetConnMaxLifetime(time.Minute)
	}

	return db
}
