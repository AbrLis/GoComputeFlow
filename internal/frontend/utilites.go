package frontend

import (
	"bytes"
	"github.com/gin-gonic/gin"
	"io"
	"log"
	"net/http"
	"strconv"
	"strings"
)

// sendAPIRequest вспомогательная функция запроса к API
func sendAPIRequest(path string, method string, data *bytes.Reader, header string) ([]byte, error) {
	var req *http.Request
	var err error

	if data != nil {
		req, err = http.NewRequest(method, APIPath+path, data)
	} else {
		req, err = http.NewRequest(method, APIPath+path, nil)
	}

	if err != nil {
		return nil, err
	}
	req.Header.Set("Content-Type", "application/json")
	if header != "" {
		req.Header.Set("Authorization", header)
	}
	resp, err := http.DefaultClient.Do(req)
	if err != nil {
		return nil, err
	}
	defer resp.Body.Close()
	if resp.StatusCode == 401 {
		return nil, errorUnauthorized
	}
	if resp.StatusCode != 200 {
		return nil, errorAPI
	}
	body, err := io.ReadAll(resp.Body)
	if err != nil {
		return nil, err
	}
	return body, nil
}

// parsingTimeOut вспомогательная функция, которая парсит таймауты и отдаёт числовые значения
func parsingTimeOut(timeout string) (float32, error) {
	if timeout == "" {
		return 0, errorsTimeout
	}
	timeout = strings.Split(timeout, " ")[0]
	value, err := strconv.ParseFloat(timeout, 32)
	if err != nil {
		log.Println("Error parsingTimeOut: ", err)
		return 0, errorsTimeout
	}
	return float32(value), nil
}

// checkPagination вспомогательная функция изменения параметров пагинации страниц
func checkPaginate(c *gin.Context) {
	page := c.Query("page")
	pageInt, err := strconv.Atoi(page)
	if err != nil || pageInt < 1 {
		return
	}
	indexPage = pageInt
}

// add вспомогательная функция используется в шаблоне
func add(args ...int) (int, error) {
	result := args[0]
	for _, arg := range args[1:] {
		result += arg
	}
	return result, nil
}
