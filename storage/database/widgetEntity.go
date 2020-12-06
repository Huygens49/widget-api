package database

import "gorm.io/gorm"

type WidgetEntity struct {
	gorm.Model
	Description string
	Owner       string
}

func (WidgetEntity) TableName() string {
	return "widgets"
}
