package utils

import (
	"fmt"
	"time"

	"github.com/golang-jwt/jwt/v4"
	"github.com/haitao-sun03/donation/backend-end/config"
)

var secretKey = []byte(config.Config.Jwt.SecretKey)

type Roles map[string]string

func GenerateJWT(address string, roles map[string]string) (string, error) {
	// payload
	claims := jwt.MapClaims{
		"address": address,                               // 用户钱包地址
		"roles":   roles,                                 // 用户角色
		"exp":     time.Now().Add(time.Hour * 24).Unix(), // 过期时间：24小时
		"iat":     time.Now().Unix(),                     // 签发时间
	}

	// 创建 token
	token := jwt.NewWithClaims(jwt.SigningMethodHS256, claims)

	// 签名并生成字符串
	tokenString, err := token.SignedString(secretKey)
	if err != nil {
		return "", err
	}
	return tokenString, nil
}

func VerifyJWT(tokenString string) (string, Roles, error) {
	// 解析 token
	token, err := jwt.Parse(tokenString, func(token *jwt.Token) (any, error) {
		// 验证签名方法
		if _, ok := token.Method.(*jwt.SigningMethodHMAC); !ok {
			return nil, fmt.Errorf("unexpected signing method: %v", token.Header["alg"])
		}
		return secretKey, nil
	})

	if err != nil {
		return "", nil, err
	}

	// 检查 token 是否有效
	if claims, ok := token.Claims.(jwt.MapClaims); ok && token.Valid {
		address := claims["address"].(string)
		rolesMap, ok := claims["roles"].(map[string]any)
		if !ok {
			return "", nil, fmt.Errorf("invalid roles format")
		}
		roles := make(Roles)
		for key, value := range rolesMap {
			if role, ok := value.(string); ok {
				roles[key] = role
			} else {
				return "", nil, fmt.Errorf("invalid role value for key %s", key)
			}
		}
		return address, roles, nil
	}
	return "", nil, fmt.Errorf("invalid token")
}
