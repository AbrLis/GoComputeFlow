package frontend

import (
	"github.com/gin-gonic/gin"
	"net/http"
)

func loggedIn() gin.HandlerFunc {
	return func(c *gin.Context) {
		setLoggedIn := func(key, value string) {
			if val, err := c.Cookie(key); err == nil || val != "" {
				c.Set(value, val)
			} else {
				c.Set("is_logged_in", false)
				c.Redirect(http.StatusFound, "/login")
				c.Abort()
			}
		}
		setLoggedIn("jwt_key", "is_logged_in")
		setLoggedIn("user_id", "user_id")
		c.Set("is_logged_in", true)
	}
}
