package di

import "gorm.io/gorm"

type DependencyInjection struct {
	DB *gorm.DB
}

func InitDI(db *gorm.DB) *DependencyInjection {
	return &DependencyInjection{
		DB: db,
	}
}
