package api

import (
	"log"
	"strconv"
	"time"

	"GoComputeFlow/pkg/database"
	"github.com/gin-gonic/gin"
	"github.com/golang-jwt/jwt/v5"
)

type registerUserRequest struct {
	Login    string `json:"login"`
	Password string `json:"password"`
}

// RegisterUser регистрация нового пользователя по логину и паролю
func RegisterUser(c *gin.Context) {
	var req registerUserRequest
	if err := c.ShouldBindJSON(&req); err != nil {
		c.JSON(400, gin.H{"Необходимые поля для регистрации: login, password": err.Error()})
		return
	}

	id, err := database.CreateNewUser(req.Login, req.Password)
	if err != nil {
		c.JSON(500, gin.H{"error": err.Error()})
		return
	}

	log.Println("Успешно зарегистрирован пользователь с ID", id)
}

// LoginUser получение токена по логину и паролю
func LoginUser(c *gin.Context) {
	var req registerUserRequest
	if err := c.ShouldBindJSON(&req); err != nil {
		c.JSON(400, gin.H{"Необходимые поля для получения токена: login, password": err.Error()})
		return
	}

	// Проверка, что такой пользователь зарегистрирован
	if !database.UserExists(req.Login) {
		c.JSON(401, gin.H{"error": "Пользователь не зарегистрирован"})
		return
	}

	// Создание токена
	secret := strconv.FormatInt(time.Now().UnixNano(), 10)
	now := time.Now()
	token := jwt.NewWithClaims(
		jwt.SigningMethodHS256, jwt.MapClaims{
			"login": req.Login,
			"nbf":   now.Unix(),
			"exp":   now.Add(time.Hour * 24).Unix(),
			"iat":   now.Unix(),
		},
	)

	tokenString, err := token.SignedString([]byte(secret))
	if err != nil {
		c.JSON(500, gin.H{"error": err.Error()})
		return
	}

	log.Println("Успешно получен токен для пользователя", req.Login)
	c.JSON(200, gin.H{"token": tokenString})
}
