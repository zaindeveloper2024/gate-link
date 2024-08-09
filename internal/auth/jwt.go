package auth

import (
	"time"

	"github.com/golang-jwt/jwt/v5"
	"github.com/zaindeveloper2024/gate-link/internal/conf"
)

func GenerateToken() (string, error) {
	claims := jwt.MapClaims{
		"authorized": true,
		"exp":        time.Now().Add(time.Minute * 30).Unix(),
	}
	token := jwt.NewWithClaims(jwt.SigningMethodHS256, claims)
	return token.SignedString([]byte(conf.Conf.Authenticate.Jwtsecret))
}
