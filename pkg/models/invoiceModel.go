package models

import "time"

type Invoice struct {
	ID             uint      `json:"_id,omitempty"`
	InvoiceID      string    `json:"invoice_id,omitempty"`
	OrderID        string    `json:"order_id,omitempty"`
	PaymentMethod  string    `json:"payment_method,omitempty"`
	PaymentStatus  string    `json:"payment_status,omitempty"`
	PaymentDueDate time.Time `json:"payment_due_date,omitempty"`
}
