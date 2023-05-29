package attribute_value

// 装备词条
const (
	ElementalDamage        = "全元素攻击"
	FixedDefeat            = "克敌"
	CriticalHitProbability = "会心"
	CriticalHitDamage      = "会心伤害"
)

// 天玑词条
const ()

// EquipmentPool 装备属性池
type EquipmentPool struct {
	Weapons   map[string]float64 //武器
	Ring1     map[string]float64 //戒指1
	Ring2     map[string]float64 //戒指2
	Bracelet1 map[string]float64 //手镯1
	Bracelet2 map[string]float64 //手镯2
	Necklace  map[string]float64 //项链

	GlovesEnchant   map[string]float64 //手套天玑
	NecklaceEnchant map[string]float64 //项链天玑
	WeaponsEnchant  map[string]float64 //武器天玑
	BeltEnchant     map[string]float64 //腰带天玑
	ShoesEnchant    map[string]float64 //鞋子天玑
}

// RoleEquipment  角色装备属性
type RoleEquipment struct {
	Weapons   Equipment //武器
	Ring1     Equipment //戒指1
	Ring2     Equipment //戒指2
	Bracelet1 Equipment //手镯1
	Bracelet2 Equipment //手镯2
	Necklace  Equipment //项链

	//GlovesEnchant   Equipment //手套天玑
	//NecklaceEnchant Equipment //项链天玑
	//WeaponsEnchant  Equipment //武器天玑
	//BeltEnchant     Equipment //腰带天玑
	//ShoesEnchant    Equipment //鞋子天玑
	//
	//Helmet     Equipment //头盔
	//Gloves     Equipment //手套
	//WristGuard Equipment //护腕
	//Clothes    Equipment //衣服
	//Belt       Equipment //腰带
	//Shoes      Equipment //鞋子
}

// Equipment 装备属性
type Equipment struct {
	ElementalDamage        float64 //全元素攻击
	FixedDefeat            float64 //克敌
	CriticalHitProbability float64 //会心
	CriticalHitDamage      float64 //会心伤害
}

// GenerationEquipmentPool 生成太虚装备属性池
func GenerationEquipmentPool() (equipmentPool EquipmentPool) {
	return EquipmentPool{
		Weapons: map[string]float64{
			ElementalDamage:   178,
			FixedDefeat:       251,
			CriticalHitDamage: 0.18,
		},
		Ring1: map[string]float64{
			ElementalDamage:        178,
			FixedDefeat:            251,
			CriticalHitProbability: 73,
		},
		Ring2: map[string]float64{
			ElementalDamage:        178,
			FixedDefeat:            251,
			CriticalHitProbability: 73,
		},
		Bracelet1: map[string]float64{
			ElementalDamage:        178,
			FixedDefeat:            251,
			CriticalHitProbability: 73,
		},
		Bracelet2: map[string]float64{
			ElementalDamage:        178,
			FixedDefeat:            251,
			CriticalHitProbability: 73,
		},
		Necklace: map[string]float64{
			ElementalDamage:   178,
			FixedDefeat:       251,
			CriticalHitDamage: 0.18,
		},

		//GlovesEnchant: map[string]float64{
		//	ElementalDamage:        200,
		//	FixedDefeat:            250,
		//	CriticalHitProbability: 0.05,
		//},
		//NecklaceEnchant: map[string]float64{
		//	ElementalDamage:        200,
		//	FixedDefeat:            250,
		//	CriticalHitProbability: 0.05,
		//},
		//WeaponsEnchant: map[string]float64{
		//	ElementalDamage:        200,
		//	FixedDefeat:            250,
		//	CriticalHitProbability: 0.05,
		//},
		//BeltEnchant: map[string]float64{
		//	ElementalDamage:        200,
		//	FixedDefeat:            250,
		//	CriticalHitProbability: 0.05,
		//},
		//ShoesEnchant: map[string]float64{
		//	ElementalDamage:        200,
		//	FixedDefeat:            250,
		//	CriticalHitProbability: 0.05,
		//},
	}
}

// GenerationRoleEquipmentPool 生成角色装备属性池
func GenerationRoleEquipmentPool(equipmentPool EquipmentPool) (roleEquipmentPool []RoleEquipment) {
	roleEquipmentPool = make([]RoleEquipment, 1000)

	for weaponsType, weapons := range equipmentPool.Weapons {
		for ring1Type, ring1 := range equipmentPool.Ring1 {
			for ring2Type, ring2 := range equipmentPool.Ring2 {
				for bracelet1Type, bracelet1 := range equipmentPool.Bracelet1 {
					for bracelet2Type, bracelet2 := range equipmentPool.Bracelet2 {
						for necklaceType, necklace := range equipmentPool.Necklace {
							roleEquipment := RoleEquipment{}
							//fmt.Println(necklaceType, necklace)
							if weaponsType == ElementalDamage {
								roleEquipment.Weapons.ElementalDamage = weapons
							}
							if weaponsType == FixedDefeat {
								roleEquipment.Weapons.FixedDefeat = weapons
							}
							if weaponsType == CriticalHitDamage {
								roleEquipment.Weapons.CriticalHitDamage = weapons
							}

							if ring1Type == ElementalDamage {
								roleEquipment.Ring1.ElementalDamage = ring1
							}
							if ring1Type == FixedDefeat {
								roleEquipment.Ring1.FixedDefeat = ring1
							}
							if ring1Type == CriticalHitProbability {
								roleEquipment.Ring1.CriticalHitProbability = ring1
							}

							if ring2Type == ElementalDamage {
								roleEquipment.Ring2.ElementalDamage = ring2
							}
							if ring2Type == FixedDefeat {
								roleEquipment.Ring2.FixedDefeat = ring2
							}
							if ring2Type == CriticalHitProbability {
								roleEquipment.Ring2.CriticalHitProbability = ring2
							}

							if bracelet1Type == ElementalDamage {
								roleEquipment.Bracelet1.ElementalDamage = bracelet1
							}
							if bracelet1Type == FixedDefeat {
								roleEquipment.Bracelet1.FixedDefeat = bracelet1
							}
							if bracelet1Type == CriticalHitProbability {
								roleEquipment.Bracelet1.CriticalHitProbability = bracelet1
							}

							if bracelet2Type == ElementalDamage {
								roleEquipment.Bracelet2.ElementalDamage = bracelet2
							}
							if bracelet2Type == FixedDefeat {
								roleEquipment.Bracelet2.FixedDefeat = bracelet2
							}
							if bracelet2Type == CriticalHitProbability {
								roleEquipment.Bracelet2.CriticalHitProbability = bracelet2
							}

							if necklaceType == ElementalDamage {
								roleEquipment.Necklace.ElementalDamage = necklace
							}
							if necklaceType == FixedDefeat {
								roleEquipment.Necklace.FixedDefeat = necklace
							}
							if necklaceType == CriticalHitDamage {
								roleEquipment.Necklace.CriticalHitDamage = necklace
							}
							//log.Println(roleEquipment)
							roleEquipmentPool = append(roleEquipmentPool, roleEquipment)
						}
					}
				}
			}
		}
	}

	return roleEquipmentPool
}
