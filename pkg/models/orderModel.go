package models

import "time"

type Order struct {
	ID        uint      `json:"_id,omitempty"`
	OrderDate time.Time `json:"order_date,omitempty"`
	OrderID   string    `json:"order_id,omitempty"`
	TableID   *string   `json:"table_id,omitempty"`
}
