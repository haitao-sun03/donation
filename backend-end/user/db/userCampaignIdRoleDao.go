package db

import (
	"errors"

	"github.com/haitao-sun03/donation/backend-end/user/config"
	"github.com/haitao-sun03/donation/backend-end/user/model"
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

func (userCampaignIdRoleDao *UserCampaignIdRoleDao) GetUserRole(address string, role string, campaignId int64) bool {
	var userActivityRoleModel model.UserActivityRoleModel
	result := userCampaignIdRoleDao.db.
		Where("address = ?", address).
		Where("campaign_id = ?", campaignId).
		Where("role = ?", role).
		First(&userActivityRoleModel)
	// 若该用户address之前没有对campaign_id进行捐赠
	if errors.Is(result.Error, gorm.ErrRecordNotFound) || result.RowsAffected == 0 {
		return false
	}
	return true
}

func (userCampaignIdRoleDao *UserCampaignIdRoleDao) AddUserRole(address string, role string, campaignId int64) error {
	if res := userCampaignIdRoleDao.db.Exec("INSERT INTO user_activity_roles (address, campaign_id, role) VALUES (?, ?, ?)",
		address,
		campaignId,
		role); res.Error != nil {
		return res.Error
	}
	return nil
}
