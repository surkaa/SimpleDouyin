package util

import (
	"fmt"
	"github.com/dgrijalva/jwt-go"
	"time"
)

// token 密钥
var secret = []byte("douyin")

type UserClaims struct {
	Id       int64  `json:"id"`
	Username string `json:"username"`
	jwt.StandardClaims
}

// GenerateToken 生成token
func GenerateToken(id int64, username string) (string, error) {
	// 设置过期时间 一天
	expirationTime := time.Now().Add(24 * time.Hour).Unix()
	// 创建声明
	claims := jwt.MapClaims{
		"id":       id,
		"username": username,
		"exp":      expirationTime,
	}
	// 使用密钥进行签名
	token := jwt.NewWithClaims(jwt.SigningMethodHS256, claims)
	// 生成 Token 字符串
	tokenString, err := token.SignedString(secret)
	if err != nil {
		return "", err
	}
	return tokenString, nil
}

// ParseToken 解析Token
func ParseToken(tokenString string) (UserClaims, error) {
	token, err := jwt.Parse(tokenString, func(token *jwt.Token) (interface{}, error) {
		return secret, nil
	})

	if err != nil {
		return createUserClaims(0, ""), err
	}

	// 验证 Token 签名方法是否正确
	if _, ok := token.Method.(*jwt.SigningMethodHMAC); !ok {
		return createUserClaims(0, ""), fmt.Errorf("unexpected signing method")
	}

	// 检查 Token 是否有效
	if !token.Valid {
		return createUserClaims(0, ""), fmt.Errorf("token is not valid")
	}

	// 获取声明中的数据
	claims, ok := token.Claims.(jwt.MapClaims)
	if !ok {
		return createUserClaims(0, ""), fmt.Errorf("invalid token claims")
	}

	// 提取 id 和 username
	id, ok := claims["id"].(float64)
	if !ok {
		return createUserClaims(0, ""), fmt.Errorf("invalid id claim")
	}
	username, ok := claims["username"].(string)
	if !ok {
		return createUserClaims(0, ""), fmt.Errorf("invalid username claim")
	}
	return createUserClaims(int64(id), username), nil
}

// createUserClaims 创建UserClaims
func createUserClaims(id int64, username string) UserClaims {
	return UserClaims{
		Id:       id,
		Username: username,
	}
}
