package util

import (
	"strings"

	"github.com/gin-gonic/gin"
)

func GetToken(c *gin.Context) string {
	authHeader := c.GetHeader("Authorization")
	return strings.Replace(authHeader, "Bearer ", "", 1)
}
