package model

import "gorm.io/gorm"

type NFTMetaDataMedal struct {
	gorm.Model
	Level string `gorm:"type:varchar(255);not null"`
	URI   string `gorm:"type:varchar(255);not null"`
}

func (NFTMetaDataMedal) TableName() string {
	return "nft_meta_data"
}
