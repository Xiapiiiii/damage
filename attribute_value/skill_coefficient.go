package attribute_value

const (
	ImageOfGod         = iota + 1 //神像
	RiverOfBlood                  //血河
	DragonMinding                 //龙吟（重剑）
	DragonMindingLight            //龙吟（轻剑）
	SimpleQuestion                //素问（天问）
	NineSpirits                   //九灵
	TheMystery                    //玄机
	BrokenDreams                  //碎梦
	HighTone                      //鸿音
	GarmentOfIron                 //铁衣（破）
)

// PVEFixM PVE修正系数
var PVEFixM = map[int64]float64{
	ImageOfGod:         1.18,
	RiverOfBlood:       1.25,
	DragonMinding:      1.08,
	DragonMindingLight: 1.08,
	SimpleQuestion:     1.21,
	NineSpirits:        1.20,
	TheMystery:         1.27,
	BrokenDreams:       1.23,
	HighTone:           1.12,
	GarmentOfIron:      1 * 1.15 * 1.05,
}

// DefeatFixM 克敌系数修正
var DefeatFixM = map[int64]float64{
	ImageOfGod:         1,
	RiverOfBlood:       0.8,
	DragonMinding:      1,
	DragonMindingLight: 0.75 * 1.5 * 1.1,
	SimpleQuestion:     1,
	NineSpirits:        0.5 * 1.4,
	TheMystery:         0.47 * 1.4,
	BrokenDreams:       1 * 1.12, //已修正
	HighTone:           1,
	GarmentOfIron:      1 * 1.4, //已修正
}

// SkillBaseCoefficient 技能基础系数
type SkillBaseCoefficient struct {
	AttackCoefficient   float64 //攻击力系数
	ElementsCoefficient float64 //元素伤害系数
	ShootRate           float64 //攻击频率, 1为一秒一次
	FactorOfDefeat      float64 //克敌系数,默认1
}

func NewSkillCoefficient(occupation int) map[string]SkillBaseCoefficient {
	switch occupation {
	case ImageOfGod:
		return newImageOfGod()
	case RiverOfBlood:
		return newRiverOfBlood()
	case DragonMinding:
		return newDragonMinding()
	case DragonMindingLight:
		return newDragonMindingLight()
	case SimpleQuestion:
		return newSimpleQuestion()
	case NineSpirits:
		return newNineSpirits()
	case TheMystery:
		return newTheMystery()
	case BrokenDreams:
		return newBrokenDreams()
	case HighTone:
		return newHighTone()
	case GarmentOfIron:
		return newGarmentOfIron()
	default:
		return newImageOfGod()
	}
}

// 神像系数
func newImageOfGod() map[string]SkillBaseCoefficient {
	imageOfGod := make(map[string]SkillBaseCoefficient)
	imageOfGod["阳关三叠"] = SkillBaseCoefficient{
		AttackCoefficient:   0.57,
		ElementsCoefficient: 0.54,
		ShootRate:           1.29,
		FactorOfDefeat:      1,
	}
	imageOfGod["平沙落雁"] = SkillBaseCoefficient{
		AttackCoefficient:   1.74 * 1.1,
		ElementsCoefficient: 1.65 * 1.1,
		ShootRate:           0.154,
		FactorOfDefeat:      1.0,
	}
	imageOfGod["凤凰展翅"] = SkillBaseCoefficient{
		AttackCoefficient:   1.97 * 1.4,
		ElementsCoefficient: 1.83 * 1.4,
		ShootRate:           0.185,
		FactorOfDefeat:      1.0,
	}
	imageOfGod["灼烧"] = SkillBaseCoefficient{
		AttackCoefficient:   0.34,
		ElementsCoefficient: 0.32,
		ShootRate:           0.9,
		FactorOfDefeat:      1,
	}
	imageOfGod["百鸟朝凤"] = SkillBaseCoefficient{
		AttackCoefficient:   0.51,
		ElementsCoefficient: 0.48,
		ShootRate:           0.90,
		FactorOfDefeat:      1.0,
	}
	imageOfGod["羽碎"] = SkillBaseCoefficient{
		AttackCoefficient:   1.74,
		ElementsCoefficient: 1.8,
		ShootRate:           0.25,
		FactorOfDefeat:      1.0,
	}
	imageOfGod["剑胆琴心"] = SkillBaseCoefficient{
		AttackCoefficient:   1.77 * 1.4,
		ElementsCoefficient: 1.83 * 1.4,
		ShootRate:           0.20,
		FactorOfDefeat:      1.0,
	}
	imageOfGod["御风"] = SkillBaseCoefficient{
		AttackCoefficient:   1.55,
		ElementsCoefficient: 1.67,
		ShootRate:           0.32,
		FactorOfDefeat:      1,
	}
	imageOfGod["幽谷飞泉"] = SkillBaseCoefficient{
		AttackCoefficient:   0.27,
		ElementsCoefficient: 0.29,
		ShootRate:           0.85,
		FactorOfDefeat:      1,
	}
	imageOfGod["飞羽空蝉"] = SkillBaseCoefficient{
		AttackCoefficient:   0.46,
		ElementsCoefficient: 0.48,
		ShootRate:           0.68,
		FactorOfDefeat:      1,
	}
	imageOfGod["高山流水"] = SkillBaseCoefficient{
		AttackCoefficient:   1.39,
		ElementsCoefficient: 1.44,
		ShootRate:           0.06,
		FactorOfDefeat:      1.0,
	}

	return imageOfGod
}

