package types

type GetMobileAdventureListReq struct {
	Name     string `form:"name,optional"`     // 名字
	Location int64  `form:"location,optional"` //地点
	Quality  int64  `form:"quality,optional"`  //品质
}

type GetMobileAdventureListResp struct {
	Code int64               `json:"code"`
	Data MobileAdventureList `json:"data"`
}

type MobileAdventureList struct {
	Items []MobileAdventureListData `json:"items"`
}

type MobileAdventureListData struct {
	ID               string `json:"id"`
	Name             string `json:"name"`              //名字
	CoordX           int64  `json:"coord_x"`           //x坐标
	CoordY           int64  `json:"coord_y"`           //y坐标
	TriggerNpc       string `json:"trigger_npc"`       //触发npc
	TriggerCondition string `json:"trigger_condition"` //触发条件
	Location         int64  `json:"location"`          //地点
	Quality          int64  `json:"quality"`           //品质
	Award            string `json:"award"`             //奖励
}
