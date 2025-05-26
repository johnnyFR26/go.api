package models

type User struct {
	ID    string `gorm:"primaryKey" json:"id"`
	Name  string `json:"name"`
	Tasks []Task `gorm:"foreignKey:UserID" json:"tasks"`
}
