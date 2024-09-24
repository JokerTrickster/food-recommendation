package mysql

import "gorm.io/gorm"

// 알레르기 상수 정의
const (
	AllergyEtc = 1 + iota
	AllergyEgg
	AllergyMilk
	AllergyBuckwheat
	AllergyPeanut
	AllergySoybean
	AllergyWheat
)

// 시나리오 상수 정의
// 연인, 혼반, 가족, 다이어트, 회식, 친구
const (
	ScenarioCouple = 1 + iota
	ScenarioSolo
	ScenarioFamily
	ScenarioDiet
	ScenarioCompany
	ScenarioFriend
)

// 식사 시간 상수 정의
// 아침, 점심, 저녁, 브런치, 간식, 야식
const (
	TimeMorning  = 1 + iota //아침
	TimeLunch               //점심
	TimeDinner              //저녁
	TimeBrunch              //브런치
	TimeSnack               //간식
	TimeMidnight            //야식
)

// 음식 종류 상수 정의
// 한식, 중식, 일식, 양식, 분식,베트남 음식, 인도 음식, 패스트 푸드, 디저트, 퓨전 요리
const (
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
	ThemeStressRelief    = 1 + iota // 스트레스 해소
	ThemeFatigueRecovery            // 피로 회복
	ThemeMoodChange                 //기분 전환
	ThemeSeasonalFood               //제철 음식
	ThemeNutrition                  //영양식
	ThemeSpecialDay                 //특별한 날
)

// 맛 상수 정의
// 매운맛, 감칠맛, 고소한맛, 단맛, 짠맛, 싱거운맛
const (
	FlavorSpicy  = 1 + iota //매운맛
	FlavorSavory            //감칠맛
	FlavorNutty             //고소한맛
	FlavorSweet             //단맛
	FlavorSalty             //짠맛
	FlavorSour              //싱거운맛
)

var AllergyMap = map[string]int{
	"기타": AllergyEtc,
	"계란": AllergyEgg,
	"우유": AllergyMilk,
	"메밀": AllergyBuckwheat,
	"땅콩": AllergyPeanut,
	"대두": AllergySoybean,
	"밀":  AllergyWheat,
}

// 맵 정의
var ScenarioMap = map[string]int{
	"연인":   ScenarioCouple,
	"혼밥":   ScenarioSolo,
	"가족":   ScenarioFamily,
	"다이어트": ScenarioDiet,
	"회식":   ScenarioCompany,
	"친구":   ScenarioFriend,
}

var TimeMap = map[string]int{
	"아침":  TimeMorning,
	"점심":  TimeLunch,
	"저녁":  TimeDinner,
	"브런치": TimeBrunch,
	"간식":  TimeSnack,
	"야식":  TimeMidnight,
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
	"짠맛":   FlavorSalty,
	"싱거운맛": FlavorSour,
}

var ThemeMap = map[string]int{
	"스트레스 해소": ThemeStressRelief,
	"피로 회복":   ThemeFatigueRecovery,
	"기분 전환":   ThemeMoodChange,
	"제철 음식":   ThemeSeasonalFood,
	"영양식":     ThemeNutrition,
	"특별한 날":   ThemeSpecialDay,
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
	Name       string `json:"name" gorm:"column:name"`
	Image      string `json:"image" gorm:"column:image"`
	ScenarioID int    `json:"scenarioID" gorm:"column:scenario_id"`
	TimeID     int    `json:"timeID" gorm:"column:time_id"`
	TypeID     int    `json:"typeID" gorm:"column:type_id"`
	ThemeID    int    `json:"themeID" gorm:"column:theme_id"`
	FlavorID   int    `json:"flavorID" gorm:"column:flavor_id"`
}
type FoodHistory struct {
	gorm.Model
	UserID uint `json:"userID" gorm:"column:user_id"`
	FoodID uint `json:"foodID" gorm:"column:food_id"`
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

// 알레르기 정보
type Allergies struct {
	gorm.Model
	Name string `json:"name" gorm:"column:name"`
}

// 유저 알레르기 정보
type UserAllergies struct {
	gorm.Model
	UserID    uint `json:"userID" gorm:"column:user_id"`
	AllergyID uint `json:"allergyID" gorm:"column:allergy_id"`
}