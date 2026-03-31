package myjwt

import (
	"time"

	"github.com/golang-jwt/jwt/v4"
	"github.com/yuanji6666/gopherAI/config"
)

type Claims struct {
	ID			int64`json:"id"`
	Username	string`json:"username"`
	jwt.RegisteredClaims
}


func GenerateJwt(id int64, username string)(string, error){
	claims := Claims{
		ID: id,
		Username: username,
		RegisteredClaims: jwt.RegisteredClaims{
			ExpiresAt: jwt.NewNumericDate(time.Now().Add(time.Duration(config.GetConfig().ExpireDuration)*time.Hour)),
			Issuer: config.GetConfig().Issuer,
			Subject: config.GetConfig().Subject,
			IssuedAt: jwt.NewNumericDate(time.Now()),
		},
	}

	token := jwt.NewWithClaims(jwt.SigningMethodHS256, claims)

	return token.SignedString([]byte(config.GetConfig().Key))
}

func ParseToken(token string)(string, bool){
	claims := new(Claims)

	t, err := jwt.ParseWithClaims(token, claims, func(t *jwt.Token) (interface{}, error) {
		return []byte(config.GetConfig().Key), nil
	})

	if !t.Valid || err != nil || claims == nil {
		return "", false
	} 

	return claims.Username, true
}

