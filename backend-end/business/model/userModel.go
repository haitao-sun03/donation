package model

import "gorm.io/gorm"

type UserModel struct {
	gorm.Model
	Address string `gorm:"type:varchar(255);not null"`
	Role    string `gorm:"type:varchar(255);;default:'user'"`
}

func (UserModel) TableName() string {
	return "users"
}

const (
	UserRole    string = "user"
	StarterRole string = "starter"
	DonorRole   string = "donor"
)
