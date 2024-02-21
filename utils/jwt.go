package utils

import (
	"errors"
	"os"
	"time"

	"github.com/ChenMiaoQiu/go-cloud-disk/model"
	"github.com/golang-jwt/jwt/v5"
)

type MyClaims struct {
	UserId   string `json:"user_id"`
	UserName string `json:"user_name"`
	Status   string `json:"status"`
	jwt.RegisteredClaims
}

// GenToken generate jwt token
func GenToken(issuer string, expireHour int, user *model.User) (string, error) {
	mySigningKey := []byte(os.Getenv("JWT_KEY"))
	claims := MyClaims{
		UserId:   user.Uuid,
		UserName: user.UserName,
		Status:   user.Status,
		RegisteredClaims: jwt.RegisteredClaims{
			Issuer:    issuer,
			ExpiresAt: jwt.NewNumericDate(time.Now().Add(time.Duration(expireHour) * time.Hour)),
			IssuedAt:  jwt.NewNumericDate(time.Now()),
		},
	}
	token := jwt.NewWithClaims(jwt.SigningMethodHS256, claims)
	return token.SignedString(mySigningKey)
}

// ParseToken parse jwt token
func ParseToken(tokenString string) (*MyClaims, error) {
	mySigningKey := []byte(os.Getenv("JWT_KEY"))
	token, err := jwt.ParseWithClaims(tokenString, &MyClaims{}, func(t *jwt.Token) (interface{}, error) {
		return mySigningKey, nil
	})

	if err != nil {
		return nil, err
	}

	if claims, ok := token.Claims.(*MyClaims); ok && token.Valid {
		return claims, nil
	}

	return nil, errors.New("token err")
}
