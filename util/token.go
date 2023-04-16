package util

import (
	"time"

	"github.com/dgrijalva/jwt-go"
)

// JWT 相关的配置
const (
	// 签名密钥，保密
	jwtSecret = "mysecret"
	// 过期时间为 1 小时
	jwtExpireDuration = time.Hour * 1
)

// CustomClaims 自定义的 token 中的 payload
type CustomClaims struct {
	UserID int64 `json:"user_id"`
	jwt.StandardClaims
}

// GenerateJWT 生成 JWT Token
func GenerateJWT(userID int64) (string, error) {
	// 构造自定义的 token 中的 payload
	claims := CustomClaims{
		UserID: userID,
		StandardClaims: jwt.StandardClaims{
			ExpiresAt: time.Now().Add(jwtExpireDuration).Unix(), // token 过期时间
			Issuer:    "myapp",
		},
	}

	// 生成 token
	token := jwt.NewWithClaims(jwt.SigningMethodHS256, claims)

	// 使用密钥签名 token
	tokenString, err := token.SignedString([]byte(jwtSecret))
	if err != nil {
		return "", err
	}

	return tokenString, nil
}

// ParseToken 解析 JWT Token
func ParseToken(tokenString string) (*CustomClaims, error) {
	// 解析 token
	token, err := jwt.ParseWithClaims(tokenString, &CustomClaims{}, func(token *jwt.Token) (interface{}, error) {
		return []byte(jwtSecret), nil
	})
	if err != nil {
		return nil, err
	}

	// 校验 token 是否有效
	if claims, ok := token.Claims.(*CustomClaims); ok && token.Valid {
		return claims, nil
	}

	return nil, jwt.ErrInvalidKey
}
