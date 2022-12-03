package middleware

import (
	"errors"
	"log"
	"net/http"
	"shoveler/auth"

	"github.com/gin-gonic/gin"
	"github.com/golang-jwt/jwt"
)

func Authorization() gin.HandlerFunc {
	return func(context *gin.Context) {
		if token, err := auth.ParseToken(context); err != nil {
			context.AbortWithStatusJSON(http.StatusForbidden, gin.H{"message": err.Error()})
		} else {
			claims, ok := token.Claims.(jwt.MapClaims)
			if !ok {
				log.Print(errors.New("Unable to access token claims."))
				return
			}

			accountId := claims["sub"].(string)

			context.Set("accountId", accountId)
			context.Next()
		}
	}
}
