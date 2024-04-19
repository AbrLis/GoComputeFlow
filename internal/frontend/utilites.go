package frontend

import (
	"bytes"
	"encoding/json"
	"errors"
	"fmt"
	"io"
	"log"
	"net/http"
	"strconv"
	"strings"

	"github.com/gin-gonic/gin"
)

var indexData indexPageData // Данные для шаблона index.html

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
	indexData.MyPage = pageInt
}

// add вспомогательная функция используется в шаблоне
func add(args ...int) (int, error) {
	result := args[0]
	for _, arg := range args[1:] {
		result += arg
	}
	return result, nil
}

// getIndexPageData вспомогательная функция для получения данных для шаблона
func getIndexPageData(c *gin.Context) {
	checkPaginate(c)
	offset := 0
	if indexData.MyPage > 1 {
		offset = (indexData.MyPage - 1) * (CountExpression - 1)
	}
	jwt, _ := c.Get("jwt_key")
	header := fmt.Sprintf("Bearer %s", jwt.(string))
	url := fmt.Sprintf("/get-expressions?limit=%d&page=%d&offset=%d", CountExpression, indexData.MyPage, offset)
	data, err := sendAPIRequest(url, "GET", nil, header)
	if errors.Is(err, errorUnauthorized) {
		c.Redirect(http.StatusSeeOther, "/login")
		return
	}
	if err != nil {
		log.Println("Error sendAPIRequest: ", err)
		indexData.Message = err.Error()
	} else {
		err = json.Unmarshal(data, &indexData.Expressions)
		if err != nil {
			indexData.Message = err.Error()
		} else {
			indexData.Message, _ = c.Cookie("message")
		}
	}

	indexData.IsNext = false
	if len(indexData.Expressions) == CountExpression {
		indexData.Expressions = indexData.Expressions[:CountExpression-1]
		indexData.IsNext = true
	}
}

// getIndexTimeoutsData вспомогательная функция для получения данных для шаблона
func getIndexTimeoutsData(c *gin.Context) indexPageMonitoring {
	var result indexPageMonitoring
	operations, err := sendAPIRequest("/get-operations", "GET", nil, "")
	if err != nil {
		showErrorTimeoutsPage(c, err.Error())
		return result
	}
	var response map[string]string
	err = json.Unmarshal(operations, &response)
	if err != nil {
		showErrorTimeoutsPage(c, err.Error())
		return result
	}

	var err1, err2, err3, err4 error
	result.Add, err1 = parsingTimeOut(response["+"])
	result.Sub, err2 = parsingTimeOut(response["-"])
	result.Mul, err3 = parsingTimeOut(response["*"])
	result.Div, err4 = parsingTimeOut(response["/"])

	if err1 != nil || err2 != nil || err3 != nil || err4 != nil {
		showErrorTimeoutsPage(c, errorParsingTimeout.Error())
		return result
	}

	if message, _ := c.Cookie("message"); message != "" {
		result.Message = message
	}

	return result
}

func init() {
	indexData = indexPageData{
		MyPage: 1,
	}
}
