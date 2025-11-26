package main

import (
	"go-basesource/src/adapters"
	"go-basesource/src/persentation/controller"
)

// @title Go Clean Architecture API
// @version 1.0
// @description A RESTful API built with Go, Gin, GORM, and PostgreSQL using Clean Architecture principles
// @termsOfService http://swagger.io/terms/

// @contact.name API Support
// @contact.email support@example.com

// @license.name MIT
// @license.url https://opensource.org/licenses/MIT

// @host localhost:8080
// @BasePath /
// @schemes http https
func main() {

	// Initialize storage
	adapters.StorageInit()

	// Initialize controllers
	controller.Initialize()

	// Start HTTP server
	adapters.StartHttpServer()
}
