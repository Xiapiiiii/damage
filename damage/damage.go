package damage

import (
	"damage/attribute_value"
	"damage/types"
)

func FinalSecondDamageComputation(req types.CharacterDataReq, character attribute_value.BasicCharacter, skillCoefficients map[string]attribute_value.SkillBaseCoefficient) (secondDamage map[string]float64) {

	secondDamage = make(map[string]float64)
	for skillName, skillCoefficient := range skillCoefficients {
		secondDamage[skillName] = SecondDamageComputation(req, skillCoefficient, character)
	}
	return
}

func SecondDamageComputation(req types.CharacterDataReq, skillCoefficient attribute_value.SkillBaseCoefficient, character attribute_value.BasicCharacter) float64 {
	basicMonster := attribute_value.NewBasicMonster(req.TestType)

	//怪穿系数计算
	monsterCoefficient := 1.0
	if character.MonsterPenetration < basicMonster.MonsterResistance {
		monsterCoefficient = 1 - (basicMonster.MonsterResistance-character.MonsterPenetration)/100*0.035
	}

	//防御校准
	if character.FixedBrokenGuard > basicMonster.Defense {
		character.FixedBrokenGuard = basicMonster.Defense
		character.PercentageBrokenGuard = 0
	}

	//防御系数计算
	defenseCoefficient := 3588 / (3588 + (basicMonster.Defense-character.FixedBrokenGuard)*(1-character.PercentageBrokenGuard))

	//会心率计算
	criticalHitRate := character.CriticalHitProbability /
		(character.CriticalHitProbability + 69*10)

	//会心校准
	if criticalHitRate < basicMonster.CriticalStrikeResistance {
		criticalHitRate = basicMonster.CriticalStrikeResistance
	}

	//攻击伤害
	attackDamage := (character.MinAttack + character.MaxAttack) /
		2 * skillCoefficient.AttackCoefficient *
		defenseCoefficient *
		monsterCoefficient
	//元素伤害
	elementalDamage := character.ElementalDamage *
		skillCoefficient.ElementsCoefficient *
		0.945

	//克敌附加
	defeatDamage := character.FixedDefeat *
		monsterCoefficient *
		attribute_value.DefeatFixM[req.Occupation]

	secondDamage := (attackDamage + elementalDamage + defeatDamage) *
		skillCoefficient.ShootRate *
		(1 + (criticalHitRate-basicMonster.CriticalStrikeResistance)*
			character.CriticalHitDamage) *
		(1 + character.PercentageDefeat) *
		attribute_value.PVEFixM[req.Occupation]

	return secondDamage
}
