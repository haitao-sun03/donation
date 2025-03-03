package db

import (
	"github.com/haitao-sun03/donation/backend-end/common"
	"github.com/haitao-sun03/donation/backend-end/config"
	"github.com/haitao-sun03/donation/backend-end/model"
	"github.com/haitao-sun03/donation/backend-end/vo"
	"gorm.io/gorm"
)

type CampaignDao struct {
	db *gorm.DB
}

func NewCampaignDao() *CampaignDao {
	db := config.GetDB()
	return &CampaignDao{db}
}

func (campaignDao *CampaignDao) GetCampaignList(pageCampaign vo.PageCampaign) ([]model.CampaignModel, int64, error) {
	var campaigns []model.CampaignModel
	var total int64

	// 构建查询条件
	dbQuery := campaignDao.db.Model(&model.CampaignModel{})

	// 根据Status和Starter添加条件
	if pageCampaign.Status != -1 { // -1表示所有状态
		dbQuery = dbQuery.Where("status = ?", pageCampaign.Status)
	}
	if pageCampaign.Starter != "" {
		dbQuery = dbQuery.Where("starter = ?", pageCampaign.Starter)
	}

	// 计算总数
	if err := dbQuery.Count(&total).Error; err != nil {
		return nil, 0, err
	}

	// 使用封装的分页函数
	common.PaginationFunc(dbQuery, pageCampaign.Page, pageCampaign.PageSize)

	// 执行查询
	if err := dbQuery.Find(&campaigns).Error; err != nil {
		return nil, 0, err
	}

	return campaigns, total, nil
}
