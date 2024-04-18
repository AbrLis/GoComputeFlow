package frontend

import (
	"github.com/gin-gonic/gin"
	"html/template"
	"log"
)

var frontendRoute *gin.Engine

func StartFront() {
	frontendRoute = gin.Default()
	html := template.New("").Funcs(template.FuncMap{
		"add": add,
	})
	html, err := html.ParseGlob("internal/frontend/templates/*")
	if err != nil {
		log.Fatalf("Failed to parse templates: %v", err)
	}
	frontendRoute.SetHTMLTemplate(html)

	frontendRoute.Static("/static", "internal/frontend/static/")
	initializeRoutes()
	err = frontendRoute.Run(":8080")
	log.Println("Frontend server started on port 8080")
	if err != nil {
		log.Fatalf("Failed to start server: %v", err)
	}
}
