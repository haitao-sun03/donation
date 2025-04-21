package vo

type PageDonation struct {
	Pagination
	CampaignId uint `json:"campaignId" binding:"required"`
}
