package middlewares

import (
	"log"
	"net/http"
	"strings"

	"multiverse/core/shared"

	"github.com/gin-gonic/gin"
)

func JwtAuthorizationMiddleware() gin.HandlerFunc {
	return func(c *gin.Context) {
		authHeader := c.Request.Header.Get("Authorization")
		if authHeader == "" {
			c.AbortWithStatus(http.StatusUnauthorized)
			return
		}
		tokenString := strings.TrimSpace(strings.SplitN(authHeader, "Bearer", 2)[1])
		claims, err := shared.ValidateAndGetClaims(tokenString)
		if err == nil {
			log.Println(claims["exp"])
			c.Set("user_email", claims["user_email"])
			c.Next()
		} else {
			c.AbortWithStatus(http.StatusUnauthorized)
		}
	}
}
