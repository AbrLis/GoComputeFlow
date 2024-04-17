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
	"GoComputeFlow/internal/models"
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
	})
}

// showTimeoutsPage отображает страницу таймаутов операций
func showTimeoutsPage(c *gin.Context) {
	var (
		errorMessage = ""
		response     map[string]string
	)
	operations, err := sendAPIRequest("/get-operations", "GET", nil, "")
	if err != nil {
		showErrorTimeoutsPage(c, err.Error())
		return
	}
	err = json.Unmarshal(operations, &response)
	if err != nil {
		showErrorTimeoutsPage(c, err.Error())
		return
	}

	addTimout, err1 := parsingTimeOut(response["+"])
	subTimout, err2 := parsingTimeOut(response["-"])
	mulTimout, err3 := parsingTimeOut(response["*"])
	divTimout, err4 := parsingTimeOut(response["/"])

	if err1 != nil || err2 != nil || err3 != nil || err4 != nil {
		showErrorTimeoutsPage(c, err.Error())
		return
	}

	if message, _ := c.Cookie("message"); message != "" {
		errorMessage += message
	}
	render(c, "indexChangeOperations.html", gin.H{
		"is_logged_in": true,
		"errorMessage": errorMessage,
		"add":          addTimout,
		"sub":          subTimout,
		"mul":          mulTimout,
		"div":          divTimout,
	})
}

// showErrorTimeoutsPage отображает страницу ошибки
func showErrorTimeoutsPage(c *gin.Context, message string) {
	render(c, "indexChangeOperations.html", gin.H{
		"is_logged_in": true,
		"errorMessage": message,
	})
}

// performChangeTimeouts изменяет таймауты операций - обработка формы
func performChangeTimeouts(c *gin.Context) {
	add := c.PostForm("add")
	sub := c.PostForm("sub")
	mul := c.PostForm("mul")
	div := c.PostForm("div")
	query := fmt.Sprintf("/set-operations?add=%s&sub=%s&mul=%s&div=%s", add, sub, mul, div)
	jwtKey, ok := c.Get("jwt_key")
	if !ok {
		log.Println("Error get jwt_key")
		c.SetCookie("message", "Error get jwt_key", timeLifeCookie, "/", "", false, true)
		c.Redirect(http.StatusFound, "/changeTimeouts")
		c.Abort()
	}
	header := fmt.Sprintf("Bearer %s", jwtKey)
	resp, err := sendAPIRequest(query, "POST", nil, header)
	log.Println(string(resp))
	if err != nil {
		log.Println("Error sendAPIRequest: ", err)
		c.SetCookie("message", err.Error(), timeLifeCookie, "/", "", false, true)
	}
	c.Redirect(http.StatusFound, "/changeTimeouts")
}

// showIndexPage отображает главную страницу
func showIndexPage(c *gin.Context) {
	// Запрос информации об операциях и её добавление в шаблон
	var (
		message    string
		dataStruct []models.Expression
	)

	jwt, _ := c.Get("jwt_key")
	header := fmt.Sprintf("Bearer %s", jwt.(string))
	data, err := sendAPIRequest(fmt.Sprintf("/get-expressions?limit=%s&page=1", CountExpression), "GET", nil, header)
	if errors.Is(err, errorUnauthorized) {
		c.Redirect(http.StatusFound, "/login")
		return
	}
	if err != nil {
		log.Println("Error sendAPIRequest: ", err)
		message = err.Error()
	} else {
		err = json.Unmarshal(data, &dataStruct)
		if err != nil {
			message = err.Error()
		} else {
			message, _ = c.Cookie("message")
		}
	}

	render(c, "index.html", gin.H{
		"expressions":      dataStruct,
		"is_logged_in":     true,
		"count_expression": CountExpression,
		"message":          message,
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
