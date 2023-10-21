package token

import (
	"tech-challenge-auth/internal/config"
	"time"

	jwt "github.com/dgrijalva/jwt-go"
)

var (
	cfg = &config.Cfg
)

func GenerateToken(userId string) (string, error) {
	permissions := jwt.MapClaims{}

	if userId == "" {
		permissions["guest"] = true
	}
	permissions["authorized"] = true
	permissions["exp"] = time.Now().Add(time.Hour * 6).Unix()
	permissions["userId"] = userId
	token := jwt.NewWithClaims(jwt.SigningMethodHS256, permissions)

	return token.SignedString([]byte(cfg.Token.Key))
}
