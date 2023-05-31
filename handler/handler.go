package handler

import (
	"damage/logic"
	"damage/types"
	"errors"
	"github.com/gin-gonic/gin"
	"log"
)

func CalculatedDamage(c *gin.Context) {
	var data types.CharacterDataReq
	// 你可以使用 c.ShouldBind() 或者 c.Bind()，它们之间的区别是 ShouldBind() 不会写入响应体，而 Bind() 会。
	// 在这种情况下，我们只需要检查数据是否存在，所以我们使用 ShouldBind()。
	if err := c.ShouldBind(&data); err != nil {
		c.JSON(400, gin.H{"error": errors.New("提交数据错误!请检查")})
		return
	}
	log.Println(data)

	characterDataResp, err := logic.CalculatedDamageLogic(c, data)
	if err != nil {
		c.JSON(500, gin.H{"error": err})
		return
	}
	c.JSON(200, characterDataResp)
}

func GetCollocation(c *gin.Context) {
	var data types.GetCollocationReq
	// 你可以使用 c.ShouldBind() 或者 c.Bind()，它们之间的区别是 ShouldBind() 不会写入响应体，而 Bind() 会。
	// 在这种情况下，我们只需要检查数据是否存在，所以我们使用 ShouldBind()。
	if err := c.ShouldBind(&data); err != nil {
		c.JSON(400, gin.H{"error": err.Error()})
		return
	}
	log.Println(data)

	collocationDataResp := logic.GetCollocationLogic(data)

	c.JSON(200, collocationDataResp)
}

func GetVisits(c *gin.Context) {

	collocationDataResp, err := logic.GetVisitsLogic(c)
	if err != nil {
		c.JSON(500, gin.H{"error": err})
		return
	}
	c.JSON(200, collocationDataResp)
}

func GetRankingList(c *gin.Context) {
	var data types.GetRankingReq
	if err := c.ShouldBind(&data); err != nil {
		c.JSON(400, gin.H{"error": errors.New("提交数据错误!请检查")})
		return
	}
	log.Println(data)
	rankingResp, err := logic.GetRankingList(c, data)
	if err != nil {
		c.JSON(500, gin.H{"error": err})
		return
	}
	c.JSON(200, rankingResp)
}

func UpdateRankingRecord(c *gin.Context) {
	var data types.UpdateRankingReq
	if err := c.ShouldBind(&data); err != nil {
		c.JSON(400, gin.H{"error": errors.New("提交数据错误!请检查")})
		return
	}

	record, err := logic.UpdateRankingRecord(c, data)
	if err != nil {
		c.JSON(500, gin.H{"error": err})
		return
	}

	c.JSON(200, record)

}
