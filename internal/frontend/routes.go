package frontend

func initializeRoutes() {
	frontendRoute.GET("/", loggedIn(), showIndexPage)
	frontendRoute.GET("/login", showLoginPage)
	frontendRoute.POST("/login", performLogin)
	frontendRoute.GET("/logout", logOut)
}
