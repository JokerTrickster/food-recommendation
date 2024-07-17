package request

type ReqUpdateUser struct {
	Birth string `json:"birth" validate:"required"`
	Sex   string `json:"sex" validate:"required"`
}
