package models

type Todo struct {
	ID    string `gorm:"primaryKey" json:"id"`
	Title string `json:"title"`
	Done  bool   `json:"done"`
}
