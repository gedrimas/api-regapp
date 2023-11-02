package middleware

import (
	"fmt"
	"net/http"
	"api-regapp/helpers"
	"github.com/gin-gonic/gin"
)



func AuthenticateUser() gin.HandlerFunc {
	return func(c *gin.Context) {
		clientToken := c.Request.Header.Get("token")
		if clientToken == "" {
			c.JSON(http.StatusInternalServerError, gin.H{"error": fmt.Sprintf("No Authorization header provided")})
			c.Abort()
			return
		}

		claims, err := helpers.ValidateToken(clientToken)
		if err != "" {
			c.JSON(http.StatusInternalServerError, gin.H{"error": err})
			c.Abort()
			return
		}

		c.Set("username", claims.Username)
		c.Set("user_id", claims.User_id)
		c.Set("user_type", claims.User_type)
		c.Set("company_id", claims.Company_id)
		c.Next()
	}
}
