package service

import (
	"strconv"

	"github.com/haitao-sun03/donation/backend-end/db"
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
