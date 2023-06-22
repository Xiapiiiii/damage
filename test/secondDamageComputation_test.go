package test

import (
	"damage/attribute_value"
	"damage/damage"
	"damage/types"
	"fmt"
	"testing"
)

// 玄机 元素 绝世3.3
func Test_SecondDamageComputation(t *testing.T) {
	skillCoefficients := attribute_value.NewSkillCoefficient(attribute_value.TheMystery)
	character := attribute_value.BasicCharacter{
		MinAttack:              3545,
		MaxAttack:              5017,
		ElementalDamage:        2943,
		FixedDefeat:            761,
		PercentageDefeat:       0.241,
		MonsterPenetration:     1448,
		CriticalHitProbability: 278,
		CriticalHitDamage:      0.50,
		FixedBrokenGuard:       1024,
		PercentageBrokenGuard:  0.0,
	}

	req := types.CharacterDataReq{
		TestType:   3,
		Occupation: attribute_value.TheMystery,
	}

	secondDamage := damage.FinalSecondDamageComputation(req, character, skillCoefficients)

	sum := 0.0
	for _, v := range secondDamage {
		sum += v
	}
	fmt.Printf("%f万 \n", sum/10000)

	for k, v := range secondDamage {
		fmt.Printf("技能:%s 秒伤:%f 万  占比: %f  \n", k, v/10000, v/sum)
	}

}

// 玄机 克敌 绝世3
func Test_SecondDamageComputation0(t *testing.T) {
	skillCoefficients := attribute_value.NewSkillCoefficient(attribute_value.TheMystery)
	character := attribute_value.BasicCharacter{
		MinAttack:              3545,
		MaxAttack:              5017,
		ElementalDamage:        1812,
		FixedDefeat:            1579,
		PercentageDefeat:       0.234,
		MonsterPenetration:     1322,
		CriticalHitProbability: 351,
		CriticalHitDamage:      0.55,
		FixedBrokenGuard:       790,
		PercentageBrokenGuard:  0.0,
	}

	req := types.CharacterDataReq{
		TestType:   3,
		Occupation: attribute_value.TheMystery,
	}

	secondDamage := damage.FinalSecondDamageComputation(req, character, skillCoefficients)

	sum := 0.0
	for _, v := range secondDamage {
		sum += v
	}
	fmt.Printf("%f万 \n", sum/10000)

	for k, v := range secondDamage {
		fmt.Printf("技能:%s 秒伤:%f 万  占比: %f  \n", k, v/10000, v/sum)
	}

}

// 夏天神相 3.6 绝世 3.49
func Test_SecondDamageComputation1(t *testing.T) {
	skillCoefficients := attribute_value.NewSkillCoefficient(attribute_value.ImageOfGod)

	character := attribute_value.BasicCharacter{
		MinAttack:              3498,
		MaxAttack:              5065,
		ElementalDamage:        2239,
		FixedDefeat:            1111,
		PercentageDefeat:       0.237,
		MonsterPenetration:     1504,
		CriticalHitProbability: 442,
		CriticalHitDamage:      0.757,
		FixedBrokenGuard:       699,
		PercentageBrokenGuard:  0.0,
	}

	req := types.CharacterDataReq{
		TestType:   3,
		Occupation: attribute_value.ImageOfGod,
	}

	secondDamage := damage.FinalSecondDamageComputation(req, character, skillCoefficients)

	sum := 0.0
	for _, v := range secondDamage {
		sum += v
	}
	fmt.Printf("%f万 \n", sum/10000)

	for k, v := range secondDamage {
		fmt.Printf("技能:%s 秒伤:%f 万  占比: %f  \n", k, v/10000, v/sum)
	}

}

// 加菲猫九灵 4.66
func Test_SecondDamageComputation2(t *testing.T) {
	skillCoefficients := attribute_value.NewSkillCoefficient(attribute_value.NineSpirits)

	character := attribute_value.BasicCharacter{
		MinAttack:              4618,
		MaxAttack:              6638,
		ElementalDamage:        3100,
		FixedDefeat:            495,
		PercentageDefeat:       0.25,
		MonsterPenetration:     1811,
		CriticalHitProbability: 274 + 721,
		CriticalHitDamage:      1.224,
		FixedBrokenGuard:       754,
		PercentageBrokenGuard:  0.0,
	}

	req := types.CharacterDataReq{
		TestType:   3,
		Occupation: attribute_value.NineSpirits,
	}

	secondDamage := damage.FinalSecondDamageComputation(req, character, skillCoefficients)

	sum := 0.0
	for _, v := range secondDamage {
		sum += v
	}
	fmt.Printf("%f万 \n", sum/10000)

	for k, v := range secondDamage {
		fmt.Printf("技能:%s 秒伤:%f 万  占比: %f  \n", k, v/10000, v/sum)
	}

}

