package mobile

import (
	"damage/svc"
	"damage/types"
	"github.com/gin-gonic/gin"
)

func GetMobileAdventureList(c *gin.Context, req types.GetMobileAdventureListReq) (adventureList types.GetMobileAdventureListResp, err error) {
	ctx := c.MustGet("ctx").(*svc.ServiceContext)

	//t := &model.Adventure{
	//	Name:             "失孤",
	//	CoordX:           1008,
	//	CoordY:           141,
	//	TriggerNpc:       "王婆婆",
	//	TriggerCondition: "动作观察",
	//	Location:         1,
	//	Quality:          1,
	//	Award:            "",
	//	CreatedAt:        time.Now(),
	//	UpdatedAt:        time.Now(),
	//}
	//
	//_, err = ctx.AdventureModel.CreateAdventure(c.Request.Context(), t)
	//if err != nil {
	//	return types.GetMobileAdventureListResp{}, err
	//}

	adventureListData, err := ctx.AdventureModel.GetAdventureList(c.Request.Context(), &req)
	if err != nil {
		return adventureList, err
	}

	data := make([]types.MobileAdventureListData, 0)
	for _, adventure := range adventureListData {
		adventureData := types.MobileAdventureListData{
			ID:               adventure.ID.Hex(),
			Name:             adventure.Name,
			CoordX:           adventure.CoordX,
			CoordY:           adventure.CoordY,
			TriggerCondition: adventure.TriggerCondition,
			TriggerNpc:       adventure.TriggerNpc,
			Location:         adventure.Location,
			Quality:          adventure.Quality,
			Award:            adventure.Award,
		}
		data = append(data, adventureData)
	}

	adventureList.Data.Items = data
	adventureList.Code = 200

	return
}
