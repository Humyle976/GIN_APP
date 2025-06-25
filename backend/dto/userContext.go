package dto

type UserContext struct {
	ID          uint   `json:"user_id" binding:"required"`
	Fullname    string `json:"fullname" binding:"required"`
	CountryCode string `gorm:"size:2" json:"country_code"`
}