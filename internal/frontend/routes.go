package frontend

func initializeRoutes() {
	frontendRoute.GET("/", loggedIn(), showIndexPage)
	frontendRoute.POST("/", loggedIn(), addExpression)
	frontendRoute.GET("/login", showLoginPage)
	frontendRoute.POST("/login", performLogin)
	frontendRoute.GET("/logout", logOut)
	frontendRoute.GET("/monitoring", showMonitoring)
	frontendRoute.GET("/changeTimeouts", loggedIn(), showTimeoutsPage)
	frontendRoute.POST("/changeTimeouts", loggedIn(), performChangeTimeouts)
}
