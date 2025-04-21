package vo

type NftVO struct {
	User string `json:"user" binding:"required"`
}
