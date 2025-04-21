package routers

import (
	"net/http/pprof"

	"github.com/gin-gonic/gin"
	"github.com/haitao-sun03/donation/backend-end/controllers"
)

func Router() *gin.Engine {
	r := gin.Default()
	UseMiddleWare(r)
	initPprofFun(r)

	core := r.Group("/core")
	campaign := core.Group("/campaign")
	campaign.POST("/list", controllers.CampaignController{}.GetList)

	donation := core.Group("/donation")
	donation.POST("/list", controllers.DonationController{}.GetDonationsOfCampaign)

	nft := core.Group("/nft")
	nft.POST("/nftGroupByUser", controllers.NftController{}.GetNfts)

	return r
}

func initPprofFun(r *gin.Engine) {
	// 注册pprof处理函数
	r.GET("/debug/pprof/", gin.WrapF(pprof.Index))
	r.GET("/debug/pprof/cmdline", gin.WrapF(pprof.Cmdline))
	r.GET("/debug/pprof/profile", gin.WrapF(pprof.Profile))
	r.GET("/debug/pprof/symbol", gin.WrapF(pprof.Symbol))
	r.GET("/debug/pprof/trace", gin.WrapF(pprof.Trace))
	r.POST("/debug/pprof/symbol", gin.WrapF(pprof.Symbol))
}
