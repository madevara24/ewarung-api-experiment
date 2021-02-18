package jwt

import (
	"ewarung-api-experiment/config"
	"strings"
	"time"

	"github.com/dgrijalva/jwt-go"
)

type UserClaims struct {
	ID       int    `json:"id"`
	Username string `json:"username"`
	Role     string `json:"role"`
	jwt.StandardClaims
}

func CreateToken(userId int, username string, role string) (string, error) {
	conf, err := config.LoadConfig()
	claims := &UserClaims{
		userId,
		username,
		role,
		jwt.StandardClaims{
			ExpiresAt: time.Now().Add(time.Hour * 72).Unix(),
		},
	}
	sign := jwt.NewWithClaims(jwt.SigningMethodHS256, claims)
	secret := conf.JWTSecret
	token, err := sign.SignedString([]byte(secret))
	if err != nil {
		return "", err
	}
	return token, nil
}

func ExtractClaims(rawToken string) (jwt.MapClaims, bool) {
	token := formatBearerToke(rawToken)
	conf, err := config.LoadConfig()
	secret := conf.JWTSecret
	claims := jwt.MapClaims{}
	_, err = jwt.ParseWithClaims(token, claims, func(token *jwt.Token) (interface{}, error) {
		return []byte(secret), nil
	})

	if err != nil {
		return nil, false
	}

	return claims, true
}

func formatBearerToke(rawToken string) string {
	return strings.Replace(rawToken, "Bearer ", "", -1)
}
