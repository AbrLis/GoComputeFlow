package frontend

func initializeRoutes() {
	frontendRoute.GET("/", showIndexPage)
	frontendRoute.GET("/login", showLoginPage)
	frontendRoute.POST("/login", performLogin)
}
