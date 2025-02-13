package vo

type Pagination struct {
	Page     int `json:"page" binding:"required"`
	PageSize int `json:"pageSize" binding:"required"`
}

type PageCampaign struct {
	Pagination
	Starter string `json:"starter"`
	Status  int    `json:"status"`
}