// 血河系数
func newRiverOfBlood() map[string]SkillBaseCoefficient {
	riverOfBlood := make(map[string]SkillBaseCoefficient)
	riverOfBlood["刺"] = SkillBaseCoefficient{
		AttackCoefficient:   0.94 * 1.08,
		ElementsCoefficient: 0.95,
		ShootRate:           1.67,
		FactorOfDefeat:      1.0,
	}
	riverOfBlood["铁画银钩"] = SkillBaseCoefficient{
		AttackCoefficient:   4.34 * 1.08,
		ElementsCoefficient: 3.26,
		ShootRate:           0.245,
		FactorOfDefeat:      1.0,
	}
	riverOfBlood["熔岩霸气"] = SkillBaseCoefficient{
		AttackCoefficient:   1.4 * 1.08,
		ElementsCoefficient: 0.0,
		ShootRate:           0.52,
		FactorOfDefeat:      1.0,
	}
	riverOfBlood["红莲咆哮戟"] = SkillBaseCoefficient{
		AttackCoefficient:   1.84 * 1.08,
		ElementsCoefficient: 0.5,
		ShootRate:           0.32,
		FactorOfDefeat:      1.0,
	}
	riverOfBlood["挑"] = SkillBaseCoefficient{
		AttackCoefficient:   0.44 * 1.08,
		ElementsCoefficient: 0.44,
		ShootRate:           0.37,
		FactorOfDefeat:      1.0,
	}
	riverOfBlood["扫六合"] = SkillBaseCoefficient{
		AttackCoefficient:   1.07 * 1.08,
		ElementsCoefficient: 0.0,
		ShootRate:           0.164,
		FactorOfDefeat:      1.0,
	}
	riverOfBlood["丹心饮意"] = SkillBaseCoefficient{
		AttackCoefficient:   1.1 * 1.08,
		ElementsCoefficient: 1.1,
		ShootRate:           0.1,
		FactorOfDefeat:      1.0,
	}

	return riverOfBlood
}

// 龙吟（重剑）系数
func newDragonMinding() map[string]SkillBaseCoefficient {
	dragonMinding := make(map[string]SkillBaseCoefficient)
	dragonMinding["阳关三叠"] = SkillBaseCoefficient{
		AttackCoefficient:   0.57,
		ElementsCoefficient: 0.58,
		ShootRate:           1.17,
	}
	dragonMinding["平沙落雁"] = SkillBaseCoefficient{
		AttackCoefficient:   1.74,
		ElementsCoefficient: 1.85,
		ShootRate:           0.16,
	}
	dragonMinding["凤凰展翅"] = SkillBaseCoefficient{
		AttackCoefficient:   1.73,
		ElementsCoefficient: 1.51,
		ShootRate:           0.79,
	}
	dragonMinding["灼烧"] = SkillBaseCoefficient{
		AttackCoefficient:   0.34,
		ElementsCoefficient: 1.51,
		ShootRate:           0.33,
	}

	return dragonMinding
}

