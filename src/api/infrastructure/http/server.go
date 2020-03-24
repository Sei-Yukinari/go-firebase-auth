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
	//Create SQL Handler
	sqlHandler, err := db.NewSQLHandler()
	if err != nil {
		fmt.Printf("Error : [%s]", err)
	}

	//Database migration
	migration.Exec(sqlHandler)
	seed.Exec(sqlHandler)

	//Create Auth Handler
	authHandler, err := auth.NewAuthHandler()
	if err != nil {
		fmt.Printf("Error : [%s]", err)
	}
	logger := log.NewLogger()
	if err := endless.ListenAndServe(":"+env.Load().APIPort, router.Handler(sqlHandler, authHandler, logger)); err != nil {
		fmt.Printf("Error : [%s]", err)
	}
}
