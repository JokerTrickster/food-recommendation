package usecase

import (
	"main/features/user/model/entity"
	"main/features/user/model/response"
	"main/utils/db/mysql"

	"gorm.io/gorm"
)

func CreateUpdateUserDTO(entity *entity.UpdateUserEntity) (*mysql.Users, error) {
	//유저 정보를 업데이트 할 때 사용할 DTO를 생성한다.
	result := &mysql.Users{
		Model: gorm.Model{
			ID: entity.UserID,
		},
	}
	if entity.Birth != "" {
		result.Birth = entity.Birth
	}
	if entity.Name != "" {
		result.Name = entity.Name
	}
	if entity.Sex != "" {
		result.Sex = entity.Sex
	}
	if entity.Email != "" {
		result.Email = entity.Email
	}
	if entity.PrevPassword != "" && entity.NewPassword != "" {
		result.Password = entity.NewPassword
	}
	return result, nil
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
