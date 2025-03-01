package model

import "gorm.io/gorm"

type NFTMetaDataModel struct {
	gorm.Model
	Level string `gorm:"type:varchar(255);not null"`
	URI   string `gorm:"type:varchar(255);not null"`
}

func (NFTMetaDataModel) TableName() string {
	return "nft_meta_data"
}
