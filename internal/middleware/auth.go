package middleware

import (
	"net/http"
	"strings"

	"github.com/gin-gonic/gin"
	"github.com/hubogle/chatcode-server/internal/pkg/jwt"
)

func Auth(c *gin.Context) {
	token := c.GetHeader("Authorization")
	token = strings.TrimPrefix(token, "Bearer ")
	userClaims, err := jwt.ParseToken(token)
	if err != nil {
		c.Abort()
		c.JSON(http.StatusOK, gin.H{
			"code": -1,
			"msg":  "token invalid",
		})
		return
	}
	c.Set("user_claims", userClaims)
	c.Next()
}
