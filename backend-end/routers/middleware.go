package routers

import (
	"strings"

	"github.com/gin-gonic/gin"
	"github.com/haitao-sun03/donation/backend-end/config"
	"github.com/haitao-sun03/donation/backend-end/utils"
	log "github.com/sirupsen/logrus"
)

// 定义中间
func middleWare() gin.HandlerFunc {

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

func authMiddleware() gin.HandlerFunc {
	return func(c *gin.Context) {

		for _, whiteUrl := range config.Config.Jwt.WhiteList {
			url := c.Request.URL.Path
			if strings.Contains(url, whiteUrl) {
				c.Next()
				return
			}
		}
		tokenString := c.GetHeader("Authorization")
		log.Infoln("Authorization:", tokenString)
		if tokenString == "" {
			c.JSON(401, gin.H{"error": "missing token"})
			c.Abort()
			return
		}
		address, roles, _, err := utils.VerifyJWT(tokenString)
		if err != nil {
			c.JSON(401, gin.H{"error": "invalid token"})
			c.Abort()
			return
		}
		log.Infoln("address:", address)
		log.Infoln("roles:", roles)
		c.Set("address", address)
		c.Set("roles", roles)
		c.Next()
	}
}

func UseMiddleWare(r *gin.Engine) {
	r.Use(middleWare())
	r.Use(authMiddleware())
}
