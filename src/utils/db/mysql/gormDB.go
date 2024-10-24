package mysql

import "gorm.io/gorm"

// 시나리오 상수 정의
// 연인, 혼반, 가족, 다이어트, 회식, 친구
const (
	ScenarioAll    = 0
	ScenarioCouple = 1 + iota
	ScenarioSolo
	ScenarioFamily
	ScenarioCompany
	ScenarioFriend
)

// 식사 시간 상수 정의
// 아침, 점심, 저녁, 브런치, 간식, 야식
const (
	TimeAll      = 0
	TimeMorning  = 1 + iota //아침
	TimeLunch               //점심
	TimeDinner              //저녁
	TimeSnack               //간식
	TimeMidnight            //야식
)

// 음식 종류 상수 정의
// 한식, 중식, 일식, 양식, 분식,베트남 음식, 인도 음식, 패스트 푸드, 디저트, 퓨전 요리
const (
	TypeAll        = 0
	TypeKorean     = 1 + iota //한식
	TypeChinese               //중식
	TypeJapanese              //일식
	TypeWestern               //양식
	TypeStreetFood            //분식
	TypeFastFood              //패스트 푸드
	TypeVietnamese            //베트남 음식
	TypeIndian                //인도 음식
	TypeDessert               //디저트
	TypeFusion                //퓨전 요리
)

// 기분/테마 상수 정의
// 스트레스 해소, 피로 회복, 기분 전환, 제철 음식, 영양식, 특별한 날
const (
	ThemeAll             = 0
	ThemeStressRelief    = 1 + iota // 스트레스 해소
	ThemeHangover                   // 해장
	ThemeFatigueRecovery            // 피로 회복
	ThemeDiet                       //다이어트
	ThemeSeasonalFood               //제철 음식
)

// 맛 상수 정의
// 매운맛, 감칠맛, 고소한맛, 단맛, 짠맛, 싱거운맛
const (
	FlavorAll    = 0
	FlavorSpicy  = 1 + iota //매운맛
	FlavorSavory            //감칠맛
	FlavorNutty             //고소한맛
	FlavorSweet             //단맛
	FlavorSour              //싱거운맛
)

// 맵 정의
var ScenarioMap = map[string]int{
	"연인": ScenarioCouple,
	"혼밥": ScenarioSolo,
	"가족": ScenarioFamily,
	"회식": ScenarioCompany,
	"친구": ScenarioFriend,
}

var TimeMap = map[string]int{
	"아침": TimeMorning,
	"점심": TimeLunch,
	"저녁": TimeDinner,
	"간식": TimeSnack,
	"야식": TimeMidnight,
}

var TypeMap = map[string]int{
	"한식":     TypeKorean,
	"중식":     TypeChinese,
	"일식":     TypeJapanese,
	"양식":     TypeWestern,
	"분식":     TypeStreetFood,
	"베트남 음식": TypeVietnamese,
	"인도 음식":  TypeIndian,
	"패스트 푸드": TypeFastFood,
	"디저트":    TypeDessert,
	"퓨전 요리":  TypeFusion,
}

var FlavorMap = map[string]int{
	"매운맛":  FlavorSpicy,
	"감칠맛":  FlavorSavory,
	"고소한맛": FlavorNutty,
	"단맛":   FlavorSweet,
	"싱거운맛": FlavorSour,
}

var ThemeMap = map[string]int{
	"스트레스 해소": ThemeStressRelief,
	"해장":      ThemeHangover,
	"피로 회복":   ThemeFatigueRecovery,
	"다이어트":    ThemeDiet,
	"제철 음식":   ThemeSeasonalFood,
}

// 반대 맵 정의 (int -> string)
var ScenarioReverseMap = make(map[int]string)
var TimeReverseMap = make(map[int]string)
var TypeReverseMap = make(map[int]string)
var FlavorReverseMap = make(map[int]string)
var ThemeReverseMap = make(map[int]string)

