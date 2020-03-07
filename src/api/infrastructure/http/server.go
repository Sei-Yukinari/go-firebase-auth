package http

import (
	"fmt"

	"github.com/Sei-Yukinari/go-firebase-auth/src/api/infrastructure/auth"
	"github.com/Sei-Yukinari/go-firebase-auth/src/api/infrastructure/db"
	"github.com/Sei-Yukinari/go-firebase-auth/src/api/infrastructure/log"
	"github.com/Sei-Yukinari/go-firebase-auth/src/api/infrastructure/router"

	"github.com/fvbock/endless"

	"github.com/Sei-Yukinari/go-firebase-auth/src/api/config"
	"github.com/Sei-Yukinari/go-firebase-auth/src/api/infrastructure/env"

	_ "github.com/go-sql-driver/mysql"
)

func StartHttpServer() {
	c := config.Load()
	env.Load()
	sqlHandler, err := db.NewSQLHandler()
	if err != nil {
		fmt.Printf("Error : [%s]", err)
	}
	authHandler, err := auth.NewAuthHandler()
	if err != nil {
		fmt.Printf("Error : [%s]", err)
	}
	logger := log.NewLogger()
	if err := endless.ListenAndServe(":"+c.Server.Port, router.Handler(sqlHandler, authHandler, logger)); err != nil {
		fmt.Printf("Error : [%s]", err)
	}
}
