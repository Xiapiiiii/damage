package test

import (
	"damage/attribute_value"
	"damage/damage"
	"damage/types"
	"fmt"
	"testing"
)

// 玄机 4.4 绝世3.8
func Test_SecondDamageComputation(t *testing.T) {
	skillCoefficients := attribute_value.NewSkillCoefficient(attribute_value.TheMystery)
	character := attribute_value.BasicCharacter{
		MinAttack:              3527,
		MaxAttack:              5064,
		ElementalDamage:        1750,
		FixedDefeat:            2023,
		PercentageDefeat:       0.284,
		MonsterPenetration:     1496,
		CriticalHitProbability: 348,
		CriticalHitDamage:      0.55,
		FixedBrokenGuard:       730,
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

// 夏天神相 3.6 绝世 3.4
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

// 芥末九灵 2.77
func Test_SecondDamageComputation2(t *testing.T) {
	skillCoefficients := attribute_value.NewSkillCoefficient(attribute_value.NineSpirits)

	character := attribute_value.BasicCharacter{
		MinAttack:              3398,
		MaxAttack:              4913,
		ElementalDamage:        1900,
		FixedDefeat:            810,
		PercentageDefeat:       0.234,
		MonsterPenetration:     1182,
		CriticalHitProbability: 1200,
		CriticalHitDamage:      0.964,
		FixedBrokenGuard:       705,
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

// 龙 5.38 绝世4.78
func Test_SecondDamageComputation3(t *testing.T) {
	skillCoefficients := attribute_value.NewSkillCoefficient(attribute_value.DragonMindingLight)

	character := attribute_value.BasicCharacter{
		MinAttack:              4011,
		MaxAttack:              5802,
		ElementalDamage:        1650,
		FixedDefeat:            2330,
		PercentageDefeat:       0.25,
		MonsterPenetration:     1490,
		CriticalHitProbability: 318,
		CriticalHitDamage:      0.55,
		FixedBrokenGuard:       1298,
		PercentageBrokenGuard:  0.0,
	}

	req := types.CharacterDataReq{
		TestType:   2,
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

// 血河 2.7 绝世2.2
func Test_SecondDamageComputation4(t *testing.T) {
	skillCoefficients := attribute_value.NewSkillCoefficient(attribute_value.RiverOfBlood)

	character := attribute_value.BasicCharacter{
		MinAttack:              3746,
		MaxAttack:              5422,
		ElementalDamage:        1502,
		FixedDefeat:            617,
		PercentageDefeat:       0.25,
		MonsterPenetration:     1224,
		CriticalHitProbability: 698,
		CriticalHitDamage:      0.98,
		FixedBrokenGuard:       737,
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

// 鸿音 绝世3.5
func Test_SecondDamageComputation5(t *testing.T) {
	skillCoefficients := attribute_value.NewSkillCoefficient(attribute_value.HighTone)

	character := attribute_value.BasicCharacter{
		MinAttack:              3508,
		MaxAttack:              5116,
		ElementalDamage:        2030,
		FixedDefeat:            1929,
		PercentageDefeat:       0.281,
		MonsterPenetration:     1600,
		CriticalHitProbability: 429,
		CriticalHitDamage:      0.55,
		FixedBrokenGuard:       974,
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

// 碎梦 绝世2.9
func Test_SecondDamageComputation6(t *testing.T) {
	skillCoefficients := attribute_value.NewSkillCoefficient(attribute_value.BrokenDreams)

	character := attribute_value.BasicCharacter{
		MinAttack:              3659,
		MaxAttack:              5359,
		ElementalDamage:        2100,
		FixedDefeat:            1111,
		PercentageDefeat:       0.24,
		MonsterPenetration:     1300,
		CriticalHitProbability: 324,
		CriticalHitDamage:      0.73,
		FixedBrokenGuard:       889,
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
