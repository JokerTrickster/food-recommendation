package mysql

import "gorm.io/gorm"

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
	Sex      string `json:"sex" gorm:"column:sex"`
	Provider string `json:"provider" gorm:"column:provider"`
}

type Foods struct {
	gorm.Model
	Name       string `json:"name" gorm:"column:name"`
	UserID     int    `json:"userID" gorm:"column:user_id"`
	ScenarioID int    `json:"scenarioID" gorm:"column:scenario_id"`
	TimeID     int    `json:"timeID" gorm:"column:time_id"`
	TypeID     int    `json:"typeID" gorm:"column:type_id"`
}

type MetaTables struct {
	gorm.Model
	TableName        string `json:"tableName" gorm:"column:table_name"`
	TableDescription string `json:"tableDescription" gorm:"column:table_description"`
}

//전체, 혼밥, 친구, 가족, 회식, 연인, 기타
type Scenarios struct {
	gorm.Model
	Name        string `json:"name" gorm:"column:name"`
	Description string `json:"description" gorm:"column:description"`
}

//전체, 아침, 점심, 저녁, 간식, 야식
type Times struct {
	gorm.Model
	Name        string `json:"name" gorm:"column:name"`
	Description string `json:"description" gorm:"column:description"`
}

//전체, 한식, 중식, 일식, 양식, 분식, 패스트푸드, 카페, 술집, 기타
type Types struct {
	gorm.Model
	Name        string `json:"name" gorm:"column:name"`
	Description string `json:"description" gorm:"column:description"`
}
type UserAuths struct {
	gorm.Model
	Email    string `json:"email" gorm:"column:email"`
	AuthCode string `json:"authCode" gorm:"column:auth_code"`
	Type     string `json:"type" gorm:"column:type"`
}
