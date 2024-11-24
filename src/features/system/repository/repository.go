package repository

import "gorm.io/gorm"

type ReportSystemRepository struct {
	GormDB *gorm.DB
}

type FoodReportSystemRepository struct {
	GormDB *gorm.DB
}