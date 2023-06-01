package test

import (
	"damage/attribute_value"
	"damage/damage"
	"damage/types"
	"fmt"
	"testing"
)

func Test_imageOfGodMatch(t *testing.T) {
	skillCoefficients := attribute_value.NewSkillCoefficient(attribute_value.RiverOfBlood)

	equipmentPool := attribute_value.GenerationEquipmentPool()
	//fmt.Println(equipmentPool)
	roleEquipmentPool := attribute_value.GenerationRoleEquipmentPool(equipmentPool)
	//fmt.Println(roleEquipmentPool)
	character := attribute_value.BasicCharacter{
		MinAttack:              3498,
		MaxAttack:              5065,
		ElementalDamage:        1900,
		FixedDefeat:            611,
		PercentageDefeat:       0.237,
		MonsterPenetration:     1504,
		CriticalHitProbability: 370,
		CriticalHitDamage:      0.58,
		FixedBrokenGuard:       699,
		PercentageBrokenGuard:  0.0,
	}

	req := types.CharacterDataReq{
		TestType:   3,
		Occupation: attribute_value.RiverOfBlood,
	}

	c := attribute_value.BasicCharacter{}
	max := 0.0
	roleBestEquipment := attribute_value.RoleEquipment{}
	for _, roleEquipment := range roleEquipmentPool {
		c = character

		c.ElementalDamage += roleEquipment.Weapons.ElementalDamage
		c.ElementalDamage += roleEquipment.Ring1.ElementalDamage
		c.ElementalDamage += roleEquipment.Ring2.ElementalDamage
		c.ElementalDamage += roleEquipment.Bracelet1.ElementalDamage
		c.ElementalDamage += roleEquipment.Bracelet2.ElementalDamage
		c.ElementalDamage += roleEquipment.Necklace.ElementalDamage

		c.FixedDefeat += roleEquipment.Weapons.FixedDefeat
		c.FixedDefeat += roleEquipment.Ring1.FixedDefeat
		c.FixedDefeat += roleEquipment.Ring2.FixedDefeat
		c.FixedDefeat += roleEquipment.Bracelet1.FixedDefeat
		c.FixedDefeat += roleEquipment.Bracelet2.FixedDefeat

		c.CriticalHitProbability += roleEquipment.Ring1.CriticalHitProbability
		c.CriticalHitProbability += roleEquipment.Ring2.CriticalHitProbability
		c.CriticalHitProbability += roleEquipment.Bracelet1.CriticalHitProbability
		c.CriticalHitProbability += roleEquipment.Bracelet2.CriticalHitProbability

		c.CriticalHitDamage += roleEquipment.Weapons.CriticalHitProbability
		c.CriticalHitDamage += roleEquipment.Necklace.CriticalHitProbability

		secondDamage := damage.FinalSecondDamageComputation(req, c, skillCoefficients)

		sum := 0.0
		for _, d := range secondDamage {
			sum += d
		}

		//fmt.Println(sum)
		if sum > max {
			max = sum
			roleBestEquipment = roleEquipment
			//fmt.Println(roleBestEquipment)
		}

		c = character
	}

	fmt.Printf("最大秒伤 %f万\n", max/10000)

	if roleBestEquipment.Weapons.ElementalDamage > 0 {
		fmt.Printf("武器最佳属性全元素 %f\n", roleBestEquipment.Weapons.ElementalDamage)
	}
	if roleBestEquipment.Weapons.CriticalHitDamage > 0 {
		fmt.Printf("武器最佳属性会心伤害 %f\n", roleBestEquipment.Weapons.CriticalHitDamage)
	}
	if roleBestEquipment.Weapons.FixedDefeat > 0 {
		fmt.Printf("武器最佳属性克敌 %f\n", roleBestEquipment.Weapons.FixedDefeat)
	}
	if roleBestEquipment.Necklace.CriticalHitDamage > 0 {
		fmt.Printf("项链最佳属性会心伤害 %f\n", roleBestEquipment.Necklace.CriticalHitDamage)
	}
	if roleBestEquipment.Necklace.ElementalDamage > 0 {
		fmt.Printf("项链最佳属性全元素 %f\n", roleBestEquipment.Necklace.ElementalDamage)
	}
	if roleBestEquipment.Necklace.FixedDefeat > 0 {
		fmt.Printf("项链最佳属性克敌 %f\n", roleBestEquipment.Necklace.FixedDefeat)
	}

	if roleBestEquipment.Ring1.ElementalDamage > 0 {
		fmt.Printf("戒指1最佳属性全元素 %f\n", roleBestEquipment.Ring1.ElementalDamage)
	}
	if roleBestEquipment.Ring2.ElementalDamage > 0 {
		fmt.Printf("戒指2最佳属性全元素 %f\n", roleBestEquipment.Ring2.ElementalDamage)
	}
	if roleBestEquipment.Bracelet1.ElementalDamage > 0 {
		fmt.Printf("手镯1最佳属性全元素 %f\n", roleBestEquipment.Bracelet1.ElementalDamage)
	}
	if roleBestEquipment.Bracelet2.ElementalDamage > 0 {
		fmt.Printf("手镯2最佳属性全元素 %f\n", roleBestEquipment.Bracelet2.ElementalDamage)
	}
	if roleBestEquipment.Ring1.FixedDefeat > 0 {
		fmt.Printf("戒指1最佳属性克敌 %f\n", roleBestEquipment.Ring1.FixedDefeat)
	}
	if roleBestEquipment.Ring2.FixedDefeat > 0 {
		fmt.Printf("戒指2最佳属性克敌 %f\n", roleBestEquipment.Ring2.FixedDefeat)
	}

	if roleBestEquipment.Bracelet1.FixedDefeat > 0 {
		fmt.Printf("手镯1最佳属性克敌 %f\n", roleBestEquipment.Bracelet1.FixedDefeat)
	}
	if roleBestEquipment.Bracelet2.FixedDefeat > 0 {
		fmt.Printf("手镯2最佳属性克敌 %f\n", roleBestEquipment.Bracelet2.FixedDefeat)
	}

	if roleBestEquipment.Ring1.CriticalHitProbability > 0 {
		fmt.Printf("戒指1最佳属性会心 %f\n", roleBestEquipment.Ring1.CriticalHitProbability)
	}

	if roleBestEquipment.Ring2.CriticalHitProbability > 0 {
		fmt.Printf("戒指2最佳属性会心 %f\n", roleBestEquipment.Ring2.CriticalHitProbability)
	}

	if roleBestEquipment.Bracelet1.CriticalHitProbability > 0 {
		fmt.Printf("手镯1最佳属性会心 %f\n", roleBestEquipment.Bracelet1.CriticalHitProbability)
	}

	if roleBestEquipment.Bracelet2.CriticalHitProbability > 0 {
		fmt.Printf("手镯2最佳属性会心 %f\n", roleBestEquipment.Bracelet2.CriticalHitProbability)
	}

	//secondDamage := FinalSecondDamageComputation(character, skillCoefficients)
	//
	//for skillName, damage := range secondDamage {
	//	fmt.Printf("%s每秒伤害 %f万\n", skillName, damage/10000)
	//}

}