// 龙吟（轻剑）系数
func newDragonMindingLight() map[string]SkillBaseCoefficient {
	dragonMindingLight := make(map[string]SkillBaseCoefficient)
	dragonMindingLight["青龙怒"] = SkillBaseCoefficient{
		AttackCoefficient:   0.445 * 1.08,
		ElementsCoefficient: 0.425 * 1.08,
		ShootRate:           1.6,
		FactorOfDefeat:      1.0,
	}
	dragonMindingLight["青龙怒威"] = SkillBaseCoefficient{
		AttackCoefficient:   1.186 * 1.08,
		ElementsCoefficient: 1.133 * 1.08,
		ShootRate:           0.82,
		FactorOfDefeat:      1.0,
	}
	dragonMindingLight["剑意斩霜"] = SkillBaseCoefficient{
		AttackCoefficient:   1.24 * 1.08,
		ElementsCoefficient: 1.19 * 1.08,
		ShootRate:           0.656,
		FactorOfDefeat:      1.0,
	}
	dragonMindingLight["惊雷怒涛"] = SkillBaseCoefficient{
		AttackCoefficient:   1.099 * 1.08,
		ElementsCoefficient: 1.413 * 1.08,
		ShootRate:           0.506,
		FactorOfDefeat:      1.0,
	}
	dragonMindingLight["惊雷"] = SkillBaseCoefficient{
		AttackCoefficient:   0.725 * 1.08,
		ElementsCoefficient: 0.852 * 1.08,
		ShootRate:           0.507,
		FactorOfDefeat:      1.0,
	}
	dragonMindingLight["求败"] = SkillBaseCoefficient{
		AttackCoefficient:   4.65 * 1.08,
		ElementsCoefficient: 4.44 * 1.08,
		ShootRate:           0.09,
		FactorOfDefeat:      1.0,
	}
	dragonMindingLight["斩霜"] = SkillBaseCoefficient{
		AttackCoefficient:   0.37 * 1.08,
		ElementsCoefficient: 0.35 * 1.08,
		ShootRate:           0.672,
		FactorOfDefeat:      1.0,
	}
	dragonMindingLight["听雨"] = SkillBaseCoefficient{
		AttackCoefficient:   0.55 * 1.08,
		ElementsCoefficient: 0.52 * 1.08,
		ShootRate:           0.447,
		FactorOfDefeat:      1.0,
	}
	dragonMindingLight["诛邪"] = SkillBaseCoefficient{
		AttackCoefficient:   0.447 * 1.08,
		ElementsCoefficient: 0.44 * 1.08,
		ShootRate:           0.4,
		FactorOfDefeat:      1.0,
	}
	dragonMindingLight["雷龙天翔"] = SkillBaseCoefficient{
		AttackCoefficient:   1.13 * 1.08,
		ElementsCoefficient: 1.08 * 1.08,
		ShootRate:           0.18,
		FactorOfDefeat:      1.0,
	}
	dragonMindingLight["吟风"] = SkillBaseCoefficient{
		AttackCoefficient:   0.596 * 1.08,
		ElementsCoefficient: 0.57 * 1.08,
		ShootRate:           0.298,
		FactorOfDefeat:      1.0,
	}

	return dragonMindingLight
}

// 素问（天问）系数
func newSimpleQuestion() map[string]SkillBaseCoefficient {
	simpleQuestion := make(map[string]SkillBaseCoefficient)
	simpleQuestion["飘絮"] = SkillBaseCoefficient{
		AttackCoefficient:   0.49,
		ElementsCoefficient: 0.5,
		ShootRate:           1.45,
		FactorOfDefeat:      1.0,
	}
	simpleQuestion["惊羽"] = SkillBaseCoefficient{
		AttackCoefficient:   7.53,
		ElementsCoefficient: 8.10,
		ShootRate:           0.07,
		FactorOfDefeat:      1.0,
	}
	simpleQuestion["蝶恋花绮罗"] = SkillBaseCoefficient{
		AttackCoefficient:   3.54,
		ElementsCoefficient: 3.70,
		ShootRate:           0.11,
		FactorOfDefeat:      1.0,
	}
	simpleQuestion["花间"] = SkillBaseCoefficient{
		AttackCoefficient:   1.71,
		ElementsCoefficient: 1.78,
		ShootRate:           0.095,
		FactorOfDefeat:      1.0,
	}
	simpleQuestion["花间被动"] = SkillBaseCoefficient{
		AttackCoefficient:   0.47,
		ElementsCoefficient: 0.5,
		ShootRate:           0.9,
		FactorOfDefeat:      1.0,
	}
	simpleQuestion["翩舞"] = SkillBaseCoefficient{
		AttackCoefficient:   0.44,
		ElementsCoefficient: 0.47,
		ShootRate:           0.93,
		FactorOfDefeat:      1.0,
	}

	return simpleQuestion
}

