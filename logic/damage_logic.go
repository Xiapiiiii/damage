package logic

import (
	"damage/attribute_value"
	"damage/damage"
	"damage/model"
	"damage/svc"
	"damage/types"
	"fmt"
	"github.com/gin-gonic/gin"
	"go.mongodb.org/mongo-driver/mongo"
	"math"
	"sync"
	"time"
)

var (
	numberOfHits int64
	timeT        = time.NewTicker(1 * time.Minute)
	mutex        = &sync.Mutex{}
)

type BestMatch struct {
	Message  string //最高秒伤
	Message1 string //武器最佳属性
	Message2 string //项链最佳属性
	Message3 string //手镯1最佳属性
	Message4 string //手镯2最佳属性
	Message5 string //戒指1最佳属性
	Message6 string //戒指2最佳属性
}

// CalculatedDamageLogic 计算秒伤逻辑
func CalculatedDamageLogic(c *gin.Context, req types.CharacterDataReq) (characterDataResp types.CharacterDataResp, err error) {
	ctx := c.MustGet("ctx").(*svc.ServiceContext)
	numberOfHits++
	//log.Println(numberOfHits)

	select {
	case <-timeT.C:
		mutex.Lock()
		todayAccess, err := ctx.AccessDayModel.GetTodayAccessNum(c.Request.Context())
		if err != nil {
			if err != mongo.ErrNoDocuments {
				return types.CharacterDataResp{}, err
			}
		}

		if todayAccess == nil {
			order := model.AccessDay{
				AccessNum: numberOfHits,
				CreatedAt: time.Now(),
				UpdatedAt: time.Now(),
			}
			_, err := ctx.AccessDayModel.CreateTodayAccessNum(c.Request.Context(), &order)
			if err != nil {
				return types.CharacterDataResp{}, err
			}
			numberOfHits = 0
		} else {
			todayAccess.AccessNum += numberOfHits
			todayAccess.UpdatedAt = time.Now()
			err := ctx.AccessDayModel.UpdateTodayAccessNum(c.Request.Context(), todayAccess)
			if err != nil {
				return types.CharacterDataResp{}, err
			}
			numberOfHits = 0
		}
		mutex.Unlock()
	default:
	}

	// 逻辑处理
	skillCoefficients := attribute_value.NewSkillCoefficient(int(req.Occupation))

	character := attribute_value.BasicCharacter{
		MinAttack:              req.MinAttack,
		MaxAttack:              req.MaxAttack,
		ElementalDamage:        req.ElementalDamage,
		FixedDefeat:            req.FixedDefeat,
		PercentageDefeat:       req.PercentageDefeat / 100,
		MonsterPenetration:     req.MonsterPenetration,
		CriticalHitProbability: req.CriticalHitProbability,
		CriticalHitDamage:      req.CriticalHitDamage/100 - 1,
		FixedBrokenGuard:       req.FixedBrokenGuard,
		PercentageBrokenGuard:  req.PercentageBrokenGuard / 100,
	}

	if req.Occupation == attribute_value.NineSpirits {
		character.CriticalHitProbability += 721
	}

	damageM := damage.FinalSecondDamageComputation(req, character, skillCoefficients)

	sum := 0.0
	for _, d := range damageM {
		sum += d
	}
	characterDataResp.Message = "success"
	sum = sum / 10000
	sum = math.Round(sum*1000) / 1000
	//sumS := fmt.Sprintf("%f 万", sum)
	characterDataResp.Data.SecondDamageSum = fmt.Sprint(sum)
	return
}

// CalculatingTheBestAttributes 计算最佳属性
func CalculatingTheBestAttributes(req types.CharacterDataReq) (bestMatch BestMatch) {
	skillCoefficients := attribute_value.NewSkillCoefficient(attribute_value.ImageOfGod)
	equipmentPool := attribute_value.GenerationEquipmentPool()
	//fmt.Println(equipmentPool)
	roleEquipmentPool := attribute_value.GenerationRoleEquipmentPool(equipmentPool)
	//fmt.Println(roleEquipmentPool)
	character := attribute_value.BasicCharacter{
		MinAttack:              3398,
		MaxAttack:              4913,
		ElementalDamage:        1450,
		FixedDefeat:            612,
		PercentageDefeat:       0.234,
		MonsterPenetration:     1500,
		CriticalHitProbability: 400,
		CriticalHitDamage:      0.57,
		FixedBrokenGuard:       800,
		PercentageBrokenGuard:  0.0,
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

	bestMatch.Message = fmt.Sprintf("%f万\n", max/10000)
	if roleBestEquipment.Weapons.ElementalDamage > 0 {
		fmt.Printf("武器最佳属性全元素 %f\n", roleBestEquipment.Weapons.ElementalDamage)
		bestMatch.Message1 = "全元素 " + fmt.Sprintf("%f\n", roleBestEquipment.Weapons.ElementalDamage)
	}
	if roleBestEquipment.Weapons.CriticalHitDamage > 0 {
		fmt.Printf("武器最佳属性会心伤害 %f\n", roleBestEquipment.Weapons.CriticalHitDamage)
		bestMatch.Message1 = "会心伤害 " + fmt.Sprintf("%f\n", roleBestEquipment.Weapons.CriticalHitDamage)
	}
	if roleBestEquipment.Weapons.FixedDefeat > 0 {
		fmt.Printf("武器最佳属性克敌 %f\n", roleBestEquipment.Weapons.FixedDefeat)
		bestMatch.Message1 = "克敌 " + fmt.Sprintf("%f\n", roleBestEquipment.Weapons.FixedDefeat)
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

	return bestMatch
}

func GetCollocationLogic(req types.GetCollocationReq) (collocationDataResp types.GetCollocationResp) {

	return
}
