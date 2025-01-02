package models

import "gorm.io/gorm"

type Category struct {
	gorm.Model
	Name string `json:"name" gorm:"uniqueIndex"`

	Inventory []Inventory `json:"inventory" gorm:"foreignKey:ID"`
}