// 鸿音系数
func newHighTone() map[string]SkillBaseCoefficient {
	highTone := make(map[string]SkillBaseCoefficient)
	highTone["鸣玉落"] = SkillBaseCoefficient{
		AttackCoefficient:   0.55 * 1.2,
		ElementsCoefficient: 0.57,
		ShootRate:           2.8 * 1.1,
		FactorOfDefeat:      1.0,
	}
	highTone["醉舞狂歌"] = SkillBaseCoefficient{
		AttackCoefficient:   0.463 * 1.2,
		ElementsCoefficient: 0.472,
		ShootRate:           1.43 * 1.1,
		FactorOfDefeat:      1.0,
	}
	//频率需继续测试
	highTone["霸王卸甲"] = SkillBaseCoefficient{
		AttackCoefficient:   0.98 * 1.2,
		ElementsCoefficient: 1.02,
		ShootRate:           0.55 * 1.1,
		FactorOfDefeat:      1.0,
	}
	highTone["霓裳惊鸿"] = SkillBaseCoefficient{
		AttackCoefficient:   1.49 * 1.2,
		ElementsCoefficient: 1.53,
		ShootRate:           0.23 * 1.1,
		FactorOfDefeat:      1.0,
	}
	highTone["金沙狂舞"] = SkillBaseCoefficient{
		AttackCoefficient:   2.1 * 1.2,
		ElementsCoefficient: 2.2,
		ShootRate:           0.155 * 1.1,
		FactorOfDefeat:      1.0,
	}
	highTone["十面埋伏"] = SkillBaseCoefficient{
		AttackCoefficient:   0.82 * 1.2,
		ElementsCoefficient: 0.85,
		ShootRate:           0.31 * 1.1,
		FactorOfDefeat:      1.0,
	}
	//highTone["九霄落影"] = SkillBaseCoefficient{
	//	AttackCoefficient:   4.36 * 1.15,
	//	ElementsCoefficient: 4.5,
	//	ShootRate:           0.025 * 1.2,
	//}

	return highTone
}

// 玄机系数
func newTheMystery() map[string]SkillBaseCoefficient {
	theMystery := make(map[string]SkillBaseCoefficient)
	theMystery["疾羽"] = SkillBaseCoefficient{
		AttackCoefficient:   0.17,
		ElementsCoefficient: 0.16,
		ShootRate:           3.82,
		FactorOfDefeat:      1,
	}
	theMystery["翼火"] = SkillBaseCoefficient{
		AttackCoefficient:   0.22,
		ElementsCoefficient: 0.20,
		ShootRate:           1.67,
		FactorOfDefeat:      1,
	}
	theMystery["归燕"] = SkillBaseCoefficient{
		AttackCoefficient:   0.336,
		ElementsCoefficient: 0.312,
		ShootRate:           1.63,
		FactorOfDefeat:      1,
	}
	theMystery["伏龙箭"] = SkillBaseCoefficient{
		AttackCoefficient:   1.54,
		ElementsCoefficient: 1.46,
		ShootRate:           0.276,
		FactorOfDefeat:      1.0,
	}
	theMystery["万象玄枢"] = SkillBaseCoefficient{
		AttackCoefficient:   0.86,
		ElementsCoefficient: 0.812,
		ShootRate:           0.6,
		FactorOfDefeat:      1.0,
	}
	theMystery["千击蝶"] = SkillBaseCoefficient{
		AttackCoefficient:   0.456,
		ElementsCoefficient: 0.437,
		ShootRate:           0.82,
		FactorOfDefeat:      1,
	}
	theMystery["猎隼"] = SkillBaseCoefficient{
		AttackCoefficient:   1.9,
		ElementsCoefficient: 1.82,
		ShootRate:           0.2,
		FactorOfDefeat:      1.0,
	}
	theMystery["空翔鹰"] = SkillBaseCoefficient{
		AttackCoefficient:   0.3,
		ElementsCoefficient: 0.293,
		ShootRate:           1,
		FactorOfDefeat:      1,
	}
	theMystery["星芒"] = SkillBaseCoefficient{
		AttackCoefficient:   0.58,
		ElementsCoefficient: 0.56,
		ShootRate:           0.3,
		FactorOfDefeat:      1,
	}

	return theMystery
}