// 龙 克敌 绝世4.99
func Test_SecondDamageComputation3(t *testing.T) {
	skillCoefficients := attribute_value.NewSkillCoefficient(attribute_value.DragonMindingLight)

	character := attribute_value.BasicCharacter{
		MinAttack:              4161,
		MaxAttack:              6019,
		ElementalDamage:        2067,
		FixedDefeat:            2518,
		PercentageDefeat:       0.248,
		MonsterPenetration:     1554,
		CriticalHitProbability: 351,
		CriticalHitDamage:      0.55,
		FixedBrokenGuard:       1331,
		PercentageBrokenGuard:  0.0,
	}

	req := types.CharacterDataReq{
		TestType:   3,
		Occupation: attribute_value.DragonMindingLight,
	}

	secondDamage := damage.FinalSecondDamageComputation(req, character, skillCoefficients)

	sum := 0.0
	for _, v := range secondDamage {
		sum += v
	}
	fmt.Printf("%f万 \n", sum/10000)

	for k, v := range secondDamage {
		fmt.Printf("技能:%s 秒伤:%f 万  占比: %f  \n", k, v/10000, v/sum)
	}

}

// 龙 元素 绝世5.15
func Test_SecondDamageComputation31(t *testing.T) {
	skillCoefficients := attribute_value.NewSkillCoefficient(attribute_value.DragonMindingLight)

	character := attribute_value.BasicCharacter{
		MinAttack:              4161,
		MaxAttack:              6051,
		ElementalDamage:        2989, //1650
		FixedDefeat:            1492, //2330
		PercentageDefeat:       0.248,
		MonsterPenetration:     1574,
		CriticalHitProbability: 380,
		CriticalHitDamage:      0.55,
		FixedBrokenGuard:       1282,
		PercentageBrokenGuard:  0.0,
	}

	req := types.CharacterDataReq{
		TestType:   3,
		Occupation: attribute_value.DragonMindingLight,
	}

	secondDamage := damage.FinalSecondDamageComputation(req, character, skillCoefficients)

	sum := 0.0
	for _, v := range secondDamage {
		sum += v
	}
	fmt.Printf("%f万 \n", sum/10000)

	for k, v := range secondDamage {
		fmt.Printf("技能:%s 秒伤:%f 万  占比: %f  \n", k, v/10000, v/sum)
	}

}

// 龙 元素 绝世4.99
func Test_SecondDamageComputation32(t *testing.T) {
	skillCoefficients := attribute_value.NewSkillCoefficient(attribute_value.DragonMindingLight)

	character := attribute_value.BasicCharacter{
		MinAttack:              4353,
		MaxAttack:              6280,
		ElementalDamage:        2479, //1650
		FixedDefeat:            1103, //2330
		PercentageDefeat:       0.25,
		MonsterPenetration:     1616,
		CriticalHitProbability: 583,
		CriticalHitDamage:      0.55,
		FixedBrokenGuard:       1256,
		PercentageBrokenGuard:  0.0,
	}

	req := types.CharacterDataReq{
		TestType:   3,
		Occupation: attribute_value.DragonMindingLight,
	}

	secondDamage := damage.FinalSecondDamageComputation(req, character, skillCoefficients)

	sum := 0.0
	for _, v := range secondDamage {
		sum += v
	}
	fmt.Printf("%f万 \n", sum/10000)

	for k, v := range secondDamage {
		fmt.Printf("技能:%s 秒伤:%f 万  占比: %f  \n", k, v/10000, v/sum)
	}

}

