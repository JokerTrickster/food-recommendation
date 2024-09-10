package entity

import "mime/multipart"

type ImageUploadFoodEntity struct {
	FoodID int                   `json:"foodID"`
	Image  *multipart.FileHeader `json:"image"`
}
