package auth

import (
	"errors"
	"fmt"
	"os"
	"strings"
	"time"

	"github.com/gin-gonic/gin"
	"github.com/golang-jwt/jwt"
)

func GenerateToken(id string) (string, error) {
	token := jwt.New(jwt.SigningMethodHS256)
	claims := token.Claims.(jwt.MapClaims)

	claims["id"] = id
	claims["exp"] = time.Now().Add(time.Minute * 30).Unix()

	tokenString, err := token.SignedString([]byte(os.Getenv("JWT_SIGNING_SECRET")))
	if err != nil {
		return "", err
	}
	return tokenString, nil
}

func VerifyToken(context *gin.Context) error {
	_, err := ParseToken(context)
	if err != nil {
		return err
	}
	return nil
}

func RetrieveToken(context *gin.Context) (string, error) {
	authorizationHeader := context.GetHeader("Authorization")

	if authorizationHeader == "" {
		return "", errors.New("Must provide Authorization header in request.")
	}

	fields := strings.Fields(authorizationHeader)

	if len(fields) < 2 {
		return "", errors.New("Invalid Authorization header format.")
	}

	authorizationType := strings.ToLower(fields[0])

	if strings.ToLower(authorizationType) != "bearer" {
		return "", errors.New("Unsupported authorization type.")
	}

	return fields[1], nil
}

func ParseToken(context *gin.Context) (*jwt.Token, error) {
	tokenString, err := RetrieveToken(context)

	if err != nil {
		return nil, err
	}

	token, err := jwt.Parse(tokenString, func(token *jwt.Token) (interface{}, error) {
		if _, ok := token.Method.(*jwt.SigningMethodHMAC); !ok {
			return nil, fmt.Errorf("Unexpected signing method: %v", token.Header["alg"])
		}
		return []byte(os.Getenv("JWT_SIGNING_SECRET")), nil
	})

	if err != nil {
		if err == jwt.ErrSignatureInvalid {
			return nil, errors.New("Invalid token signature.")
		}
		return nil, err
	}

	if !token.Valid {
		return nil, errors.New("Invalid token.")
	}

	return token, err
}
