package types

type CharacterDataReq struct {
	Occupation             int64   `form:"occupation" binding:"required"`               //职业
	TestType               int64   `form:"test_type" binding:"required"`                //副本类型
	MinAttack              float64 `form:"min_attack" binding:"required"`               //最小攻击
	MaxAttack              float64 `form:"max_attack" binding:"required"`               //最大攻击
	ElementalDamage        float64 `form:"elemental_damage" binding:"required"`         //全元素攻击
	FixedDefeat            float64 `form:"fixed_defeat" `                               //克敌
	PercentageDefeat       float64 `form:"percentage_defeat" `                          //百分比克敌
	MonsterPenetration     float64 `form:"monster_penetration" binding:"required"`      //怪物穿透
	CriticalHitProbability float64 `form:"critical_hit_probability" binding:"required"` //会心
	CriticalHitDamage      float64 `form:"critical_hit_damage" binding:"required"`      //会心伤害
	FixedBrokenGuard       float64 `form:"fixed_broken_guard" `                         //固定破防
	PercentageBrokenGuard  float64 `form:"percentage_broken_guard"`                     //百分比破防
}

type CharacterDataResp struct {
	Code    int64            `json:"code"`
	Message string           `json:"message"`
	Data    SecondDamageData `json:"data"`
}

type SecondDamageData struct {
	SecondDamageSum string `json:"second_damage_sum"`
}

type GetCollocationReq struct {
	Occupation int64 `form:"occupation" binding:"required"` //职业
}

type GetCollocationResp struct {
	Code    int64           `json:"code"`
	Message string          `json:"message"`
	Data    CollocationData `json:"data"`
}

type CollocationData struct {
	Weapons   string `json:"weapons"`   //武器
	Necklace  string `json:"necklace"`  //项链
	Ring1     string `json:"ring1"`     //戒指1
	Ring2     string `json:"ring2"`     //戒指2
	Bracelet1 string `json:"bracelet1"` //手镯1
	Bracelet2 string `json:"bracelet2"` //手镯2
}

type GetVisitsResp struct {
	AccessTotal int64 `json:"access_total"`
	AccessDay   int64 `json:"access_day"`
	UserTotal   int64 `json:"user_total"`
}

type GetRankingReq struct {
	RoleUID     int64  `form:"role_uid"`
	RoleName    string `form:"role_name"`
	Occupation  int64  `form:"occupation"`
	AreaService string `form:"area_service"`
}

type GetRankingResp struct {
	Code    int64             `json:"code"`
	Message string            `json:"message"`
	Data    []RankingRespData `json:"data"`
}

type RankingRespData struct {
	ID          string  `json:"id"`
	RoleUID     int64   `json:"role_uid"`
	RoleName    string  `json:"role_name"`
	Occupation  int64   `json:"occupation"`
	AreaService string  `json:"area_service"`
	Damage      float64 `json:"damage"`
}

type UpdateRankingReq struct {
	ID          string  `form:"id"`
	RoleUID     int64   `form:"role_uid"`
	RoleName    string  `form:"role_name"`
	Occupation  int64   `form:"occupation"`
	AreaService string  `form:"area_service"`
	Damage      float64 `form:"damage"`
}

type UpdateRankingResp struct {
	Code    int64  `json:"code"`
	Message string `json:"message"`
}
