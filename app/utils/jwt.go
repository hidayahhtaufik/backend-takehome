package utils

import (
	"time"

	"github.com/golang-jwt/jwt/v5"
)

func SignJWT(secret string, uid uint64) (string, error) {
	claims := jwt.MapClaims{"uid": uid, "iat": time.Now().Unix(), "exp": time.Now().Add(24 * time.Hour).Unix()}
	return jwt.NewWithClaims(jwt.SigningMethodHS256, claims).SignedString([]byte(secret))
}
