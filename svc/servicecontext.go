package svc

import (
	"context"
	"damage/conf"
	"damage/model"
	"go.mongodb.org/mongo-driver/mongo"
	"go.mongodb.org/mongo-driver/mongo/options"
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
	Mongo            *mongo.Database
	AccessDayModel   *model.AccessDayModel
	RankingRecord    *model.RankingRecordModel
	UserInfoModel    *model.UserInfoModel
	AdventureModel   *model.AdventureModel
	RegionsKillModel *model.RegionsKillModel
}

func NewServiceContext() *ServiceContext {
	ctx := context.Background()

	client, err := connectToDB()
	if err != nil {
		log.Fatal(err)
	}
	err = client.Ping(ctx, nil)
	if err != nil {
		panic(err)
	}
	mongoDb := client.Database("expert")

	log.Println("mongoDb create success ")

	//var mysqlConf conf.MysqlConfig
	//mysqlConf.Datasource = "root:123456@tcp(127.0.0.1:3306)/assistant?charset=utf8mb4&parseTime=True&loc=Local"
	//db := initDB(mysqlConf)
	//err = db.AutoMigrate(
	//	&model.UserInfo{})
	//if err != nil {
	//	log.Println(err)
	//	panic(err)
	//}
	//log.Println("mysql create success ")

	return &ServiceContext{
		Mongo:          mongoDb,
		AccessDayModel: model.NewAccessDayModel(mongoDb),
		RankingRecord:  model.NewRankingRecordModel(mongoDb),
		//UserInfoModel:  model.NewUserInfoModel(db),
		AdventureModel:   model.NewAdventureModel(mongoDb),
		RegionsKillModel: model.NewRegionsKillModel(mongoDb),
	}
}

func initDB(mysqlConf conf.MysqlConfig) *gorm.DB {
	db, err := gorm.Open(mysql.Open(mysqlConf.Datasource), &gorm.Config{})
	if err != nil {
		log.Println("无法连接到数据库：", err)
		panic(err)
	}
	//sqlDB, err := db.DB()
	//if err != nil {
	//	log.Println("无法连接到数据库：", err)
	//	panic(err)
	//}
	//log.Println("成功连接到数据库！")
	//if mysqlConf.ConnMaxOpen != 0 {
	//	sqlDB.SetMaxOpenConns(mysqlConf.ConnMaxOpen)
	//} else {
	//	sqlDB.SetMaxOpenConns(DefaultMaxOpenCount)
	//}
	//if mysqlConf.ConnMaxIdle != 0 {
	//	sqlDB.SetMaxIdleConns(mysqlConf.ConnMaxIdle)
	//} else {
	//	sqlDB.SetMaxIdleConns(DefaultMaxIdleCount)
	//}
	//if mysqlConf.ConnMaxLifeTime != 0 {
	//	sqlDB.SetConnMaxLifetime(time.Second * time.Duration(mysqlConf.ConnMaxLifeTime))
	//} else {
	//	sqlDB.SetConnMaxLifetime(time.Minute)
	//}

	return db
}

func connectToDB() (*mongo.Client, error) {
	ctx, cancel := context.WithTimeout(context.Background(), 10*time.Second)
	defer cancel()

	client, err := mongo.Connect(ctx, options.Client().ApplyURI("mongodb://127.0.0.1:27017"))
	log.Println("mongo client success")
	if err != nil {
		return nil, err
	}

	return client, nil
}
