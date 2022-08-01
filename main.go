package main

import (
	_ "gitlab.hho-inc.com/devops/oops-api/docs"
	"gitlab.hho-inc.com/devops/oops-api/services"
)

// @title Oops API
// @version 1.0
// @description A server for Oops api

// @schemes http https
// @securityDefinitions.apiKey ApiKeyAuth
// @in header
// @name Authorization

// @BasePath /apis
func main() {
	app := services.App()
	app.Run(":9001")
}
