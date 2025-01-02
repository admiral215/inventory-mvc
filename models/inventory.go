package models

import "gorm.io/gorm"

type Inventory struct {
	gorm.Model
	Name        string `json:"name" gorm:"not null"`
	Quantity    int    `json:"quantity" gorm:"not null"`
	Description string `json:"description"`

	CategoryId uint     `json:"category_id" gorm:"not null"`
	Category   Category `json:"category" gorm:"foreignKey:CategoryId"`
}
