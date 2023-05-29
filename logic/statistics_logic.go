package logic

import (
	"damage/model"
	"damage/svc"
	"damage/types"
	"github.com/gin-gonic/gin"
	"go.mongodb.org/mongo-driver/mongo"
	"time"
)

func GetVisitsLogic(c *gin.Context) (visitsData types.GetVisitsResp, err error) {
	ctx := c.MustGet("ctx").(*svc.ServiceContext)
	todayAccess, err := ctx.AccessDayModel.GetTodayAccessNum(c.Request.Context())

	if err != nil {
		if err != mongo.ErrNoDocuments {
			return types.GetVisitsResp{}, err
		}
	}

	if todayAccess == nil {
		order := model.AccessDay{
			AccessNum: 0,
			CreatedAt: time.Now(),
			UpdatedAt: time.Now(),
		}
		_, err := ctx.AccessDayModel.CreateTodayAccessNum(c.Request.Context(), &order)
		if err != nil {
			return types.GetVisitsResp{}, err
		}
	} else {
		visitsData.AccessDay = todayAccess.AccessNum
	}

	total, err := ctx.AccessDayModel.GetAccessTotal(c.Request.Context())
	if err != nil {
		return types.GetVisitsResp{}, err
	}

	visitsData.AccessTotal = total

	return
}