// 血河 绝世3.78
func Test_SecondDamageComputation4(t *testing.T) {
	skillCoefficients := attribute_value.NewSkillCoefficient(attribute_value.RiverOfBlood)

	character := attribute_value.BasicCharacter{
		MinAttack:              4037,
		MaxAttack:              5882,
		ElementalDamage:        2687,
		FixedDefeat:            510,
		PercentageDefeat:       0.243,
		MonsterPenetration:     1645,
		CriticalHitProbability: 568,
		CriticalHitDamage:      0.908,
		FixedBrokenGuard:       877,
		PercentageBrokenGuard:  0.0,
	}

	req := types.CharacterDataReq{
		TestType:   3,
		Occupation: attribute_value.RiverOfBlood,
	}

	secondDamage := damage.FinalSecondDamageComputation(req, character, skillCoefficients)

	sum := 0.0
	for _, v := range secondDamage {
		sum += v
	}
	fmt.Printf("%f万 \n", sum/10000)

	for k, v := range secondDamage {
		fmt.Printf("技能:%s 秒伤:%f 万  占比: %f  \n", k, v/10000, v/sum)
	}

}

// 碎梦 绝世4.04
func Test_SecondDamageComputation6(t *testing.T) {
	skillCoefficients := attribute_value.NewSkillCoefficient(attribute_value.BrokenDreams)

	character := attribute_value.BasicCharacter{
		MinAttack:              3703,
		MaxAttack:              5388,
		ElementalDamage:        3449,
		FixedDefeat:            698,
		PercentageDefeat:       0.245,
		MonsterPenetration:     1684,
		CriticalHitProbability: 351,
		CriticalHitDamage:      0.55,
		FixedBrokenGuard:       784,
		PercentageBrokenGuard:  0.0,
	}

	req := types.CharacterDataReq{
		TestType:   3,
		Occupation: attribute_value.BrokenDreams,
	}

	secondDamage := damage.FinalSecondDamageComputation(req, character, skillCoefficients)

	sum := 0.0
	for _, v := range secondDamage {
		sum += v
	}
	fmt.Printf("%f万 \n", sum/10000)

	for k, v := range secondDamage {
		fmt.Printf("技能:%s 秒伤:%f 万  占比: %f  \n", k, v/10000, v/sum)
	}

}

// 铁衣 h3.41  绝世2.99
func Test_SecondDamageComputation7(t *testing.T) {
	skillCoefficients := attribute_value.NewSkillCoefficient(attribute_value.GarmentOfIron)

	character := attribute_value.BasicCharacter{
		MinAttack:              4965,
		MaxAttack:              7098,
		ElementalDamage:        2450,
		FixedDefeat:            659,
		PercentageDefeat:       0.337,
		MonsterPenetration:     1474,
		CriticalHitProbability: 376,
		CriticalHitDamage:      0.55,
		FixedBrokenGuard:       687,
		PercentageBrokenGuard:  0.0,
	}

	req := types.CharacterDataReq{
		TestType:   3,
		Occupation: attribute_value.GarmentOfIron,
	}

	secondDamage := damage.FinalSecondDamageComputation(req, character, skillCoefficients)

	sum := 0.0
	for _, v := range secondDamage {
		sum += v
	}
	fmt.Printf("%f万 \n", sum/10000)

	for k, v := range secondDamage {
		fmt.Printf("技能:%s 秒伤:%f 万  占比: %f  \n", k, v/10000, v/sum)
	}

}

// 鸿音 绝世3.88
func Test_SecondDamageComputation8(t *testing.T) {
	skillCoefficients := attribute_value.NewSkillCoefficient(attribute_value.HighTone)

	character := attribute_value.BasicCharacter{
		MinAttack:              4036,
		MaxAttack:              5845,
		ElementalDamage:        2477,
		FixedDefeat:            825,
		PercentageDefeat:       0.25,
		MonsterPenetration:     1636,
		CriticalHitProbability: 462,
		CriticalHitDamage:      0.93,
		FixedBrokenGuard:       1054,
		PercentageBrokenGuard:  0.0,
	}

	req := types.CharacterDataReq{
		TestType:   3,
		Occupation: attribute_value.HighTone,
	}

	secondDamage := damage.FinalSecondDamageComputation(req, character, skillCoefficients)

	sum := 0.0
	for _, v := range secondDamage {
		sum += v
	}
	fmt.Printf("%f万 \n", sum/10000)

	for k, v := range secondDamage {
		fmt.Printf("技能:%s 秒伤:%f 万  占比: %f  \n", k, v/10000, v/sum)
	}

}