type Tokens struct {
	gorm.Model
	UserID           uint   `json:"userID" gorm:"column:user_id"`
	AccessToken      string `json:"accessToken" gorm:"column:access_token"`
	RefreshToken     string `json:"refreshToken" gorm:"column:refresh_token"`
	RefreshExpiredAt int64  `json:"refreshExpiredAt" gorm:"column:refresh_expired_at"`
}

type Users struct {
	gorm.Model
	Email    string `json:"email" gorm:"column:email"`
	Password string `json:"password" gorm:"column:password"`
	Birth    string `json:"birth" gorm:"column:birth"`
	Name     string `json:"name" gorm:"column:name"`
	Sex      string `json:"sex" gorm:"column:sex"`
	Provider string `json:"provider" gorm:"column:provider"`
}

type Foods struct {
	gorm.Model
	Name        string `json:"name" gorm:"column:name"`
	FoodImageID int    `json:"foodImageID" gorm:"column:food_image_id"`
	ScenarioID  int    `json:"scenarioID" gorm:"column:scenario_id"`
	TimeID      int    `json:"timeID" gorm:"column:time_id"`
	TypeID      int    `json:"typeID" gorm:"column:type_id"`
	ThemeID     int    `json:"themeID" gorm:"column:theme_id"`
	FlavorID    int    `json:"flavorID" gorm:"column:flavor_id"`
}
type FoodHistory struct {
	gorm.Model
	UserID uint   `json:"userID" gorm:"column:user_id"`
	FoodID uint   `json:"foodID" gorm:"column:food_id"`
	Name   string `json:"name" gorm:"column:name"`
}

type MetaTables struct {
	gorm.Model
	TableName        string `json:"tableName" gorm:"column:table_name"`
	TableDescription string `json:"tableDescription" gorm:"column:table_description"`
}

type Scenarios struct {
	gorm.Model
	Name        string `json:"name" gorm:"column:name"`
	Image       string `json:"image" gorm:"column:image"`
	Description string `json:"description" gorm:"column:description"`
}

type Times struct {
	gorm.Model
	Name        string `json:"name" gorm:"column:name"`
	Image       string `json:"image" gorm:"column:image"`
	Description string `json:"description" gorm:"column:description"`
}

type Types struct {
	gorm.Model
	Name        string `json:"name" gorm:"column:name"`
	Image       string `json:"image" gorm:"column:image"`
	Description string `json:"description" gorm:"column:description"`
}

type Flavors struct {
	gorm.Model
	Name        string `json:"name" gorm:"column:name"`
	Image       string `json:"image" gorm:"column:image"`
	Description string `json:"description" gorm:"column:description"`
}
type Themes struct {
	gorm.Model
	Name        string `json:"name" gorm:"column:name"`
	Image       string `json:"image" gorm:"column:image"`
	Description string `json:"description" gorm:"column:description"`
}

type UserAuths struct {
	gorm.Model
	Email    string `json:"email" gorm:"column:email"`
	AuthCode string `json:"authCode" gorm:"column:auth_code"`
	Type     string `json:"type" gorm:"column:type"`
}

type FoodImages struct {
	gorm.Model
	Name  string `json:"name" gorm:"column:name"`
	Image string `json:"image" gorm:"column:image"`
}

type Reports struct {
	gorm.Model
	UserID int    `json:"userID" gorm:"column:user_id"`
	Reason string `json:"reason" gorm:"column:reason"`
}

type Nutrients struct {
	gorm.Model
	FoodName     string  `json:"foodName" gorm:"column:food_name"`
	Amount       string  `json:"amount" gorm:"column:amount"`
	Calorie      float64 `json:"calories" gorm:"column:calorie"`
	Carbohydrate float64 `json:"carbohydrate" gorm:"column:carbohydrate"`
	Protein      float64 `json:"protein" gorm:"column:protein"`
	Fat          float64 `json:"fat" gorm:"column:fat"`
}
