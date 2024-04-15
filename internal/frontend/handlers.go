package frontend

import (
	"bytes"
	"encoding/json"
	"errors"
	"io"
	"log"
	"net/http"
	"strconv"
	"time"

	"github.com/gin-gonic/gin"

	"GoComputeFlow/internal/api"
	"GoComputeFlow/internal/database"
)

type tokenJWT struct {
	Token  string `json:"token"`
	UserID string `json:"user_id"`
}

func render(c *gin.Context, templateName string, data gin.H) {
	c.HTML(200, templateName, data)
}

func showIndexPage(c *gin.Context) {
	// TODO: Запрос информации об операциях и её добавление в шаблон
	userID, exist := c.Get("user_id")
	userIDUint, err := strconv.ParseUint(userID.(string), 10, 64)
	if err != nil {
		log.Println("Ошибка преобразования user_id: ", err)
		c.Redirect(http.StatusFound, "/login")
		c.Abort()
	}
	var data []database.Expression
	if exist {
		data, _ = database.GetNTasks(uint(userIDUint), 10)
	}
	render(c, "index.html", gin.H{
		"expressions":  data,
		"is_logged_in": true,
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
	request, _ := http.NewRequest("POST", APIPath+"/register", bytes.NewReader(data))
	request.Header.Set("Content-Type", "application/json")
	if resp, err := http.DefaultClient.Do(request); err == nil {
		resp.Body.Close()
	}

	// обращение к API логина
	request, _ = http.NewRequest("POST", APIPath+"/login", bytes.NewReader(data))
	request.Header.Set("Content-Type", "application/json")
	resp, err := http.DefaultClient.Do(request)
	if err != nil {
		errorLogin(c, "Error", err)
		return
	}
	defer resp.Body.Close()

	body, err := io.ReadAll(resp.Body)
	if err != nil {
		errorLogin(c, "Error", err)
		return
	}
	var token tokenJWT
	err = json.Unmarshal(body, &token)
	if err != nil {
		errorLogin(c, "Error", err)
		return
	}

	// Логин успешно пройден, установить куку и перейти на глвную

	c.SetCookie("jwt_key", token.Token, int(time.Hour*24), "/", "", false, true)
	c.SetCookie("user_id", token.UserID, int(time.Hour*24), "/", "", false, true)
	c.Redirect(302, "/")
}

func errorLogin(c *gin.Context, errTitle string, err error) {

	c.HTML(http.StatusBadRequest, "login.html", gin.H{
		"ErrorTitle":   errTitle,
		"ErrorMessage": err,
	})
}

func logOut(c *gin.Context) {
	c.SetCookie("jwt_key", "", -1, "/", "", false, true)
	c.Set("is_logged_in", false)
	c.Redirect(302, "/login")
}
