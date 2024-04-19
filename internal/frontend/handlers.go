package frontend

import (
	"bytes"
	"encoding/json"
	"errors"
	"fmt"
	"log"
	"net/http"
	"time"

	"github.com/gin-gonic/gin"

	"GoComputeFlow/internal/api"
)

type tokenJWT struct {
	Token  string `json:"token"`
	UserID string `json:"user_id"`
}

func render(c *gin.Context, templateName string, data gin.H) {
	c.HTML(200, templateName, data)
}

// showMonitoring отображает пинги вычислителей
func showMonitoring(c *gin.Context) {
	var (
		message string
		result  map[string]string
	)
	resp, err := sendAPIRequest("/monitoring", "GET", nil, "")
	if err != nil {
		message = err.Error()
	} else {
		err = json.Unmarshal(resp, &result)
		if err != nil {
			message = err.Error()
		}
	}
	render(c, "indexMonitoring.html", gin.H{
		"monitoring":   result,
		"errorMessage": message,
		"is_logged_in": true,
		"activePage":   "Monitoring",
	})
}

// showTimeoutsPage отображает страницу таймаутов операций
func showTimeoutsPage(c *gin.Context) {
	data := getIndexTimeoutsData(c)

	render(c, "indexChangeOperations.html", gin.H{
		"is_logged_in": true,
		"errorMessage": data.Message,
		"add":          data.Add,
		"sub":          data.Sub,
		"mul":          data.Mul,
		"div":          data.Div,
		"activePage":   "Timeouts",
	})
}

// showErrorTimeoutsPage отображает страницу ошибки
func showErrorTimeoutsPage(c *gin.Context, message string) {
	render(c, "indexChangeOperations.html", gin.H{
		"is_logged_in": true,
		"errorMessage": message,
		"activePage":   "Timeouts",
	})
}

// performChangeTimeouts изменяет таймауты операций - обработка формы
func performChangeTimeouts(c *gin.Context) {
	adds := c.PostForm("add")
	sub := c.PostForm("sub")
	mul := c.PostForm("mul")
	div := c.PostForm("div")
	jwtKey, ok := c.Get("jwt_key")
	if !ok {
		log.Println("Error get jwt_key")
		c.SetCookie("message", "Error get jwt_key", timeLifeCookie, "/", "", false, true)
		c.Redirect(http.StatusSeeOther, "/changeTimeouts")
		return
	}
	header := fmt.Sprintf("Bearer %s", jwtKey)
	query := fmt.Sprintf("/set-operations?add=%s&sub=%s&mul=%s&div=%s", adds, sub, mul, div)
	resp, err := sendAPIRequest(query, "POST", nil, header)
	log.Println(string(resp))
	if err != nil {
		log.Println("Error sendAPIRequest: ", err)
		c.SetCookie("message", err.Error(), timeLifeCookie, "/", "", false, true)
	}
	c.Redirect(http.StatusSeeOther, "/changeTimeouts")
}

// showIndexPage отображает главную страницу
func showIndexPage(c *gin.Context) {
	// Запрос информации об операциях и её добавление в шаблон
	getIndexPageData(c)

	render(c, "index.html", gin.H{
		"expressions":      indexData.Expressions,
		"is_logged_in":     true,
		"count_expression": CountExpression - 1,
		"message":          indexData.Message,
		"isNext":           indexData.IsNext,
		"isPrevious":       indexData.MyPage > 1,
		"myPage":           indexData.MyPage,
		"activePage":       "Home",
	})
}

// addExpression отправляет операцию на вычисления
func addExpression(c *gin.Context) {
	expression := c.PostForm("expression")
	var message = "Empty expression"
	if expression != "" {
		jwt, _ := c.Get("jwt_key")
		header := fmt.Sprintf("Bearer %s", jwt.(string))
		resp, err := sendAPIRequest("/add-expression", "POST", bytes.NewReader([]byte(expression)), header)
		if err != nil {
			message = fmt.Sprintf("Error sendAPIRequest: %s, %s", err, string(resp))
		} else {
			message = string(resp)
		}
	}

	c.SetCookie("message", message, timeLifeCookie, "/", "", false, true)
	c.Redirect(302, "/")
}

func showLoginPage(c *gin.Context) {
	c.HTML(200, "login.html", nil)
}

func performLogin(c *gin.Context) {
	username := c.PostForm("username")
	password := c.PostForm("password")
	if len(username) <= 3 || len(password) <= 3 {
		// выдать ошибку и перейти на логин экран
		errorLogin(c, "Error:", errors.New("length of username and password must be > 3"))
		return
	}

	user := api.RegisterUserRequest{Login: username, Password: password}
	data, err := json.Marshal(user)
	if err != nil {
		// выдать ошибку и перейти на логин экран
		errorLogin(c, "No valid data", err)
		return
	}

	// обращение к API регистрации
	_, _ = sendAPIRequest("/register", "POST", bytes.NewReader(data), "")

	// обращение к API логина
	resp, err := sendAPIRequest("/login", "POST", bytes.NewReader(data), "")
	if err != nil {
		errorLogin(c, "Error", err)
		return
	}

	var token tokenJWT
	err = json.Unmarshal(resp, &token)
	if err != nil {
		errorLogin(c, "Error", err)
		return
	}

	// Логин успешно пройден, установить куку и перейти на глвную

	c.SetCookie("jwt_key", token.Token, int((time.Hour * 20).Seconds()), "/", "", false, true)
	c.SetCookie("user_id", token.UserID, int((time.Hour * 20).Seconds()), "/", "", false, true)
	c.Redirect(302, "/")
}

func errorLogin(c *gin.Context, errTitle string, err error) {

	c.HTML(http.StatusBadRequest, "login.html", gin.H{
		"ErrorTitle":   errTitle,
		"ErrorMessage": err,
	})
	c.Abort()
}

func logOut(c *gin.Context) {
	c.SetCookie("jwt_key", "", -1, "/", "", false, true)
	c.Set("is_logged_in", false)
	c.Redirect(302, "/login")
}
