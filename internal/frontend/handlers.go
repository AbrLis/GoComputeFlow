package frontend

import (
	"bytes"
	"encoding/json"
	"errors"
	"io"
	"log"
	"net/http"
	"time"

	"github.com/gin-gonic/gin"

	"GoComputeFlow/internal/api"
)

type tokenJWT struct {
	Token string `json:"token"`
}

func render(c *gin.Context, templateName string, data gin.H) {
	jwt, err := c.Cookie("jwt_key")
	log.Println(jwt)
	if err != nil {
		// переходим на логин экран
		c.Redirect(http.StatusFound, "/login")
	}

	// TODO: Запрос информации и её отображение
	c.HTML(200, templateName, data)
}

func showIndexPage(c *gin.Context) {
	render(c, "index.html", nil)
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
	request, _ := http.NewRequest("POST", "http://localhost:3000/api/v1/register", bytes.NewReader(data))
	request.Header.Set("Content-Type", "application/json")
	resp, err := http.DefaultClient.Do(request)
	resp.Body.Close()

	// обращение к API логина
	request, _ = http.NewRequest("POST", "http://localhost:3000/api/v1/login", bytes.NewReader(data))
	request.Header.Set("Content-Type", "application/json")
	resp, err = http.DefaultClient.Do(request)
	if err != nil {
		errorLogin(c, "Error", err)
		return
	}

	body, err := io.ReadAll(resp.Body)
	if err != nil {
		errorLogin(c, "Error", err)
		return
	}
	resp.Body.Close()
	var token tokenJWT
	err = json.Unmarshal(body, &token)
	if err != nil {
		errorLogin(c, "Error", err)
		return
	}

	// Логин успешно пройден, установить куку и перейти на глвную
	c.SetCookie("jwt_key", token.Token, int(time.Hour*24), "/", "", false, true)
	c.Redirect(302, "/")
}

func errorLogin(c *gin.Context, errTitle string, err error) {

	c.HTML(http.StatusBadRequest, "login.html", gin.H{
		"ErrorTitle":   errTitle,
		"ErrorMessage": err,
	})
}
