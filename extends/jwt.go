package extends

import (
	"fmt"
	"github.com/golang-jwt/jwt"
	"github.com/sirupsen/logrus"
)

type JwtUtils struct {
}

var hmacSampleSecret []byte

type AdminClaims struct {
	Id  int
	Key string
}

func (AdminClaims) Valid() error {
	return nil
}
func (JwtUtils) NewToken(id int, key string) string {

	token := jwt.NewWithClaims(jwt.SigningMethodHS256, AdminClaims{
		Id:  id,
		Key: key,
	})
	// Sign and get the complete encoded token as a string using the secret
	tokenString, err := token.SignedString(hmacSampleSecret)
	if err != nil {
		logrus.Error(err)
	}
	return tokenString
}
func (JwtUtils) ParseToken(tokenString string) *AdminClaims {
	token, err := jwt.ParseWithClaims(tokenString, &AdminClaims{}, func(token *jwt.Token) (interface{}, error) {
		if _, ok := token.Method.(*jwt.SigningMethodHMAC); !ok {
			return nil, fmt.Errorf("Unexpected signing method: %v", token.Header["alg"])
		}
		return hmacSampleSecret, nil
	})
	if err != nil {
		logrus.Error(err)
		return nil
	}
	if claims, ok := token.Claims.(*AdminClaims); ok && token.Valid {
		return claims
	}
	return nil
}
