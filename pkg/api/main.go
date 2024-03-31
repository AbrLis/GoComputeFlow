package api

import (
	"github.com/gin-gonic/gin"
	"log"
)

func StartServer() {
	router := gin.Default()

	apiRouters := router.Group(apiVersion)
	{
		apiRouters.POST(registerPath, RegisterUser)
		apiRouters.POST(LoginPath, LoginUser)

		//apiRouters.POST(addExpressionPath, ensureAuth(), AddExpressionHandler)
		//apiRouters.GET(getExpressionsPath, ensureAuth(), GetExpressionsHandler)
		//apiRouters.GET(getValuePath, ensureAuth(), GetValueHandler)
		//apiRouters.GET(getOperationsPath, ensureAuth(), GetOperationsHandler)
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
