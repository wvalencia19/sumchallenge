package api

import (
	"errors"
	"time"

	"github.com/dgrijalva/jwt-go"
)

type JWTClaim struct {
	Username string
	jwt.StandardClaims
}

func GenerateToken(username, secretJWTKey string, expiration time.Duration) (string, error) {
	expirationTime := time.Now().Add(expiration)
	claims := &JWTClaim{
		Username: username,
		StandardClaims: jwt.StandardClaims{
			ExpiresAt: expirationTime.Unix(),
		},
	}
	token := jwt.NewWithClaims(jwt.SigningMethodHS256, claims)
	tokenString, err := token.SignedString([]byte(secretJWTKey))
	return tokenString, err
}
func ValidToken(signedToken, secretJWTKey string) error {
	token, err := jwt.ParseWithClaims(
		signedToken,
		&JWTClaim{},
		func(token *jwt.Token) (interface{}, error) {
			return []byte(secretJWTKey), nil
		},
	)
	if err != nil {
		return err
	}
	claims, ok := token.Claims.(*JWTClaim)
	if !ok {
		return errors.New("couldn't parse claims")
	}
	if claims.ExpiresAt < time.Now().Local().Unix() {
		return errors.New("token expired")
	}
	return nil
}
