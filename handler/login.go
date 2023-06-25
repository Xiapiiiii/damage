package handler

import (
	"damage/logic"
	"damage/types"
	"errors"
	"github.com/gin-gonic/gin"
)

func CreateUserRegisterInfo(c *gin.Context) {
	var data types.UserRegisterInfoReq
	if err := c.ShouldBind(&data); err != nil {
		c.JSON(400, gin.H{"error": errors.New("提交数据错误!请检查")})
		return
	}

	record, err := logic.CreateUserRegister(c, data)
	if err != nil {
		c.JSON(500, gin.H{"error": err})
		return
	}

	c.JSON(200, record)

}
