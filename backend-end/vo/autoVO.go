package vo

type NonceVO struct {
	Address string `json:"currentAccount" binding:"required"`
}

type LoginVO struct {
	Address   string `json:"currentAccount" binding:"required"`
	Signature string `json:"signature" binding:"required"`
}

type RefeshJwtVO struct {
	Address string `json:"currentAccount" binding:"required"`
}
