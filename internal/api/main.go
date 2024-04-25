package api

import (
	"log"

	"github.com/gin-gonic/gin"

	"GoComputeFlow/internal/api/apiConfig"
	"GoComputeFlow/internal/api/auth"
)

func StartServer(host, port string) {
	gin.SetMode(gin.ReleaseMode)
	router := gin.Default()

	apiRouters := router.Group(apiConfig.ApiVersion)
	{
		apiRouters.POST(apiConfig.RegisterPath, RegisterUser)
		apiRouters.POST(apiConfig.LoginPath, LoginUser)

		apiRouters.POST(apiConfig.AddExpressionPath, auth.EnsureAuth(), AddExpressionHandler)
		apiRouters.GET(apiConfig.GetExpressionsPath, auth.EnsureAuth(), GetExpressionsHandler)
		apiRouters.GET(apiConfig.GetValuePath, auth.EnsureAuth(), GetValueHandler)
		apiRouters.GET(apiConfig.GetOperationsPath, GetOperationsHandler)
		apiRouters.POST(apiConfig.SetOperationsPath, auth.EnsureAuth(), SetOperationsHandler)
		apiRouters.GET(apiConfig.MonitoringPath, GetMonitoringHandler)
	}

	log.Printf("Starting server on %s%s ", host, port)
	err := router.Run(host + port)
	if err != nil {
		log.Println("Error starting server: ", err)
		panic(err)
	}
}
