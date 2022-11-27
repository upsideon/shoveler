package middleware

import (
	"net/http"
	"shoveler/auth"

	"github.com/gin-gonic/gin"
)

func Authorization() gin.HandlerFunc {
	return func(context *gin.Context) {
		if err := auth.VerifyToken(context); err != nil {
			context.AbortWithStatusJSON(http.StatusForbidden, gin.H{"message": err.Error()})
		}
		context.Next()
	}
}
