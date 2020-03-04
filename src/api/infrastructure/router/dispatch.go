package router

import (
	"api/interfaces"
	"api/interfaces/auth"
	"api/interfaces/user"
	"api/middleware"
	"api/usecases"

	"github.com/gin-gonic/gin"
)

func Dispatch(r *gin.Engine, sqlHandler interfaces.SQLHandler, authHandler interfaces.AuthHandler, logger usecases.Logger) {
	apiV1 := r.Group("/api/v1")
	userRouter(apiV1, sqlHandler, logger)
	authRouter(apiV1, authHandler, logger)
}

func userRouter(r *gin.RouterGroup, sqlHandler interfaces.SQLHandler, logger usecases.Logger) {
	userController := user.NewUserController(sqlHandler, logger)
	user := r.Group("/user")
	user.GET("/", func(c *gin.Context) { userController.Index(c) })
	user.GET("/:id", func(c *gin.Context) { userController.Show(c) })
}

func authRouter(r *gin.RouterGroup, authHandler interfaces.AuthHandler, logger usecases.Logger) {
	authController := auth.NewAuthController(logger)
	auth := r.Group("/auth")

	auth.Use(middleware.Authorize(authHandler))
	auth.GET("/", func(c *gin.Context) { authController.Index(c) })
}
