package db

import (
	"github.com/haitao-sun03/go/common"
	"github.com/haitao-sun03/go/config"
	"github.com/haitao-sun03/go/model"
	"github.com/haitao-sun03/go/vo"
	"gorm.io/gorm"
)

type DonationDao struct {
	db *gorm.DB
}

func NewDonationDao() *DonationDao {
	db := config.GetDB()
	return &DonationDao{db}
}

func (donationDao *DonationDao) GetDonationListOfCampaign(pageDonation vo.PageDonation) ([]model.DonationModel, int64, error) {
	var donations []model.DonationModel
	var total int64

	// 构建查询条件
	dbQuery := donationDao.db.Model(&model.DonationModel{})

	// 根据Status和Starter添加条件
	dbQuery = dbQuery.Where("campaign_id = ?", pageDonation.CampaignId)

	// 计算总数
	if err := dbQuery.Count(&total).Error; err != nil {
		return nil, 0, err
	}

	// 使用封装的分页函数
	common.PaginationFunc(dbQuery, pageDonation.Page, pageDonation.PageSize)

	// 执行查询
	if err := dbQuery.Find(&donations).Error; err != nil {
		return nil, 0, err
	}

	return donations, total, nil
}
