package attribute_value

// BasicCharacter 人物基础属性
type BasicCharacter struct {
	MinAttack              float64 //最小攻击
	MaxAttack              float64 //最大攻击
	ElementalDamage        float64 //全元素攻击
	FixedDefeat            float64 //克敌
	PercentageDefeat       float64 //百分比克敌
	MonsterPenetration     float64 //怪物穿透
	CriticalHitProbability float64 //会心
	CriticalHitDamage      float64 //会心伤害
	FixedBrokenGuard       float64 //破防
	PercentageBrokenGuard  float64 //百分比破防
}
