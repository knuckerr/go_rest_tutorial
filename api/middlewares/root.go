package middlewares

import (
	"github.com/gin-gonic/gin"
	"github.com/knuckerr/go_rest/api/auth"
	"net/http"
	"strings"
)

func AuthenticationRequired() gin.HandlerFunc {
	return func(c *gin.Context) {
		reqToken := c.Request.Header.Get("Authorization")
		splitToken := strings.Split(reqToken, "Bearer")
		if len(splitToken) == 1 {
			c.JSON(http.StatusUnauthorized, gin.H{"error": "user needs to be signed in to access this service"})
			c.Abort()
			return
		}
		reqToken = strings.TrimSpace(splitToken[1])
		var claims = &auth.Claims{}
		err := auth.Vaildtoken(reqToken, claims)
		if err != nil {
			c.JSON(http.StatusUnauthorized, gin.H{"error": "user needs to be signed in to access this service"})
			c.Abort()
			return
		}
		c.Next()
	}
}
