package frontend

import (
	"github.com/gin-gonic/gin"
	"net/http"
)

func loggedIn() gin.HandlerFunc {
	return func(c *gin.Context) {
		setLoggedIn := func(key string) {
			if val, err := c.Cookie(key); err == nil || val != "" {
				c.Set(key, val)
			} else {
				c.Set("is_logged_in", false)
				c.Redirect(http.StatusFound, "/login")
				c.Abort()
			}
		}
		setLoggedIn("jwt_key")
		setLoggedIn("user_id")
		c.Set("is_logged_in", true)
	}
}
