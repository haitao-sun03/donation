package vo

type NonceVO struct {
	Address string `json:"currentAccount" binding:"required"`
}

type LoginVO struct {
	Address   string `json:"currentAccount" binding:"required"`
	Signature string `json:"signature" binding:"required"`
}

type RenewVO struct {
	Address    string `json:"currentAccount" binding:"required"`
	RefreshJwt string `json:"refreshJwt" binding:"required"`
}

type RefeshJwtVO struct {
	Address string `json:"currentAccount" binding:"required"`
}

type UserRoleVO struct {
	Address    string `json:"address" binding:"required"`
	Role       string `json:"role" binding:"required"`
	CampaignId int64  `json:"campaignId" binding:"required"`
}
