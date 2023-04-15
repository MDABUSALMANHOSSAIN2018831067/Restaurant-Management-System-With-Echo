package models

type Table struct {
	ID             uint   `json:"_id,omitempty"`
	NumberOfGuests *int   `json:"number_of_guests,omitempty"`
	TableNumber    *int   `json:"table_number,omitempty"`
	TableID        string `json:"table_id,omitempty"`
}
