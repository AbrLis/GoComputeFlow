package frontend

import (
	"github.com/gin-gonic/gin"
	"net/http"
)

func loggedIn() gin.HandlerFunc {
	return func(c *gin.Context) {
		loggedInInterface, _ := c.Get("is_logged_in")
		loggedIn := loggedInInterface.(bool)
		if !loggedIn {
			c.Redirect(http.StatusFound, "/login")
		}
	}
}
func setUserStatus() gin.HandlerFunc {
	return func(c *gin.Context) {
		if token, err := c.Cookie("jwt_key"); err == nil || token != "" {
			c.Set("is_logged_in", true)
		} else {
			c.Set("is_logged_in", false)
		}
	}
}
