package middleware

import (
	"net/http"
	"shoveler/auth"

	"github.com/gin-gonic/gin"
)

func Authorization() gin.HandlerFunc {
	return func(context *gin.Context) {
		if token, err := auth.ParseToken(context); err != nil {
			context.AbortWithStatusJSON(http.StatusForbidden, gin.H{"message": err.Error()})
		} else {
			context.Set("token", token)
			context.Next()
		}
	}
}
