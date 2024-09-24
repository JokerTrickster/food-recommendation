package usecase

import (
	"main/features/user/model/entity"
	"main/features/user/model/response"
	"main/utils/db/mysql"

	"gorm.io/gorm"
)

func CreateUpdateUserDTO(entity *entity.UpdateUserEntity) (*mysql.Users, error) {
	//유저 정보를 업데이트 할 때 사용할 DTO를 생성한다.
	return &mysql.Users{
		Model: gorm.Model{
			ID: entity.UserID,
		},
		Name:  entity.Name,
		Email: entity.Email,
		Sex:   entity.Sex,
		Birth: entity.Birth,
	}, nil

}

func CreateResGetUser(user *mysql.Users) response.ResGetUser {
	//유저 정보를 가져올 때 사용할 DTO를 생성한다.
	return response.ResGetUser{
		Name:  user.Name,
		Email: user.Email,
		Sex:   user.Sex,
		Birth: user.Birth,
	}
}
