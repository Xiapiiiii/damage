package define

// 太虚装备属性
const (
	Elemental = iota + 1
)

// 奇遇品质
const (
	QualityBig    = iota + 1 //大吉
	QualityMedium            //中吉
	QualitySmall             //小吉
	QualityFierce            //凶
)

// 地点
const (
	TaoXi      = iota + 1 //桃溪村
	BianJing              //汴京
	HangZhou              //杭州
	CiZhou                //磁州
	CangZhou              //沧州
	SanQinShan            //三清山
	YanMenGuan            //雁门关
	SongLiao              //宋辽边境
	DiXian                //谪仙岛
	YaoWang               //药王谷
	BiXue                 //碧血营
)

var Location = map[int64]string{
	1:  "桃溪村",
	2:  "汴京",
	3:  "杭州",
	4:  "磁州",
	5:  "沧州",
	6:  "三清山",
	7:  "雁门关",
	8:  "宋辽边境",
	9:  "谪仙岛",
	10: "药王谷",
	11: "碧血营",
}

var LocationS = map[string]int64{
	"桃溪村":  1,
	"汴京":   2,
	"杭州":   3,
	"磁州":   4,
	"沧州":   5,
	"三清山":  6,
	"雁门关":  7,
	"宋辽边境": 8,
	"谪仙岛":  9,
	"药王谷":  10,
	"碧血营":  11,
}

var Quality = map[int64]string{
	1: "大吉",
	2: "中吉",
	3: "小吉",
	4: "凶",
}

var QualityS = map[string]int64{
	"大吉": 1,
	"中吉": 2,
	"小吉": 3,
	"凶":  4,
}
