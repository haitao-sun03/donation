package model

import (
	"time"

	"gorm.io/gorm"
)

// 定义活动的状态
const (
	CampaignStatusPending   uint8 = 0
	CampaignStatusActive    uint8 = 1
	CampaignStatusCompleted uint8 = 2
	CampaignStatusCancelled uint8 = 3
)

// 定义活动是否被withdraw
const (
	NotWithdraw uint8 = 0
	Withdraw    uint8 = 1
)

// Campaign 代表一个募捐活动
type CampaignModel struct {
	gorm.Model
	CampaignID     string    `gorm:"type:varchar(64);index:idx_campaign,unique"` // 唯一索引防重复
	Title          string    `gorm:"type:varchar(255);not null"`
	Goal           string    `gorm:"type:varchar(64);not null"`             // 存储big.Int的字符串表示
	TotalDonated   string    `gorm:"type:varchar(64);not null;default:'0'"` // 存储big.Int的字符串表示
	StartTime      time.Time `gorm:"not null"`
	EndTime        time.Time `gorm:"not null"`
	Status         uint8     `gorm:"type:int;not null"`
	Starter        string    `gorm:"type:varchar(42);not null"` // Ethereum addresses are typically 42 characters long including '0x'
	IsWithdraw     uint8     `gorm:"type:int;not null;default:0"`
	Nature         uint8     `gorm:"type:int;not null"`
	Beneficiary    uint8     `gorm:"type:int;not null"`
	Purpose        string    `gorm:"type:varchar(255);not null"`
	ExpectedImpact string    `gorm:"type:varchar(255);not null"`
}

// 表初始化，迁移时使用
func (CampaignModel) TableName() string {
	return "campaigns"
}
