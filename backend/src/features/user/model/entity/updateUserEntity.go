package entity

type UpdateUserEntity struct {
	UserID uint   `json:"user_id"`
	Birth  string `json:"birth"`
	Sex    string `json:"sex"`
	Email  string `json:"email"`
}
