package database

import (
	"fmt"
	"log"
	"os"
	"time"

	"golang.org/x/crypto/bcrypt"
	"gorm.io/driver/sqlite"
	"gorm.io/gorm"
	"gorm.io/gorm/logger"
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
		&Expression{},
		&User{},
		&Timeouts{},
	)
	if err != nil {
		log.Println("Error migrating database: ", err)
		return err
	}

	dafaultTimeouts := &Timeouts{
		AddTimeout:      ADDTIMEOUT,
		SubtractTimeout: SUBTRACTTIMEOUT,
		MutiplyTimeout:  MULTIPLYTIMEOUT,
		DivideTimeout:   DIVIDETIMEOUT,
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

	user := User{Login: login, HashPassword: string(hashPassword)}
	result := dbExpr.Create(&user)
	if result.Error != nil {
		log.Println("Error creating user: ", result.Error)
		return 0, result.Error
	}

	return int(user.ID), nil
}

// UserExists проверяет, существует ли пользователь с указанным логином в базе данных
func UserExists(login string) bool {
	var user User
	dbExpr.First(&user, "login = ?", login)
	return user.ID != 0
}

// GetUser возвращает пользователя по его логину
func GetUser(login string) (User, error) {
	var user User
	result := dbExpr.First(&user, "login = ?", login)
	if result.Error != nil {
		return user, result.Error
	}
	return user, nil
}

// AddExprssion добавляет задачу в базу данных и возвращает её ID
func AddExprssion(userId uint, expression string) (uint, bool) {
	expr := Expression{UserId: userId, Expression: expression, Status: StatusInProgress}
	dbExpr.Create(&expr)

	return expr.InfoModel.ID, true
}

// SetTaskResult устанавливает результат выполнения задачи
func SetTaskResult(userId, exprId int, status TaskStatus, result float32) {
	var expression Expression
	db := dbExpr.First(&expression, "user_id = ? AND id = ?", userId, exprId)
	if db.Error != nil {
		log.Println("!!Error getting expression: ", db.Error)
		log.Println("!!User ID:", userId, "Expression ID:", exprId, "Expression:", expression)
		// Фатальная ошибка, возможно потеря вычислений, неверно передан параметр userId или exprId
		return
	}

	expression.Status = status
	if status == StatusCompleted {
		expression.Result = fmt.Sprintf("%v", result)
	}
	dbExpr.Save(&expression)
}

// GetAllTasks возвращает все задачи определённого юзера в базе данных
func GetAllTasks(userId uint) ([]Expression, error) {
	var expressions []Expression
	dbExpr.Find(&expressions, "user_id = ?", userId)
	return expressions, nil
}

// GetTask возвращает задачу по её идентификатору
func GetTask(userId uint, exprId int) (Expression, error) {
	var expression Expression
	dbExpr.First(&expression, "user_id = ? AND id = ?", userId, exprId)
	return expression, nil
}

// GetTimeouts возвращает таймауты вычислителя
func GetTimeouts() Timeouts {
	var timeouts Timeouts
	dbExpr.First(&timeouts)
	return timeouts
}

// SetTimeouts обновляет таймауты вычислителя
func SetTimeouts(add, subtract, multiply, divide time.Duration) {
	dbExpr.Model(&Timeouts{}).Where("id = ?", "1").Updates(
		Timeouts{
			AddTimeout:      add,
			SubtractTimeout: subtract,
			MutiplyTimeout:  multiply,
			DivideTimeout:   divide,
		},
	)
}

// GetAllUnfinishedTasks возвращает все незавершенные задачи из базы данных
func GetAllUnfinishedTasks() ([]Expression, error) {
	var expressions []Expression
	result := dbExpr.Find(&expressions, "status = ?", StatusInProgress)
	if result.Error != nil {
		return expressions, result.Error
	}
	return expressions, nil
}
