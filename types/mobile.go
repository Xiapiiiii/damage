package types

type GetMobileAdventureListReq struct {
	Name     string `form:"name,optional"`     // 名字
	Location string `form:"location,optional"` //地点
	Quality  string `form:"quality,optional"`  //品质
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
	Coord            string `json:"coord"`             //坐标
	TriggerNpc       string `json:"trigger_npc"`       //触发npc
	TriggerCondition string `json:"trigger_condition"` //触发条件
	Location         string `json:"location"`          //地点
	Quality          string `json:"quality"`           //品质
	Award            string `json:"award"`             //奖励
}

// 江湖绝学
type GetMobileRegionsKillListReq struct {
	Name     string `form:"name,optional"`     // 名字
	Location string `form:"location,optional"` //地点
}

type GetMobileRegionsKillListResp struct {
	Code int64                 `json:"code"`
	Data MobileRegionsKillList `json:"data"`
}

type MobileRegionsKillList struct {
	Items []MobileRegionsKillListData `json:"items"`
}

type MobileRegionsKillListData struct {
	ID               string `json:"id"`
	Name             string `json:"name"`              //名字
	Coord            string `json:"coord"`             //坐标
	TriggerNpc       string `json:"trigger_npc"`       //触发npc
	TriggerCondition string `json:"trigger_condition"` //触发条件
	Location         string `json:"location"`          //地点
}
