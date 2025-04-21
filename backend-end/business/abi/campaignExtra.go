package abi

import (
	"math/big"

	"github.com/haitao-sun03/donation/backend-end/model"
)

//go:generate abigen --abi .\donation.abi --type donationManage --pkg abi --out .\donationManage.go
//go:generate abigen --abi .\nft.abi --type nft --pkg abi --out .\nft.go

type CampaignEvent interface {
	GetCampaignID() *big.Int
	GetStatus() uint8
}

func (e *DonationManageActiveCampaign) GetCampaignID() *big.Int { return e.Id }
func (e *DonationManageActiveCampaign) GetStatus() uint8        { return model.CampaignStatusActive }

func (e *DonationManageCancellCampaign) GetCampaignID() *big.Int { return e.Id }
func (e *DonationManageCancellCampaign) GetStatus() uint8        { return model.CampaignStatusCancelled }

func (e *DonationManageCompletedCampaign) GetCampaignID() *big.Int { return e.Id }
func (e *DonationManageCompletedCampaign) GetStatus() uint8        { return model.CampaignStatusCompleted }
