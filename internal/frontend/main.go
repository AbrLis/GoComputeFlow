package frontend

import (
	"github.com/gin-gonic/gin"
	"log"
)

var frontendRoute *gin.Engine

func StartFront() {
	frontendRoute = gin.Default()
	frontendRoute.Static("/static", "internal/frontend/static/")
	frontendRoute.LoadHTMLGlob("internal/frontend/templates/*")
	initializeRoutes()
	err := frontendRoute.Run(":8080")
	log.Println("Frontend server started on port 8080")
	if err != nil {
		log.Fatalf("Failed to start server: %v", err)
	}
}
