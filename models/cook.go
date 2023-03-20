package models

type Cook struct {
	ID         uint   `json:"id" gorm:"primary_key"`
	Name       string `json:"name"`
	Ingredient string `json:"ingredient"`
}
