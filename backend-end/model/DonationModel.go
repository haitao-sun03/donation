package model

import (
	"time"

	"gorm.io/gorm"
)

// 捐赠记录表结构定义（与DonationRecord结构体映射）
type DonationModel struct {
	gorm.Model
	CampaignID string    `gorm:"type:varchar(64);index:idx_campaign_donor,unique"` // 唯一索引防重复
	Donor      string    `gorm:"type:varchar(42);index:idx_campaign_donor,unique"` // 以太坊地址长度: Ethereum addresses are typically 42 characters long including '0x'
	Amount     string    `gorm:"type:varchar(78)"`                                 // 支持大整数（最长78位）
	BlockTime  time.Time `gorm:"index"`
	IsRefund   uint8     `gorm:"type:int;not null;default:0"`
	MintLevel  string    `gorm:"type:varchar(255)"`
}

func (DonationModel) TableName() string {
	return "donation_records"
}

// 是否退款
const (
	NotRefund uint8 = 0
	Refund    uint8 = 1
)

// MintLevel
const (
	Diamond string = "Diamond"
	Gold    string = "Gold"
	Silver  string = "Silver"
	None    string = "None"
)
