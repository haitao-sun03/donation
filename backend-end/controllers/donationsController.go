package controllers

import (
	"net/http"

	"github.com/gin-gonic/gin"
	"github.com/haitao-sun03/donation/backend-end/common"
	"github.com/haitao-sun03/donation/backend-end/db"
	"github.com/haitao-sun03/donation/backend-end/vo"
)

type DonationController struct{}

func (d DonationController) GetDonationsOfCampaign(ctx *gin.Context) {
	var total int64

	pageDonations := vo.PageDonation{}

	if err := ctx.ShouldBindJSON(&pageDonations); err != nil {
		ctx.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}

	donationDao := db.NewDonationDao()
	donations, total, err := donationDao.GetDonationListOfCampaign(pageDonations)

	if err != nil {
		common.Fail(ctx, http.StatusInternalServerError, err)
		return
	}

	common.Success(ctx, http.StatusOK, common.MessageSuccess, donations, total)
}
