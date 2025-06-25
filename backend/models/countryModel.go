package models

type Country struct {
	Code  string `gorm:"primaryKey;size:2" json:"code"`
	Name  string `gorm:"size:100;not null" json:"name"`
}
