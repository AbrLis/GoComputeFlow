package api

import (
	"io"
	"log"
	"net/http"
	"strconv"
	"time"

	"github.com/gin-gonic/gin"
	"github.com/golang-jwt/jwt/v5"

	"GoComputeFlow/internal/api/apiConfig"
	"GoComputeFlow/internal/api/auth"
	"GoComputeFlow/internal/calculator"
	"GoComputeFlow/internal/database"
	"GoComputeFlow/internal/models"
)

type RegisterUserRequest struct {
	Login    string `json:"login"`
	Password string `json:"password"`
}

// RegisterUser регистрация нового пользователя по логину и паролю
func RegisterUser(c *gin.Context) {
	var req RegisterUserRequest
	if err := c.ShouldBindJSON(&req); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"Необходимые поля для регистрации: login, password": err.Error()})
		return
	}

	id, err := database.CreateNewUser(req.Login, req.Password)
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
		return
	}

	log.Println("Успешно зарегистрирован пользователь с ID", id)
	c.JSON(http.StatusOK, gin.H{"msg": "Пользователь с логином " + req.Login + " успешно зарегистрирован"})

}

// LoginUser получение токена по логину и паролю
func LoginUser(c *gin.Context) {
	var req RegisterUserRequest
	if err := c.ShouldBindJSON(&req); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"Необходимые поля для получения токена: login, password": err.Error()})
		return
	}
	if len(req.Login) <= 3 || len(req.Password) <= 3 {
		c.JSON(http.StatusBadRequest, gin.H{"error": "length of username and password must be > 3"})
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

	tokenString, err := token.SignedString([]byte(apiConfig.SECRETKEY))
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
		return
	}

	// Добавим токен в бд
	err = database.AddTokenToUser(uint(msg.Code), tokenString)
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
		return
	}

	log.Println("Успешно получен токен для пользователя", req.Login)
	c.JSON(http.StatusOK, gin.H{"token": tokenString, "user_id": strconv.Itoa(msg.Code)})
}

// AddExpressionHandler обработчик для добавления арифметического выражения
func AddExpressionHandler(c *gin.Context) {
	bodyBytes, err := io.ReadAll(c.Request.Body)
	if err != nil || len(bodyBytes) == 0 {
		c.JSON(http.StatusBadRequest, gin.H{"error": "Empty body or error reading body"})
		return
	}

	// Отправка строки калькулятору
	userId, ok := c.Get("user_id")
	if !ok {
		c.JSON(http.StatusInternalServerError, gin.H{"error": "invalid user_id in context"})
		return
	}
	if ok := calculator.AddExpressionToQueue(string(bodyBytes), userId.(uint), true, 0); !ok {
		c.JSON(http.StatusBadRequest, gin.H{"error": "Error parsing exprssion"})
		return
	}

	c.JSON(http.StatusOK, gin.H{"msg": "Expression added to queue: " + string(bodyBytes)})
}

// GetExpressionsHandler обработчик для получения списка арифметических выражений пользователя
func GetExpressionsHandler(c *gin.Context) {
	userId, ok := c.Get("user_id")
	if !ok {
		c.JSON(http.StatusInternalServerError, gin.H{"error": "invalid userID"})
		return
	}

	var (
		expressions []models.Expression
		err         error
		limit       = apiConfig.DefaultCountExpressions
		page        = 1
	)

	limit, err = getQueryParameter(c, "limit", limit)
	if err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": "Неверное значение limit"})
		return
	}
	page, err = getQueryParameter(c, "page", page)
	if err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": "Неверное значение page"})
		return
	}
	offset, _ := getQueryParameter(c, "offset", (page-1)*limit)

	expressions, err = database.GetNTasks(userId.(uint), limit, offset)
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
		return
	}

	c.JSON(http.StatusOK, expressions)
}

// GetOperationsHandler обработчик для получения списка времени выполнения операций
func GetOperationsHandler(c *gin.Context) {
	c.JSON(http.StatusOK, calculator.GetTimeoutsOperations())
}

// SetOperationsHandler обработчик для установки времени выполнения операции
func SetOperationsHandler(c *gin.Context) {
	add := c.Query("add")
	sub := c.Query("sub")
	mul := c.Query("mul")
	div := c.Query("div")

	res, err := calculator.SetTimeoutsOperations(add, sub, mul, div)
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
	}
	c.JSON(200, map[string]string{"msg": res})
}

// GetValueHandler возвращает конкретную задачу по ID
func GetValueHandler(c *gin.Context) {
	userId, ok := c.Get("user_id")
	if !ok {
		c.JSON(http.StatusInternalServerError, gin.H{"error": "invalid userID"})
		return
	}
	taskID, err := strconv.Atoi(c.Param("task_id"))
	if err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": "invalid taskID"})
		return
	}
	result, _ := database.GetTask(userId.(uint), taskID)
	if result.UserId == 0 {
		c.JSON(http.StatusNotFound, gin.H{"error": "task not found"})
		return
	}
	c.JSON(http.StatusOK, result)
}

// GetMonitoringHandler возвращает мониторинг задач
func GetMonitoringHandler(c *gin.Context) {
	data, err := calculator.GetWorkersTimeouts()
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
		return
	}
	c.JSON(http.StatusOK, data)
}
