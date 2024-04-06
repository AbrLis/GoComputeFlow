package api

import (
	"GoComputeFlow/internal/calculator"
	"io"
	"log"
	"os"
	"strconv"
	"time"

	"GoComputeFlow/internal/api/auth"
	"GoComputeFlow/internal/database"
	"github.com/gin-gonic/gin"
	"github.com/golang-jwt/jwt/v5"
)

var SECRETKEY = os.Getenv("SECRETKEY")

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
	c.JSON(200, gin.H{"msg": "Пользователь с логином " + req.Login + " успешно зарегистрирован"})

}

// LoginUser получение токена по логину и паролю
func LoginUser(c *gin.Context) {
	var req registerUserRequest
	if err := c.ShouldBindJSON(&req); err != nil {
		c.JSON(400, gin.H{"Необходимые поля для получения токена: login, password": err.Error()})
		return
	}

	// Проверка, существует ли пользователь с указанным логином и паролем
	ok, msg := auth.CheckUserExists(req.Login, req.Password)
	if !ok {
		c.JSON(msg.Code, msg.Msg)
		return
	}

	// Создание токена
	now := time.Now()
	token := jwt.NewWithClaims(
		jwt.SigningMethodHS256, jwt.MapClaims{
			"user_id": msg.Code, // ID пользователя
			"login":   req.Login,
			"nbf":     now.Unix(),
			"exp":     now.Add(time.Hour * 24).Unix(),
			"iat":     now.Unix(),
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

// AddExpressionHandler обработчик для добавления арифметического выражения
func AddExpressionHandler(c *gin.Context) {
	bodyBytes, err := io.ReadAll(c.Request.Body)
	if err != nil || len(bodyBytes) == 0 {
		c.JSON(500, gin.H{"error": "Empty body or error reading body"})
		return
	}

	// Отправка строки калькулятору
	userId, ok := c.Get("user_id")
	if !ok {
		c.JSON(500, gin.H{"error": "invalid user_id in context"})
		return
	}
	if ok := calculator.AddExpressionToQueue(string(bodyBytes), userId.(uint)); !ok {
		c.JSON(500, gin.H{"error": "Error parsing exprssion"})
		return
	}

	c.JSON(200, gin.H{"msg": "Expression added to queue: " + string(bodyBytes)})
}

// GetExpressionsHandler обработчик для получения списка арифметических выражений пользователя
func GetExpressionsHandler(c *gin.Context) {
	userId, ok := c.Get("user_id")
	if !ok {
		c.JSON(500, gin.H{"error": "invalid userID"})
		return
	}

	expressions, err := database.GetAllTasks(userId.(uint))
	if err != nil {
		c.JSON(500, gin.H{"error": "Ошибка получения списка из базы данных"})
		return
	}

	c.JSON(200, expressions)
}

// GetOperationsHandler обработчик для получения списка времени выполнения операций
func GetOperationsHandler(c *gin.Context) {
	c.JSON(200, calculator.GetTimeoutsOperations())
}

// GetValueHandler возвращает конкретную задачу по ID
func GetValueHandler(c *gin.Context) {
	userId, ok := c.Get("user_id")
	if !ok {
		c.JSON(500, gin.H{"error": "invalid userID"})
		return
	}
	taskID, err := strconv.Atoi(c.Param("task_id"))
	if err != nil {
		c.JSON(400, gin.H{"error": "invalid taskID"})
		return
	}
	result, _ := database.GetTask(userId.(uint), taskID)
	if result.UserId == 0 {
		c.JSON(404, gin.H{"error": "task not found"})
		return
	}
	c.JSON(200, result)
}
