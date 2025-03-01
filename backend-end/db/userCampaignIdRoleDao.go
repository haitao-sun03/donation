package db

import (
	"github.com/haitao-sun03/go/config"
	"github.com/haitao-sun03/go/model"
	"gorm.io/gorm"
)

type UserCampaignIdRoleDao struct {
	db *gorm.DB
}

func NewUserCampaignIdRoleDao() *UserCampaignIdRoleDao {
	db := config.GetDB()
	return &UserCampaignIdRoleDao{db}
}

func (userCampaignIdRoleDao *UserCampaignIdRoleDao) GetUserCampaignRolesByUser(user string) ([]model.UserActivityRoleModel, error) {
	var userActivityRoles []model.UserActivityRoleModel
	if err := userCampaignIdRoleDao.db.Where("address = ?", user).Find(&userActivityRoles).Error; err != nil {
		return nil, err
	}
	return userActivityRoles, nil
}
