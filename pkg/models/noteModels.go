package models

type Note struct {
	ID     uint   `json:"_id,omitempty"`
	Text   string `json:"text,omitempty"`
	Title  string `json:"title,omitempty"`
	NoteID string `json:"note_id,omitempty"`
}
