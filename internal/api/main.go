package api

import (
	"log"

	"github.com/gin-gonic/gin"

	"GoComputeFlow/internal/api/auth"
)

func StartServer(host, port string) {
	gin.SetMode(gin.ReleaseMode)
	router := gin.Default()

	apiRouters := router.Group(apiVersion)
	{
		apiRouters.POST(registerPath, RegisterUser)
		apiRouters.POST(LoginPath, LoginUser)

		apiRouters.POST(addExpressionPath, auth.EnsureAuth(), AddExpressionHandler)
		apiRouters.GET(getExpressionsPath, auth.EnsureAuth(), GetExpressionsHandler)
		apiRouters.GET(getValuePath, auth.EnsureAuth(), GetValueHandler)
		apiRouters.GET(getOperationsPath, GetOperationsHandler)
		apiRouters.POST(setOperationsPath, auth.EnsureAuth(), SetOperationsHandler)
		apiRouters.GET(monitoringPath, GetMonitoringHandler)
	}

	log.Printf("Starting server on %s%s ", host, port)
	err := router.Run(host + port)
	if err != nil {
		log.Println("Error starting server: ", err)
		panic(err)
	}
}
