package frontend

import "github.com/gin-gonic/gin"

func showIndexPage(c *gin.Context) {
	c.HTML(200, "index2.html", nil)
}
