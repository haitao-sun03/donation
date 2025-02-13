package routers

import (
	"net/http/pprof"

	"github.com/gin-gonic/gin"
	"github.com/haitao-sun03/go/controllers"
	log "github.com/sirupsen/logrus"
)

// 定义中间
func MiddleWare() gin.HandlerFunc {

	return func(c *gin.Context) {
		defer func() {
			if err := recover(); err != nil {
				if errs, ok := err.(error); ok {
					log.WithError(errs).Error("global error occurred")
				} else {
					log.WithField("panic", errs).Error("global error occurred")
				}

			}
		}()
		c.Next()
	}
}

func Router() *gin.Engine {
	r := gin.Default()
	// r.Use(MiddleWare())
	initPprofFun(r)

	campaign := r.Group("/campaign")
	campaign.POST("/list", controllers.CampaignController{}.GetList)

	donation := r.Group("/donation")
	donation.POST("/list", controllers.DonationController{}.GetDonationsOfCampaign)
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
