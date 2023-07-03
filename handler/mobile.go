package handler

import (
	"damage/logic/mobile"
	"damage/types"
	"errors"
	"github.com/gin-gonic/gin"
)

func GetMobileAdventureList(c *gin.Context) {
	var data types.GetMobileAdventureListReq
	if err := c.ShouldBind(&data); err != nil {
		c.JSON(400, gin.H{"error": errors.New("提交数据错误!请检查")})
		return
	}

	record, err := mobile.GetMobileAdventureList(c, data)
	if err != nil {
		c.JSON(500, gin.H{"error": err})
		return
	}

	c.JSON(200, record)

}

func GetMobileRegionsKillList(c *gin.Context) {
	var data types.GetMobileRegionsKillListReq
	if err := c.ShouldBind(&data); err != nil {
		c.JSON(400, gin.H{"error": errors.New("提交数据错误!请检查")})
		return
	}

	record, err := mobile.GetMobileRegionsKillList(c, data)
	if err != nil {
		c.JSON(500, gin.H{"error": err})
		return
	}

	c.JSON(200, record)

}
