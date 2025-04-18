package controllers

import (
	"context"
	"errors"
	"net/http"
	"strconv"

	"github.com/gin-gonic/gin"
	"github.com/haitao-sun03/donation/backend-end/common"
	"github.com/haitao-sun03/donation/backend-end/config"
	"github.com/haitao-sun03/donation/backend-end/db"
	"github.com/haitao-sun03/donation/backend-end/model"
	"github.com/haitao-sun03/donation/backend-end/service"
	"github.com/haitao-sun03/donation/backend-end/utils"
	"github.com/haitao-sun03/donation/backend-end/vo"
)

type AuthController struct{}

const NONCE_KEY string = "NONCE_"

func (au AuthController) GetNonce(ctx *gin.Context) {
	var nonceVO vo.NonceVO
	if err := ctx.ShouldBindJSON(&nonceVO); err != nil {
		common.Fail(ctx, http.StatusInternalServerError, err)
		return
	}
	nonce, err := utils.GenerateNonce()

	if err != nil {
		common.Fail(ctx, http.StatusInternalServerError, err)
		return
	}

	status := config.RedisClient.Set(context.Background(), NONCE_KEY+nonceVO.Address, nonce, 0)
	if status.Err() != nil {
		common.Fail(ctx, http.StatusInternalServerError, status.Err())
		return
	}
	common.Success(ctx, http.StatusOK, common.MessageSuccess, nonce, -1)
}

func (au AuthController) Login(ctx *gin.Context) {
	var loginVO vo.LoginVO
	if err := ctx.ShouldBindJSON(&loginVO); err != nil {
		common.Fail(ctx, http.StatusInternalServerError, err)
		return
	}

	// 获取该地址的nonce
	nonce, err := config.RedisClient.Get(context.Background(), NONCE_KEY+loginVO.Address).Result()
	if err != nil {
		common.Fail(ctx, http.StatusInternalServerError, err)
		return
	}

	// 验证签名
	valid, err := utils.VerifySignature(loginVO.Address, loginVO.Signature, nonce)
	if err != nil {
		common.Fail(ctx, http.StatusInternalServerError, err)
		return
	}

	if !valid {
		common.Fail(ctx, http.StatusInternalServerError, errors.New("signature is not valid"))
		return
	}

	//验证通过后，删除nonce，防止signature重放
	config.RedisClient.Del(context.Background(), NONCE_KEY+loginVO.Address)

	// 若用户不存在，则注册
	user := model.UserModel{
		Address: loginVO.Address,
		Role:    model.UserRole,
	}

	var claimMap = make(utils.Roles)

	userDao := db.NewUserDao()
	userCampaignIdRoleDao := db.NewUserCampaignIdRoleDao()
	_, err = userDao.GetUserByAddress(loginVO.Address)
	// 用户没有注册过，注册用户
	if err != nil {
		if err = userDao.CreateUser(user); err != nil {
			common.Fail(ctx, http.StatusInternalServerError, err)
			return
		}

	} else {
		// 用户之前已注册
		userActivityRoleModels, err := userCampaignIdRoleDao.GetUserCampaignRolesByUser(loginVO.Address)
		if err != nil {
			common.Fail(ctx, http.StatusInternalServerError, err)
			return
		}

		for _, userActivityRole := range userActivityRoleModels {
			campaignIdStr := strconv.FormatUint(uint64(userActivityRole.CampaignId), 10)
			claimMap[campaignIdStr] = userActivityRole.Role
		}
	}

	var jwts []string = make([]string, 2)
	accessJwt, err := utils.GenerateAccessJWT(loginVO.Address, claimMap)
	if err != nil {
		common.Fail(ctx, http.StatusInternalServerError, err)
		return
	}
	refreshJwt, err := utils.GenerateRefreshJWT(loginVO.Address, claimMap)
	if err != nil {
		common.Fail(ctx, http.StatusInternalServerError, err)
		return
	}
	jwts[0] = accessJwt
	jwts[1] = refreshJwt

	common.Success(ctx, http.StatusOK, common.MessageSuccess, jwts, -1)

}

// 通过refreshToken续期AccessToken
func (au AuthController) RenewJwt(ctx *gin.Context) {
	var renewVO vo.RenewVO
	if err := ctx.ShouldBindJSON(&renewVO); err != nil {
		common.Fail(ctx, http.StatusInternalServerError, err)
		return
	}

	// 验证刷新令牌有效性
	address, _, jwtType, err := utils.VerifyJWT(renewVO.RefreshJwt)
	if err != nil || jwtType != utils.JwtTypeRefresh {
		common.Fail(ctx, http.StatusUnauthorized, errors.New("invalid refresh token"))
		return
	}
	if address != renewVO.Address {
		common.Fail(ctx, http.StatusUnauthorized, errors.New("request address mismatch address in refreshToken"))
		return
	}
	authService := service.NewAuthService()
	// 查询某用户目前所有角色
	roles, shouldReturn := authService.GetRolesByAddress(address)
	if shouldReturn {
		return
	}
	accessJwt, err := utils.GenerateAccessJWT(address, roles)
	if err != nil {
		common.Fail(ctx, http.StatusUnauthorized, err)
		return
	}

	common.Success(ctx, http.StatusOK, common.MessageSuccess, accessJwt, -1)

}

// 当用户新增活动，捐赠时，更新token
func (au AuthController) RefreshJWT(ctx *gin.Context) {

	var refreshJwtVO vo.RefeshJwtVO
	if err := ctx.ShouldBindBodyWithJSON(&refreshJwtVO); err != nil {
		common.Fail(ctx, http.StatusInternalServerError, err)
		return
	}
	authService := service.NewAuthService()
	// 查询某用户目前所有角色
	roles, shouldReturn := authService.GetRolesByAddress(refreshJwtVO.Address)
	if shouldReturn {
		return
	}
	var jwts []string = make([]string, 2)
	accessJwt, err := utils.GenerateAccessJWT(refreshJwtVO.Address, roles)
	if err != nil {
		common.Fail(ctx, http.StatusInternalServerError, err)
		return
	}

	refreshJwt, err := utils.GenerateRefreshJWT(refreshJwtVO.Address, roles)
	if err != nil {
		common.Fail(ctx, http.StatusInternalServerError, err)
		return
	}
	jwts[0] = accessJwt
	jwts[1] = refreshJwt
	common.Success(ctx, http.StatusOK, common.MessageSuccess, jwts, -1)
}
