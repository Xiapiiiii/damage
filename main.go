package main

import (
	"damage/handler"
	"damage/svc"
	"github.com/gin-gonic/gin"
)

func main() {

	r := gin.Default()

	r.Use(MyMiddleware())

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

	//生活技能
	r.POST("/life/revenue", handler.GetLifeRevenue)

	//注册
	r.POST("/user/register", handler.CreateUserRegisterInfo)
	//注册校验
	r.POST("/user/register/verify")
	//登录
	r.POST("/user/login")

	//批量上传奇遇excel
	r.POST("/upload/adventure", handler.UploadAdventureExcel)

	//获取奇遇列表
	r.POST("/mobile/adventure/list", handler.GetMobileAdventureList)
	//获取奇遇列表
	r.POST("api/mobile/adventure/list", handler.GetMobileAdventureList)

	//获取单个奇遇详情
	r.POST("/mobile/adventure")

	//批量上传绝学excel
	r.POST("/upload/regions-kill", handler.UploadRegionsKillExcel)

	//获取江湖绝学列表
	r.POST("/mobile/regions-kill/list", handler.GetMobileRegionsKillList)
	r.POST("api/mobile/regions-kill/list", handler.GetMobileRegionsKillList)

	r.Run(":8080")
}

func MyMiddleware() gin.HandlerFunc {
	ctx := svc.NewServiceContext()
	return func(c *gin.Context) {
		// 将数据库连接存储在上下文中
		c.Set("ctx", ctx)
		c.Next()
	}
}
