package usecase

import (
	"main/features/user/model/entity"
	"main/utils/db/mysql"

	"gorm.io/gorm"
)

func CreateUpdateUserDTO(entity *entity.UpdateUserEntity) (*mysql.Users, error) {
	//유저 정보를 업데이트 할 때 사용할 DTO를 생성한다.
	return &mysql.Users{
		Model: gorm.Model{
			ID: entity.UserID,
		},
		Email: entity.Email,
		Sex:   entity.Sex,
		Birth: entity.Birth,
	}, nil

}
