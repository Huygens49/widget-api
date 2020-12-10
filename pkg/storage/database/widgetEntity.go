package database

import "gorm.io/gorm"

type WidgetEntity struct {
	gorm.Model
	Description string
	Owner       string
	Value       int
}

func (WidgetEntity) TableName() string {
	return "widgets"
}