// 碎梦系数
func newBrokenDreams() map[string]SkillBaseCoefficient {
	brokenDreams := make(map[string]SkillBaseCoefficient)
	brokenDreams["一梦千一"] = SkillBaseCoefficient{
		AttackCoefficient:   1.48 * 1.2,
		ElementsCoefficient: 1.386,
		ShootRate:           0.426 * 1.1,
		FactorOfDefeat:      1,
	}
	brokenDreams["百裂千击"] = SkillBaseCoefficient{
		AttackCoefficient:   0.308 * 1.2 * 1.1,
		ElementsCoefficient: 0.288,
		ShootRate:           1.441 * 0.8 * 1.07 * 1.1 * 1.25,
		FactorOfDefeat:      0.6,
	}

	brokenDreams["断肠"] = SkillBaseCoefficient{
		AttackCoefficient:   0.440 * 1.1 * 1.15,
		ElementsCoefficient: 0.416,
		ShootRate:           0.64 * 1.25 * 1.1,
		FactorOfDefeat:      0.75,
	}
	brokenDreams["荆轲现匕"] = SkillBaseCoefficient{
		AttackCoefficient:   2.62 * 1.15 * 1.1,
		ElementsCoefficient: 2.45,
		ShootRate:           0.156 * 1.25 * 1.1 * 0.75,
		FactorOfDefeat:      1,
	}
	brokenDreams["碧血丹青"] = SkillBaseCoefficient{
		AttackCoefficient:   0.638 * 1.1 * 1.1,
		ElementsCoefficient: 0.597,
		ShootRate:           0.516 * 1.1,
		FactorOfDefeat:      0.75,
	}
	brokenDreams["无尽闪"] = SkillBaseCoefficient{
		AttackCoefficient:   0.617 * 1.1 * 1.1,
		ElementsCoefficient: 0.577,
		ShootRate:           0.459 * 1.1,
		FactorOfDefeat:      0.9,
	}
	brokenDreams["碎影"] = SkillBaseCoefficient{
		AttackCoefficient:   0.74 * 1.1 * 1.1,
		ElementsCoefficient: 0.693,
		ShootRate:           0.40,
		FactorOfDefeat:      1,
	}
	brokenDreams["月影碎空"] = SkillBaseCoefficient{
		AttackCoefficient:   1.21 * 1.1 * 1.1,
		ElementsCoefficient: 1.155,
		ShootRate:           0.1393 * 1.25 * 1.1,
		FactorOfDefeat:      1,
	}
	brokenDreams["黯灭"] = SkillBaseCoefficient{
		AttackCoefficient:   1.047 * 1.1 * 1.1,
		ElementsCoefficient: 1,
		ShootRate:           0.115 * 1.25 * 1.1 * 1.65,
		FactorOfDefeat:      1,
	}
	brokenDreams["无尽灭"] = SkillBaseCoefficient{
		AttackCoefficient:   0.986 * 1.1 * 1.1,
		ElementsCoefficient: 0.924,
		ShootRate:           0.065 * 1.4 * 1.1,
		FactorOfDefeat:      1,
	}
	brokenDreams["十击"] = SkillBaseCoefficient{
		AttackCoefficient:   0.346 * 1.1 * 1.1,
		ElementsCoefficient: 0.324,
		ShootRate:           0.55 * 1.25 * 1.1,
		FactorOfDefeat:      0.65,
	}

	return brokenDreams
}

