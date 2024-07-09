package helper

import (
	"blog/common/confsetting"
	"errors"
	"log"
	"time"

	"github.com/golang-jwt/jwt/v5"
)

type MyClaim struct {
	Username     string
	UserId       int
	RefreshToken int64
	jwt.RegisteredClaims
}

func GenerateToken(name string, id int) (string, error) {
	myClaim := &MyClaim{
		Username:     name,
		UserId:       id,
		RefreshToken: 300,
		RegisteredClaims: jwt.RegisteredClaims{
			ExpiresAt: jwt.NewNumericDate(time.Now().Add(time.Minute * 20)),
		},
	}

	token := jwt.NewWithClaims(jwt.SigningMethodHS256, myClaim)

	tokenString, err := token.SignedString([]byte(confsetting.JwtKey))
	if err != nil {
		log.Printf("generate token error : %v\n", err)
		return "", err
	}

	return tokenString, nil
}

func AnalyseToken(tokenString string) (*MyClaim, error) {
	token, err := jwt.ParseWithClaims(tokenString, &MyClaim{}, func(t *jwt.Token) (interface{}, error) {
		return []byte(confsetting.JwtKey), nil
	})
	if err != nil {
		log.Printf("analyse token error : %v\n", err)
		return nil, err
	}

	if claim, ok := token.Claims.(*MyClaim); ok && token.Valid {
		return claim, nil
	}
	return nil, errors.New("token is invalid")
}
