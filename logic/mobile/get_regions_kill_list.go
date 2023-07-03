package mobile

import (
	"damage/define"
	"damage/svc"
	"damage/types"
	"github.com/gin-gonic/gin"
	"strconv"
)

func GetMobileRegionsKillList(c *gin.Context, req types.GetMobileRegionsKillListReq) (regionsKillList types.GetMobileRegionsKillListResp, err error) {
	ctx := c.MustGet("ctx").(*svc.ServiceContext)

	regionsKillListData, err := ctx.RegionsKillModel.GetRegionsKillList(c.Request.Context(), &req)
	if err != nil {
		return regionsKillList, err
	}

	data := make([]types.MobileRegionsKillListData, 0)
	for _, regionsKill := range regionsKillListData {
		regionsKillData := types.MobileRegionsKillListData{
			ID:               regionsKill.ID.Hex(),
			Name:             regionsKill.Name,
			Coord:            strconv.FormatInt(regionsKill.CoordX, 10) + "," + strconv.FormatInt(regionsKill.CoordY, 10),
			TriggerCondition: regionsKill.TriggerCondition,
			Location:         define.Location[regionsKill.Location],
		}
		data = append(data, regionsKillData)
	}

	regionsKillList.Data.Items = data
	regionsKillList.Code = 200

	return
}
