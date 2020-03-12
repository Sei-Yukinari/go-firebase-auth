package router

import (
	"api/interfaces"
	"api/usecases"

	"github.com/gin-gonic/gin"
)

func Handler(sqlHandler interfaces.SQLHandler, authHandler interfaces.AuthHandler, logger usecases.Logger) *gin.Engine {
	r := gin.Default()
	r.Use(gin.Logger())
	Dispatch(r, sqlHandler, authHandler, logger)

	return r
}
