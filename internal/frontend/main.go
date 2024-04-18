package frontend

import (
	"html/template"
	"log"

	"github.com/gin-gonic/gin"
)

var frontendRoute *gin.Engine

func StartFront() {
	frontendRoute = gin.Default()
	// Настройка темплейтов
	html := template.New("").Funcs(template.FuncMap{
		"add": add,
	})
	html, err := html.ParseGlob("internal/frontend/templates/*")
	if err != nil {
		log.Fatalf("Failed to parse templates: %v", err)
	}
	frontendRoute.SetHTMLTemplate(html)
	// Установка статики
	frontendRoute.Static("/static", "internal/frontend/static/")
	// Установка маршрутов
	initializeRoutes()
	// Запуск сервера
	err = frontendRoute.Run(":8080")
	log.Println("Frontend server started on port 8080")
	if err != nil {
		log.Fatalf("Failed to start server: %v", err)
	}
}
