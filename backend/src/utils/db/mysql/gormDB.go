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
const (
	ScenarioAll = 1 + iota
	ScenarioAlone
	ScenarioFriends
	ScenarioLovers
	ScenarioFamily
	ScenarioWork
)

// 식사 시간 상수 정의
const (
	TimeAll = 1 + iota
	TimeBreakfast
	TimeLunch
	TimeDinner
	TimeSnack
	TimeLateNight
)

// 음식 종류 상수 정의
const (
	TypeAll = 1 + iota
	TypeKorean
	TypeChinese
	TypeWestern
	TypeJapanese
	TypeSnack
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
	"전체": ScenarioAll,
	"혼밥": ScenarioAlone,
	"친구": ScenarioFriends,
	"연인": ScenarioLovers,
	"가족": ScenarioFamily,
	"회식": ScenarioWork,
}

var TimeMap = map[string]int{
	"전체": TimeAll,
	"조식": TimeBreakfast,
	"중식": TimeLunch,
	"석식": TimeDinner,
	"간식": TimeSnack,
	"야식": TimeLateNight,
}

var TypeMap = map[string]int{
	"전체": TypeAll,
	"한식": TypeKorean,
	"중식": TypeChinese,
	"양식": TypeWestern,
	"일식": TypeJapanese,
	"분식": TypeSnack,
}

// 반대 맵 정의 (int -> string)
var ScenarioReverseMap = make(map[int]string)
var TimeReverseMap = make(map[int]string)
var TypeReverseMap = make(map[int]string)

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
