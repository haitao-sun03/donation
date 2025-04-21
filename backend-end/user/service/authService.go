package service

import (
	"strconv"

	"github.com/haitao-sun03/donation/backend-end/user/db"
)

type AuthService struct{}

func NewAuthService() AuthService {
	return AuthService{}
}

func (AuthService) GetRolesByAddress(address string) (map[string]string, bool) {
	userCampaignIdRoleDao := db.NewUserCampaignIdRoleDao()
	userActivityRoles, err := userCampaignIdRoleDao.GetUserCampaignRolesByUser(address)
	if err != nil {
		return nil, true
	}

	roles := make(map[string]string)
	for _, userActivityRole := range userActivityRoles {
		campaignIdStr := strconv.FormatUint(uint64(userActivityRole.CampaignId), 10)
		roles[campaignIdStr] = userActivityRole.Role
	}
	return roles, false
}

func (AuthService) GetUserRole(address string, role string, campaignId int64) bool {
	userCampaignIdRoleDao := db.NewUserCampaignIdRoleDao()
	return userCampaignIdRoleDao.GetUserRole(address, role, campaignId)

}

func (AuthService) AddUserRole(address string, role string, campaignId int64) error {
	userCampaignIdRoleDao := db.NewUserCampaignIdRoleDao()
	err := userCampaignIdRoleDao.AddUserRole(address, role, campaignId)
	if err != nil {
		return err
	}

	return nil
}
