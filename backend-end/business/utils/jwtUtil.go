package utils

import (
	"crypto/rand"
	"encoding/base64"
	"errors"
	"fmt"
	"time"

	"github.com/ethereum/go-ethereum/common"
	"github.com/ethereum/go-ethereum/crypto"
	"github.com/golang-jwt/jwt/v4"
	"github.com/haitao-sun03/donation/backend-end/config"
)

var secretKey = []byte(config.Config.Jwt.SecretKey)

type Roles map[string]string

const (
	JwtTypeAccess  = "access"
	JwtTypeRefresh = "refresh"
)

func GenerateAccessJWT(address string, roles map[string]string) (string, error) {
	// payload
	claims := jwt.MapClaims{
		"address": address,                              // 用户钱包地址
		"roles":   roles,                                // 用户角色
		"exp":     time.Now().Add(time.Hour * 2).Unix(), // 过期时间：2小时
		"iat":     time.Now().Unix(),                    // 签发时间
		"type":    JwtTypeAccess,                        //令牌类型为accessToken
	}

	// 创建 accessToken
	accessToken := jwt.NewWithClaims(jwt.SigningMethodHS256, claims)

	// 签名并生成字符串
	accessTokenString, err := accessToken.SignedString(secretKey)
	if err != nil {
		return "", err
	}

	return accessTokenString, nil
}

func GenerateRefreshJWT(address string, roles map[string]string) (string, error) {
	// payload
	claims := jwt.MapClaims{
		"address": address,                                // 用户钱包地址
		"roles":   roles,                                  // 用户角色
		"exp":     time.Now().Add(time.Hour * 168).Unix(), // 过期时间：168 小时（7天）
		"iat":     time.Now().Unix(),                      // 签发时间
		"type":    JwtTypeRefresh,                         //令牌类型为refreshToken
	}

	// 创建 refreshToken
	refreshToken := jwt.NewWithClaims(jwt.SigningMethodHS256, claims)

	// 签名并生成字符串
	refreshTokenString, err := refreshToken.SignedString(secretKey)
	if err != nil {
		return "", err
	}

	return refreshTokenString, nil
}

func VerifyJWT(tokenString string) (string, Roles, string, error) {
	// 解析 token
	token, err := jwt.Parse(tokenString, func(token *jwt.Token) (any, error) {
		// 验证签名方法
		if _, ok := token.Method.(*jwt.SigningMethodHMAC); !ok {
			fmt.Println("unexpected signing method")
			return nil, fmt.Errorf("unexpected signing method: %v", token.Header["alg"])
		}
		return secretKey, nil
	})

	if err != nil {
		fmt.Println("111error: ", err)
		return "", nil, "", err
	}

	// 检查 token 是否有效
	if claims, ok := token.Claims.(jwt.MapClaims); ok && token.Valid {
		address := claims["address"].(string)
		rolesMap, ok := claims["roles"].(map[string]any)
		if !ok {
			fmt.Println("claims error: ", err)
			return "", nil, "", fmt.Errorf("invalid roles format")
		}
		roles := make(Roles)
		for key, value := range rolesMap {
			if role, ok := value.(string); ok {
				roles[key] = role
			} else {
				fmt.Println("rolesMap error: ", err)
				return "", nil, "", fmt.Errorf("invalid role value for key %s", key)
			}
		}
		return address, roles, claims["type"].(string), nil
	}
	fmt.Println("error: ", err)
	return "", nil, "", fmt.Errorf("invalid token")
}

func GenerateNonce() (string, error) {
	nonceBytes := make([]byte, 16)  // 生成 16 字节随机数
	_, err := rand.Read(nonceBytes) // 填充随机字节
	if err != nil {
		return "", err
	}
	nonceString := base64.URLEncoding.EncodeToString(nonceBytes) // 转为字符串
	return nonceString, nil
}

// 签名验证工具函数
func VerifySignature(address, signatureHex, nonce string) (bool, error) {
	// 构造预期消息
	signParam := "please login,confirm the signature that do not cost gas\n nonce: " + nonce
	message := fmt.Sprintf("\x19Ethereum Signed Message:\n%d%s", len(signParam), signParam)
	messageHash := crypto.Keccak256Hash([]byte(message))

	// 解码签名
	signature := common.Hex2Bytes(signatureHex[2:]) // 去掉0x前缀
	if len(signature) != 65 {
		return false, errors.New("signature length is not 65")
	}

	// 恢复公钥
	signature[64] -= 27 // 调整恢复标识位
	pubKey, err := crypto.SigToPub(messageHash.Bytes(), signature)
	if err != nil {
		return false, err
	}

	// 验证地址匹配
	recoveredAddr := crypto.PubkeyToAddress(*pubKey)
	return common.HexToAddress(address) == recoveredAddr, nil
}
