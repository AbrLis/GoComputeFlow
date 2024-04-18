package api

import (
	"strconv"

	"github.com/gin-gonic/gin"
)

// getQueryParameter вспомогательная утилита парсинга int параметров
func getQueryParameter(c *gin.Context, parameter string, defaultValue int) (int, error) {
	value := c.Query(parameter)
	if value != "" {
		intValue, err := strconv.Atoi(value)
		if err != nil {
			return defaultValue, err
		}
		return intValue, nil
	}
	return defaultValue, nil
}
