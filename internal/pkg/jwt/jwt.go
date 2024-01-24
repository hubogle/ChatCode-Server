package jwt

import (
	"github.com/golang-jwt/jwt/v4"
)

type UserClaims struct {
	UID     uint64 `json:"uid"`
	Account string `json:"account"`
	jwt.RegisteredClaims
}

var key = []byte("chat-code")

func GenerateToken(uid uint64, account string) (string, error) {
	UserClaim := &UserClaims{
		UID:              uid,
		Account:          account,
		RegisteredClaims: jwt.RegisteredClaims{},
	}
	token := jwt.NewWithClaims(jwt.SigningMethodHS256, UserClaim)
	tokenString, err := token.SignedString(key)
	if err != nil {
		return "", err
	}
	return tokenString, nil
}

func ParseToken(tokenString string) (*UserClaims, error) {
	token, err := jwt.ParseWithClaims(tokenString, &UserClaims{}, func(token *jwt.Token) (interface{}, error) {
		return key, nil
	})
	if err != nil {
		return nil, err
	}
	claims := token.Claims.(*UserClaims)
	return claims, nil
}
