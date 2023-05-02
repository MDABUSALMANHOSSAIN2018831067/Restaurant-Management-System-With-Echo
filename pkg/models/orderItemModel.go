package models

type OrderItem struct {
	ID       uint `json:"_id,omitempty"`
	Quantity uint `json:"quantity,omitempty"`
	MenuID   uint `json:"menu_id,omitempty"`
	Menu     Menu `gorm:"foreignKey:MenuID;constraint:OnDelete:CASCADE"`
}

// OrderItemID string   `json:"order_item_id,omitempty"`
// 	OrderID     string   `json:"order_id,omitempty"`
//TotalPrice float64 `json:"total_price,omitempty"`
