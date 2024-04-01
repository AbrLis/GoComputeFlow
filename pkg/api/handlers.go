package api

import (
	"log"
	"os"
	"time"

	"GoComputeFlow/pkg/api/auth"
	"GoComputeFlow/pkg/database"
	"github.com/gin-gonic/gin"
	"github.com/golang-jwt/jwt/v5"
)

var SECRETKEY = os.Getenv("SECRET_KEY")

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

	// Проверка, существует ли пользователь с указанным логином и паролем
	if ok, msg := auth.CheckUserExists(req.Login, req.Password); !ok {
		c.JSON(msg.Code, msg.Msg)
		return
	}

	// Создание токена
	now := time.Now()
	token := jwt.NewWithClaims(
		jwt.SigningMethodHS256, jwt.MapClaims{
			"login": req.Login,
			"nbf":   now.Unix(),
			"exp":   now.Add(time.Hour * 24).Unix(),
			"iat":   now.Unix(),
		},
	)

	tokenString, err := token.SignedString([]byte(SECRETKEY))
	if err != nil {
		c.JSON(500, gin.H{"error": err.Error()})
		return
	}

	log.Println("Успешно получен токен для пользователя", req.Login)
	c.JSON(200, gin.H{"token": tokenString})
}
