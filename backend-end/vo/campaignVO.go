package vo

type PageCampaign struct {
	Pagination
	Starter string `json:"starter"`
	Status  int    `json:"status"`
}
