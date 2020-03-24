package main

import (
	"api/infrastructure/http"
)

// @title go firebase auth
// @version 1.0
// @description This is a sample firebase auth server.

// @license.name Apache 2.0
// @license.url http://www.apache.org/licenses/LICENSE-2.0.html

// @host localhost:3302
// @BasePath /api/v1

// @securityDefinitions.apikey ApiKeyAuth
// @in header
// @name Authorization
func main() {
	http.StartHttpServer()
}
