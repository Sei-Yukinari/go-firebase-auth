package http

import (
	"api/infrastructure/auth"
	"api/infrastructure/db"
	"api/infrastructure/log"
	"api/infrastructure/router"
	"fmt"

	"github.com/fvbock/endless"

	"api/config"
	"api/infrastructure/env"

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
