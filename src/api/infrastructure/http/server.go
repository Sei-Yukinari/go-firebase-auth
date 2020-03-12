package http

import (
	"api/infrastructure/auth"
	"api/infrastructure/db/seed"
	"api/infrastructure/env"
	"api/infrastructure/log"
	"api/infrastructure/router"
	"fmt"
	"github.com/fvbock/endless"

	"api/infrastructure/db"
	"api/infrastructure/db/migration"
	_ "github.com/go-sql-driver/mysql"
)

func StartHttpServer() {
	sqlHandler, err := db.NewSQLHandler()
	if err != nil {
		fmt.Printf("Error : [%s]", err)
	}

	migration.Exec(sqlHandler)
	seed.Exec(sqlHandler)

	authHandler, err := auth.NewAuthHandler()
	if err != nil {
		fmt.Printf("Error : [%s]", err)
	}
	logger := log.NewLogger()
	if err := endless.ListenAndServe(":"+env.Load().APIPort, router.Handler(sqlHandler, authHandler, logger)); err != nil {
		fmt.Printf("Error : [%s]", err)
	}
}
