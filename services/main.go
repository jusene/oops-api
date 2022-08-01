package services

import (
	"fmt"
	"github.com/gin-gonic/gin"
	"gitlab.hho-inc.com/devops/oops-api/config"
	_ "gitlab.hho-inc.com/devops/oops-api/docs"
	"gitlab.hho-inc.com/devops/oops-api/middlewares"
	"gitlab.hho-inc.com/devops/oops-api/routes"
)

func App() *gin.Engine {
	// init logger
	if err := middlewares.InitLogger(config.LogConf); err != nil {
		fmt.Printf("init logger failed, err: %v\n", err)
	}
	gin.SetMode(gin.DebugMode)
	engine := gin.Default()
	engine.Use(middlewares.GinLogger(), middlewares.GinRecovery(true))
	routes.Route(engine)
	return engine
}
