package middlewares

import (
	"net/http"

	"github.com/Soyaib10/eba-event-booking-api/utils"
	"github.com/gin-gonic/gin"
)

func Authenticate(c *gin.Context) {
	// Authorization
	token := c.Request.Header.Get("Authorization")
	if token == "" {
		c.AbortWithStatusJSON(http.StatusUnauthorized, gin.H{"message": "Not authorized"})
		return
	}
	userId, err := utils.VerifyToken(token)
	if err != nil {
		c.AbortWithStatusJSON(http.StatusUnauthorized, gin.H{"message": "Not authorized"})
		return
	}

	c.Set("userId", userId)
	c.Next()
}
