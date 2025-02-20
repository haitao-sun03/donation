package controllers

import (
	"net/http"

	"github.com/gin-gonic/gin"
	"github.com/haitao-sun03/go/common"
	"github.com/haitao-sun03/go/db"
	"github.com/haitao-sun03/go/vo"
)

type NftController struct{}

func (n NftController) GetNfts(ctx *gin.Context) {
	nftVO := vo.NftVO{}

	if err := ctx.ShouldBindJSON(&nftVO); err != nil {
		ctx.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}

	donationDao := db.NewDonationDao()
	nftLevels, err := donationDao.GetNftOfUser(nftVO)

	if err != nil {
		common.Fail(ctx, http.StatusInternalServerError, err)
		return
	}

	common.Success(ctx, http.StatusOK, common.MessageSuccess, nftLevels, -1)
}
