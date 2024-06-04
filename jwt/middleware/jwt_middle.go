package middleware

import (
	"net/http"
	"strings"

	"github.com/gin-gonic/gin"
	"github.com/x14n/goExperimental/jwt/config"
)

func JWT() gin.HandlerFunc {
	return func(ctx *gin.Context) {
		authHeader := ctx.GetHeader("Authorization")
		if authHeader == "" {
			ctx.JSON(http.StatusUnauthorized, gin.H{
				"message": "Authorization header",
			})

			ctx.Abort()
			return
		}

		parts := strings.SplitN(authHeader, " ", 2)
		if len(parts) != 2 || parts[0] != "Bearer" {
			ctx.JSON(http.StatusUnauthorized, gin.H{"message": "Authorization"})

			ctx.Abort()
			return
		}

		tokenString := parts[1]

		claims, err := config.ValidToken(tokenString)
		if err != nil {
			ctx.JSON(http.StatusUnauthorized, gin.H{"message": "Invalid token"})

			ctx.Abort()
			return
		}

		ctx.Set("username", claims.Username)
		ctx.Set("uuid", claims.UUID)
		ctx.Next()
	}
}
