package models

import (
	_ "github.com/jinzhu/gorm/dialects/mysql"
)

type Food struct {
	ID    uint    `gorm:"primary_key;AUTO_INCREMENT" json:"_id,omitempty"`
	Name  string  `json:"name,omitempty"`
	Price float64 `json:"price,omitempty"`
}

// FoodImage string  `json:"food_image,omitempty"`
// FoodID    uint    `json:"food_id,omitempty" gorm:"unique;not null"`
// FoodID    uint       `json:"food_id,omitempty" gorm:"index"` // foreign key
// MenuID    string  `json:"menu_id,omitempty"`