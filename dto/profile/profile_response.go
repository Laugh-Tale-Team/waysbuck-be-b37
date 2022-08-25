package profiledto

import "waysbuck/models"

type ProfileResponse struct {
	ID         int    `json:"id" gorm:"primary_key:auto_increment"`
	Fullname   string `json:"fullname" gorm:"type: varchar(255)"`
	Image      string `json:"image" gorm:"type: varchar(255)"`
	Address    string `json:"address" gorm:"type: varchar(255)"`
	City       string `json:"city" form:"city" validate:"required"`
	PostalCode int    `json:"postal_code" form:"postal_code" validate:"required"`
	Phone      string `json:"phone" gorm:"type: varchar(255)"`
	UserID     int    `json:"user_id"`
	User       models.UserProfileResponse  `json:"user"`
}