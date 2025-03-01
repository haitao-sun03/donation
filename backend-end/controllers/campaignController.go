package controllers

import (
	"net/http"

	"github.com/gin-gonic/gin"
	"github.com/haitao-sun03/go/common"
	"github.com/haitao-sun03/go/db"
	"github.com/haitao-sun03/go/vo"
)

type CampaignController struct{}

func (u CampaignController) GetList(ctx *gin.Context) {
	var total int64

	pageCampaign := vo.PageCampaign{}

	if err := ctx.ShouldBindJSON(&pageCampaign); err != nil {
		common.Fail(ctx, http.StatusInternalServerError, err)
		return
	}

	campaignDao := db.NewCampaignDao()
	campaigns, total, err := campaignDao.GetCampaignList(pageCampaign)

	if err != nil {
		common.Fail(ctx, http.StatusInternalServerError, err)
		return
	}

	common.Success(ctx, http.StatusOK, common.MessageSuccess, campaigns, total)
}
