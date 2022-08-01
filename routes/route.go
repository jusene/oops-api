package routes

import (
	"github.com/gin-gonic/gin"
	swaggerFiles "github.com/swaggo/files"
	ginSwagger "github.com/swaggo/gin-swagger"
	"gitlab.hho-inc.com/devops/oops-api/controllers"
)


func Route(engine *gin.Engine)  {

	apis := engine.Group("/apis")
	{
		// scripts
		apis.GET("/scripts/:script", controllers.ExecShellScripts)

		// dingding webhook
		apis.POST("/webhook/dingding", controllers.WebHookSendDing)
		apis.POST("/webhook/callvoice", controllers.WebHookCallVoice)
		apis.POST("/webhook/loki/callvoice", controllers.WebHookLokiCallVoice)

		// kubernetes api
		apis.GET("/kubernetes/deployments", controllers.ListKubeDeployments)
		apis.GET("/kubernetes/pods", controllers.ListKubePods)
		apis.GET("/kubernetes/services", controllers.ListKubeServices)
		apis.GET("/kubernetes/nodes", controllers.ListKubeNodes)
		apis.GET("/kubernetes/pods/:namespace", controllers.ListKubeNameSpacePods)
		apis.GET("/kubernetes/deployments/:namespace", controllers.ListKubeNameSpaceDeployments)
		apis.GET("/kubernetes/services/:namespace", controllers.ListKubeNameSpaceServices)

		apis.POST("/kubernetes/namespace", controllers.CreateKubeNameSpace)
	}

	// health check
	engine.GET("/health/check", func(context *gin.Context) {
		context.String(200, "hello world!")
	})

	// redirct swagger
	engine.GET("/", func(context *gin.Context) {
		context.Redirect(302, "/swagger/index.html")
	})

	// swagger html
	engine.GET("/swagger/*any", ginSwagger.WrapHandler(swaggerFiles.Handler))
}
