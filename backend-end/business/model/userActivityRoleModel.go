package model

type UserActivityRoleModel struct {
	// gorm.Model
	Address    string `gorm:"type:varchar(42);not null;primaryKey"`
	CampaignId uint   `gorm:"type:bigint;primaryKey"`
	Role       string `gorm:"type:varchar(20);default:'user';primaryKey"`
}

func (UserActivityRoleModel) TableName() string {
	return "user_activity_roles"
}
