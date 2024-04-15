package frontend

import (
	"bytes"
	"encoding/json"
	"errors"
	"fmt"
	"io"
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

// showIndexPage отображает главную страницу
func showIndexPage(c *gin.Context) {
	// Запрос информации об операциях и её добавление в шаблон
	jwt, _ := c.Get("jwt_key")
	header := fmt.Sprintf("Bearer %s", jwt.(string))
	data, err := sendAPIRequest("/get-expressions/"+CountExpression, "GET", nil, header)
	if err != nil {
		log.Println("Error sendAPIRequest: ", err)
	}
	var dataStruct []models.Expression
	_ = json.Unmarshal(data, &dataStruct)

	render(c, "index.html", gin.H{
		"expressions":      dataStruct,
		"is_logged_in":     true,
		"count_expression": CountExpression,
	})
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

	c.SetCookie("jwt_key", token.Token, int(time.Hour*24), "/", "", false, true)
	c.SetCookie("user_id", token.UserID, int(time.Hour*24), "/", "", false, true)
	c.Redirect(302, "/")
}

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
	body, err := io.ReadAll(resp.Body)
	if err != nil {
		return nil, err
	}
	return body, nil
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
