package router

import (
	"github.com/Sei-Yukinari/go-firebase-auth/src/api/interfaces"
	"github.com/Sei-Yukinari/go-firebase-auth/src/api/usecases"

	"github.com/gin-gonic/gin"
)

func Handler(sqlHandler interfaces.SQLHandler, authHandler interfaces.AuthHandler, logger usecases.Logger) *gin.Engine {
	r := gin.Default()
	r.Use(gin.Logger())
	Dispatch(r, sqlHandler, authHandler, logger)

	return r
}
