package logic

import (
	"damage/model"
	"damage/svc"
	"damage/types"
	"fmt"
	"github.com/gin-gonic/gin"
	"log"
	"net/http"
	"time"
)

func CreateUserRegister(c *gin.Context, userRegisterInfo types.UserRegisterInfoReq) (commonResp types.CommonResp, err error) {
	ctx := c.MustGet("ctx").(*svc.ServiceContext)
	userIp := c.ClientIP()

	fmt.Println(userRegisterInfo)
	fmt.Println(userIp)

	userInfo := model.UserInfo{
		UserName:    userRegisterInfo.UserName,
		UserId:      userRegisterInfo.UserId,
		PassWord:    userRegisterInfo.PassWord,
		Avatar:      "",
		Email:       userRegisterInfo.Email,
		Level:       1,
		LevelName:   "初入江湖",
		Phone:       0,
		LastLoginAt: time.Now(),
		LastLoginIp: userIp,
		CreatedIp:   userIp,
		CreatedAt:   time.Now(),
		UpdatedAt:   time.Now(),
	}
	fmt.Println(userInfo)
	err = ctx.UserInfoModel.CreateUserInfo(userInfo)
	if err != nil {
		log.Println(err)
		return types.CommonResp{}, err
	}

	return types.CommonResp{
		Code:    http.StatusOK,
		Message: "success",
	}, nil
}
