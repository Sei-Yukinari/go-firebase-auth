package middleware

import (
	"context"
	"fmt"
	"log"
	netHttp "net/http"

	"api/infrastructure/util"
	"api/interfaces"

	"github.com/gin-gonic/gin"
)

func Authorize(authHandler interfaces.AuthHandler) gin.HandlerFunc {
	return func(c *gin.Context) {
		idToken := util.GetToken(c)
		token, err := authHandler.VerifyIDToken(context.Background(), idToken)
		if err != nil {
			fmt.Printf("error verifying ID token: %v\n", err)
			c.AbortWithStatusJSON(netHttp.StatusUnauthorized, gin.H{"Message": "Unauthorized..."})
		}
		log.Printf("Verified ID token: %v\n", token)
		c.Next()
	}
}
