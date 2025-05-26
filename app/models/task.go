package models

type Task struct {
	ID     string `gorm:"primaryKey" json:"id"`
	Title  string `json:"title"`
	Done   bool   `json:"done"`
	UserID string `json:"user_id"`
}
