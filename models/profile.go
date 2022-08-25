package models

import "time"

type Profile struct {
	ID 			int						`json:"id" gorm:"primary_key:auto_increment"`
	Fullname 	string					`json:"fullname" gorm:"type: varchar(255)"`
	Image		string					`json:"image" gorm:"type: varchar(255)"`
	Address		string					`json:"address" gorm:"type: varchar(255)"`
	Phone		string					`json:"phone" gorm:"type: varchar(255)"`
	City       	string      			`json:"city" gorm:"type: varchar(255)"`
	PostalCode 	int        				`json:"postal_code" gorm:"type: int"`
	UserID		int						`json:"user_id"`
	User		UserProfileResponse 	`json:"user" gorm:"constraint:OnUpdate:CASCADE,OnDelete:CASCADE;"`
	CreatedAt	time.Time				`json:"-"`
	UpdatedAt 	time.Time 				`json:"-"`
	
}

type ProfileResponse struct {
	Fullname	string	`json:"fullname"`
	Image		string	`json:"image"`
	Address		string	`json:"address"`
	Phone		string	`json:"phone"`
	City       	string 	`json:"city" `
	PostalCode 	int    	`json:"postal_code"`
	UserID		int		`json:"-"`
}

func (ProfileResponse) TableName() string {
	return "profiles"
}