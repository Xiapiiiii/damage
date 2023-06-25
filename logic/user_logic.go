package logic

import (
	"damage/types"
	"fmt"
	"github.com/gin-gonic/gin"
)

func CreateUserRegister(c *gin.Context, userRegisterInfo types.UserRegisterInfoReq) (commonResp types.CommonResp, err error) {
	userIp := c.ClientIP()

	fmt.Println(userIp)
	return
}
