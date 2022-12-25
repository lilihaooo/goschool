package util

import (
	"aschool/models"
	"fmt"
	"github.com/golang-jwt/jwt"
	"time"
)

// 定义签名秘钥
var jwtKey = []byte("qianminglilihaohao")

type Claims struct {
	UserId   int
	Username string
	jwt.StandardClaims
}

// 颁发token
func GenerateToken(user *models.User) (tokenStr string) {
	claims := Claims{
		UserId:   user.Id,
		Username: user.Username,
		StandardClaims: jwt.StandardClaims{
			ExpiresAt: time.Now().Add(time.Hour * 24 * 30 * 365).Unix(), // 过期时间
			IssuedAt:  time.Now().Unix(),                                // 签名时间
			Issuer:    "127.0.0.1",                                      // 签名颁发者
			Subject:   "user token",                                     // 签名主题
		},
	}
	token := jwt.NewWithClaims(
		jwt.SigningMethodHS256,
		claims,
	)
	tokenStr, err := token.SignedString(jwtKey)
	if err != nil {
		fmt.Println(err)
	}
	return
}

// 解析token
func ParseToken(tokenStr string) (*jwt.Token, *Claims, error) {
	claims := &Claims{}
	token, err := jwt.ParseWithClaims(tokenStr, claims, func(token *jwt.Token) (i interface{}, err error) {
		return jwtKey, nil
	})
	return token, claims, err
}
