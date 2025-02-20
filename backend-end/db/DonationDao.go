package db

import (
	"strconv"

	"github.com/haitao-sun03/go/common"
	"github.com/haitao-sun03/go/config"
	"github.com/haitao-sun03/go/dto"
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

func (donationDao *DonationDao) GetNftOfUser(nftVO vo.NftVO) ([]dto.NftLevelDTO, error) {
	var donations []model.DonationModel
	var metadatas []model.NFTMetaDataMedal
	var nftDTOs []dto.NftLevelDTO

	// 查询用户的所有捐赠记录，且已铸造NFT的
	err := donationDao.db.Where("donor = ? AND mint_level != ''", nftVO.User).
		Find(&donations).Error
	if err != nil {
		return nil, err
	}

	// 获取所有NFT元数据
	err = donationDao.db.Find(&metadatas).Error
	if err != nil {
		return nil, err
	}

	// 构建元数据映射
	metaMap := make(map[string]string)
	for _, meta := range metadatas {
		metaMap[meta.Level] = meta.URI
	}

	// 按NFT等级分组的映射
	nftLevelMap := make(map[string]*dto.NftLevelDTO)

	// 组装DTO
	for _, donation := range donations {
		level := donation.MintLevel
		// 如果这个等级的NFT还没有创建，就创建一个新的
		if _, exists := nftLevelMap[level]; !exists {
			nftLevelMap[level] = &dto.NftLevelDTO{
				NftLevel: level,
				MetaData: metaMap[level],
				Nfts:     []dto.Nft{},
			}
		}

		// 添加这个捐赠对应的NFT
		campaignID, _ := strconv.ParseUint(donation.CampaignID, 10, 64)
		nftLevelMap[level].Nfts = append(nftLevelMap[level].Nfts, dto.Nft{
			CampaignId: uint(campaignID),
		})
	}

	// 将map转换为slice
	for _, nftDTO := range nftLevelMap {
		nftDTOs = append(nftDTOs, *nftDTO)
	}

	return nftDTOs, nil
}
