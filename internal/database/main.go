package database

import (
	"fmt"
	"log"
	"os"
	"time"

	"github.com/glebarez/sqlite"
	"golang.org/x/crypto/bcrypt"
	"gorm.io/gorm"
	"gorm.io/gorm/logger"

	"GoComputeFlow/internal/models"
)

var dbExpr *gorm.DB

func OpenDB() error {
	var err error
	gormLogger := logger.New(
		log.New(os.Stdout, "\r\n", log.LstdFlags),
		logger.Config{
			LogLevel:                  logger.Silent,
			IgnoreRecordNotFoundError: true,
		},
	)

	dbExpr, err = gorm.Open(sqlite.Open("database.db"), &gorm.Config{Logger: gormLogger})
	if err != nil {
		log.Println("Error opening database: ", err)
		return err
	}

	err = dbExpr.AutoMigrate(
		&models.Expression{},
		&models.User{},
		&models.Timeouts{},
	)
	if err != nil {
		log.Println("Error migrating database: ", err)
		return err
	}

	dafaultTimeouts := &models.Timeouts{
		AddTimeout:      models.ADDTIMEOUT,
		SubtractTimeout: models.SUBTRACTTIMEOUT,
		MutiplyTimeout:  models.MULTIPLYTIMEOUT,
		DivideTimeout:   models.DIVIDETIMEOUT,
	}
	dbExpr.FirstOrCreate(dafaultTimeouts)

	return nil
}

// CreateNewUser создает нового пользователя в базе данных и возвращает его ID
func CreateNewUser(login, password string) (int, error) {
	if UserExists(login) {
		log.Println("User already exists: ", login)
		return 0, fmt.Errorf("user already exists: %s", login)
	}

	hashPassword, err := bcrypt.GenerateFromPassword([]byte(password), bcrypt.DefaultCost)
	if err != nil {
		log.Println("Error hashing password: ", err)
		return 0, err
	}

	user := models.User{Login: login, HashPassword: string(hashPassword)}
	result := dbExpr.Create(&user)
	if result.Error != nil {
		log.Println("Error creating user: ", result.Error)
		return 0, result.Error
	}

	return int(user.ID), nil
}

// UserExists проверяет, существует ли пользователь с указанным логином в базе данных
func UserExists(login string) bool {
	var user models.User
	dbExpr.First(&user, "login = ?", login)
	return user.ID != 0
}

// GetUser возвращает пользователя по его логину
func GetUser(login string) (models.User, error) {
	var user models.User
	result := dbExpr.First(&user, "login = ?", login)
	if result.Error != nil {
		return user, result.Error
	}
	return user, nil
}

// AddExprssion добавляет задачу в базу данных и возвращает её ID
func AddExprssion(userId uint, expression string) (uint, bool) {
	expr := models.Expression{UserId: userId, Expression: expression, Status: models.StatusInProgress}
	dbExpr.Create(&expr)

	return expr.InfoModel.ID, true
}

// SetTaskResult устанавливает результат выполнения задачи
func SetTaskResult(userId, exprId int, status models.TaskStatus, result float32) {
	var expression models.Expression
	db := dbExpr.First(&expression, "user_id = ? AND id = ?", userId, exprId)
	if db.Error != nil {
		log.Println("!!Error getting expression: ", db.Error)
		log.Println("!!User ID:", userId, "Expression ID:", exprId, "Expression:", expression)
		// Фатальная ошибка, возможно потеря вычислений, неверно передан параметр userId или exprId
		return
	}

	expression.Status = status
	if status == models.StatusCompleted {
		expression.Result = fmt.Sprintf("%v", result)
	}
	dbExpr.Save(&expression)
}

// GetNTasks возвращает limit задач на странице page
func GetNTasks(userId uint, limit, page int) ([]models.Expression, error) {
	var expressions []models.Expression
	offset := (page - 1) * limit
	db := dbExpr.Order("id desc").Limit(limit).Offset(offset).Find(&expressions, "user_id = ?", userId)
	return expressions, db.Error
}

// GetTask возвращает задачу по её идентификатору
func GetTask(userId uint, exprId int) (models.Expression, error) {
	var expression models.Expression
	db := dbExpr.First(&expression, "user_id = ? AND id = ?", userId, exprId)
	return expression, db.Error
}

// GetTimeouts возвращает таймауты вычислителя
func GetTimeouts() models.Timeouts {
	var timeouts models.Timeouts
	dbExpr.First(&timeouts)
	return timeouts
}

// SetTimeouts обновляет таймауты вычислителя
func SetTimeouts(add, subtract, multiply, divide time.Duration) {
	dbExpr.Model(&models.Timeouts{}).Where("id = ?", "1").Updates(
		models.Timeouts{
			AddTimeout:      add,
			SubtractTimeout: subtract,
			MutiplyTimeout:  multiply,
			DivideTimeout:   divide,
		},
	)
}

// GetAllUnfinishedTasks возвращает все незавершенные задачи из базы данных
func GetAllUnfinishedTasks() ([]models.Expression, error) {
	var expressions []models.Expression
	result := dbExpr.Find(&expressions, "status = ?", models.StatusInProgress)
	if result.Error != nil {
		return expressions, result.Error
	}
	return expressions, nil
}
