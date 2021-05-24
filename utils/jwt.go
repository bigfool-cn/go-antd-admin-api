package utils

import (
  jwt "github.com/dgrijalva/jwt-go"
  "strings"
  "time"
  "wechat-bot-api/configs"
)

var jwtSecret = []byte(configs.Conf.App.SecretKey)

type UserInfo struct{
	UserId   int64   `json:"userId"`
	UserName string  `json:"userName"`
}

type Claims struct {
	UserInfo
	jwt.StandardClaims
}

type JwtDo struct {

}

type JwtInterface interface {
	GenerateToken(userId int64, userName string) (string, error)
	ParseToken(token string) (*Claims, error)
}

func (*JwtDo) GenerateToken(userId int64, userName string) (string, error) {
	nowTime := time.Now()
	expireTime := nowTime.Add(2 * time.Hour)
	userInfo := UserInfo{UserId: userId, UserName: userName}
	claims := Claims{
		userInfo,
      jwt.StandardClaims {
                  ExpiresAt : expireTime.Unix(),
                  Issuer : "bigfool",
		  },
	}

	tokenClaims := jwt.NewWithClaims(jwt.SigningMethodHS256, claims)
	token, err := tokenClaims.SignedString(jwtSecret)

	var build strings.Builder
	build.WriteString("Bearer ")
	build.WriteString(token)
	token = build.String()

	return token, err
}

func (*JwtDo) ParseToken(token string) (*Claims, error) {
	if strings.Contains(token,"Bearer ") {
		token = strings.Replace(token,"Bearer ","",1)
	}

	tokenClaims, err := jwt.ParseWithClaims(token, &Claims{}, func(token *jwt.Token) (interface{}, error) {
		return jwtSecret, nil
	})

	if tokenClaims != nil {
		if claims, ok := tokenClaims.Claims.(*Claims); ok && tokenClaims.Valid {
			return claims, nil
		}
	}

	return nil, err
}

var Jwt = new(JwtDo)
