package repository

import "gorm.io/gorm"

type GetUserRepository struct {
	GormDB *gorm.DB
}
type UpdateUserRepository struct {
	GormDB *gorm.DB
}

type DeleteUserRepository struct {
	GormDB *gorm.DB
}
