package models

import (
	_ "github.com/jinzhu/gorm/dialects/mysql"
)

type Menu struct {
	ID       uint   `gorm:"primary_key;AUTO_INCREMENT" json:"_id,omitempty"`
	Category string `json:"category,omitempty"`
	FoodID   uint   `json:"food_id"`
	Food     Food   `gorm:"foreignKey:FoodID;constraint:OnDelete:CASCADE"`
}

// UserID    string     `json:"user_id"`
// Name      string     `json:"name,omitempty"`
// StartDate *time.Time `json:"start_date,omitempty"`
// EndDate   *time.Time `json:"end_date,omitempty"`
