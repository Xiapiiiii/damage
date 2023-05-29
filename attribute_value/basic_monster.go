package attribute_value

const (
	CommonWhiteStone = iota + 1 //普通白石
	HeroWhiteStone              //英雄白石
	EliteWhiteStone             //绝世白石
)

// BasicMonster 怪物基础属性
type BasicMonster struct {
	Defense                  float64 //防御力（暂不分内外）
	CriticalStrikeResistance float64 //暴击抵抗
	MonsterResistance        float64 //怪物抗性
}

func NewBasicMonster(bizType int64) *BasicMonster {
	switch bizType {
	case CommonWhiteStone:
		return &BasicMonster{
			Defense:                  2888,
			CriticalStrikeResistance: 0.10,
			MonsterResistance:        1200,
		}
	case HeroWhiteStone:
		return &BasicMonster{
			Defense:                  3588,
			CriticalStrikeResistance: 0.20,
			MonsterResistance:        1480,
		}
	case EliteWhiteStone:
		return &BasicMonster{
			Defense:                  4400,
			CriticalStrikeResistance: 0.25,
			MonsterResistance:        1800,
		}
	default:
		return &BasicMonster{
			Defense:                  2888,
			CriticalStrikeResistance: 0.10,
			MonsterResistance:        1200,
		}
	}
}
