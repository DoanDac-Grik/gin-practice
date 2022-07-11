package services

import (
	"fmt"
	"gin-practice/models"
	"gin-practice/utils"

	"time"

	"github.com/dgrijalva/jwt-go"
)

var secretKey = utils.GetEnvVariable("SECRET_KEY")

type authCustomClaims struct {
	User models.User
	jwt.StandardClaims
}

func GenerateToken(user models.User) string {
	claims := &authCustomClaims{
		user,
		jwt.StandardClaims{
			ExpiresAt: time.Now().Add(time.Hour * 48).Unix(),
			IssuedAt:  time.Now().Unix(),
		},
	}
	token := jwt.NewWithClaims(jwt.SigningMethodHS256, claims)

	//Encoded string

	t, err := token.SignedString([]byte(secretKey))
	if err != nil {
		panic(err)
	}
	return t
}

func ValidateToken(encodedToken string) (*jwt.Token, error) {
	return jwt.Parse(encodedToken, func(token *jwt.Token) (interface{}, error) {
		if _, isValid := token.Method.(*jwt.SigningMethodHMAC); !isValid {
			return nil, fmt.Errorf("Invalid Token", token.Header["alg"])
		}
		return []byte(secretKey), nil
	})
}

func DecodeingToken(tokenString string) jwt.MapClaims {
	claims := jwt.MapClaims{}
	jwt.ParseWithClaims(tokenString, claims, func(token *jwt.Token) (interface{}, error) {
		return []byte(secretKey), nil
	})
	return claims
}
