package test

import (
	"damage/attribute_value"
	"damage/damage"
	"damage/types"
	"fmt"
	"testing"
)

// 玄机 4.4
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
		TestType:   2,
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

// 夏天神相 3.1
func Test_SecondDamageComputation1(t *testing.T) {
	skillCoefficients := attribute_value.NewSkillCoefficient(attribute_value.ImageOfGod)

	character := attribute_value.BasicCharacter{
		MinAttack:              3398,
		MaxAttack:              4913,
		ElementalDamage:        1800,
		FixedDefeat:            1312,
		PercentageDefeat:       0.234,
		MonsterPenetration:     1322,
		CriticalHitProbability: 427,
		CriticalHitDamage:      0.76,
		FixedBrokenGuard:       800,
		PercentageBrokenGuard:  0.0,
	}

	req := types.CharacterDataReq{
		TestType:   2,
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
		TestType:   2,
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

// 龙 4.6
func Test_SecondDamageComputation3(t *testing.T) {
	skillCoefficients := attribute_value.NewSkillCoefficient(attribute_value.DragonMindingLight)

	character := attribute_value.BasicCharacter{
		MinAttack:              3761,
		MaxAttack:              5410,
		ElementalDamage:        2414,
		FixedDefeat:            1517,
		PercentageDefeat:       0.27,
		MonsterPenetration:     1398,
		CriticalHitProbability: 269,
		CriticalHitDamage:      0.55,
		FixedBrokenGuard:       1159,
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

// 血河 2.7
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
		TestType:   2,
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
