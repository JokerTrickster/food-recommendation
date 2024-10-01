package repository

import "gorm.io/gorm"

type ReportSystemRepository struct {
	GormDB *gorm.DB
}
