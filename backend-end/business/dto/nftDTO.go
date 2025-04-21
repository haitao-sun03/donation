package dto

type NftLevelDTO struct {
	NftLevel string
	MetaData string
	Nfts     []Nft
}

type Nft struct {
	CampaignId uint
}