// 九灵系数(独登台)
func newNineSpirits() map[string]SkillBaseCoefficient {
	nineSpirits := make(map[string]SkillBaseCoefficient)
	nineSpirits["环灵决"] = SkillBaseCoefficient{
		AttackCoefficient:   0.961,
		ElementsCoefficient: 0.735,
		ShootRate:           0.85 * 1.25,
		FactorOfDefeat:      1.0,
	}
	nineSpirits["破梦"] = SkillBaseCoefficient{
		AttackCoefficient:   2.69,
		ElementsCoefficient: 2.57,
		ShootRate:           0.26 * 1.25,
		FactorOfDefeat:      1.0,
	}
	nineSpirits["灵犀三现"] = SkillBaseCoefficient{
		AttackCoefficient:   1.12,
		ElementsCoefficient: 1.07,
		ShootRate:           0.58 * 1.25,
		FactorOfDefeat:      1.0,
	}
	nineSpirits["蚀骨销魂"] = SkillBaseCoefficient{
		AttackCoefficient:   0.293,
		ElementsCoefficient: 0.26,
		ShootRate:           1.1 * 1.25,
		FactorOfDefeat:      1.0,
	}
	nineSpirits["黄泉之烬"] = SkillBaseCoefficient{
		AttackCoefficient:   0.16,
		ElementsCoefficient: 0.14,
		ShootRate:           0.9 * 1.25,
		FactorOfDefeat:      1.0,
	}
	nineSpirits["刺魂击"] = SkillBaseCoefficient{
		AttackCoefficient:   1.46,
		ElementsCoefficient: 1.43,
		ShootRate:           0.20 * 1.25,
		FactorOfDefeat:      1.0,
	}
	nineSpirits["长风散魂"] = SkillBaseCoefficient{
		AttackCoefficient:   1.73,
		ElementsCoefficient: 1.68,
		ShootRate:           0.068 * 1.25,
		FactorOfDefeat:      1.0,
	}
	nineSpirits["蛊种"] = SkillBaseCoefficient{
		AttackCoefficient:   0.247,
		ElementsCoefficient: 0.22,
		ShootRate:           0.45 * 1.25,
		FactorOfDefeat:      1.0,
	}

	return nineSpirits
}

// 铁衣(破)系数
func newGarmentOfIron() map[string]SkillBaseCoefficient {
	garmentOfIron := make(map[string]SkillBaseCoefficient)
	garmentOfIron["无影拳"] = SkillBaseCoefficient{
		AttackCoefficient:   0.599,
		ElementsCoefficient: 0.556,
		ShootRate:           0.95,
		FactorOfDefeat:      0.75,
	}
	garmentOfIron["练火拳"] = SkillBaseCoefficient{
		AttackCoefficient:   2.76,
		ElementsCoefficient: 3.84,
		ShootRate:           0.182,
		FactorOfDefeat:      1.0,
	}
	garmentOfIron["余震"] = SkillBaseCoefficient{
		AttackCoefficient:   0.429,
		ElementsCoefficient: 0.4,
		ShootRate:           1,
		FactorOfDefeat:      1,
	}
	garmentOfIron["劈云拳"] = SkillBaseCoefficient{
		AttackCoefficient:   0.479,
		ElementsCoefficient: 0.444,
		ShootRate:           0.273 * 3,
		FactorOfDefeat:      1.0,
	}
	garmentOfIron["旋风腿"] = SkillBaseCoefficient{
		AttackCoefficient:   1.38,
		ElementsCoefficient: 1.28,
		ShootRate:           0.30,
		FactorOfDefeat:      1.0,
	}
	garmentOfIron["开山拳"] = SkillBaseCoefficient{
		AttackCoefficient:   4.37,
		ElementsCoefficient: 4.06,
		ShootRate:           0.1,
		FactorOfDefeat:      1.0,
	}
	garmentOfIron["一以贯之"] = SkillBaseCoefficient{
		AttackCoefficient:   3.355 * 1.4,
		ElementsCoefficient: 3.11 * 1.4,
		ShootRate:           0.075,
		FactorOfDefeat:      1.0,
	}
	garmentOfIron["冲劲寸拳"] = SkillBaseCoefficient{
		AttackCoefficient:   0.0,
		ElementsCoefficient: 0.0,
		ShootRate:           0.0,
		FactorOfDefeat:      1.0,
	}
	garmentOfIron["金刚拳"] = SkillBaseCoefficient{
		AttackCoefficient:   0.0,
		ElementsCoefficient: 0.0,
		ShootRate:           0.0,
		FactorOfDefeat:      1.0,
	}

	return garmentOfIron
}
