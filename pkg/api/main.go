package api

import (
	"log"

	"github.com/gin-gonic/gin"

	"GoComputeFlow/pkg/api/auth"
)

func StartServer() {
	router := gin.Default()

	apiRouters := router.Group(apiVersion)
	{
		apiRouters.POST(registerPath, RegisterUser)
		apiRouters.POST(LoginPath, LoginUser)

		apiRouters.POST(addExpressionPath, auth.EnsureAuth(), AddExpressionHandler)
		//apiRouters.GET(getExpressionsPath, auth.EnsureAuth(), GetExpressionsHandler)
		//apiRouters.GET(getValuePath, auth.EnsureAuth(), GetValueHandler)
		//apiRouters.GET(getOperationsPath, auth.EnsureAuth(), GetOperationsHandler)
		//apiRouters.GET(monitoring, GetMonitoringHandler)
	}

	log.Printf("Starting server on %s:%s ", PortHost, HostPath)
	go func() {
		err := router.Run(HostPath + PortHost)
		if err != nil {
			log.Println("Error starting server: ", err)
			panic(err)
		}
	}()
}
