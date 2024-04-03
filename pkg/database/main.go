package database

import (
	"fmt"
	"golang.org/x/crypto/bcrypt"
	"gorm.io/driver/sqlite"
	"gorm.io/gorm"
	"log"
)

var dbExpr *gorm.DB

func OpenDB() error {
	var err error
	dbExpr, err = gorm.Open(sqlite.Open("database.db"), &gorm.Config{})
	if err != nil {
		log.Println("Error opening database: ", err)
		return err
	}

	err = dbExpr.AutoMigrate(&Expression{}, &User{})
	if err != nil {
		log.Println("Error migrating database: ", err)
		return err
	}

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
	dbExpr.First(&user, "login = ?", login)
	return user, nil
}

// AddExprssion добавляет задачу в базу данных
func AddExprssion(userId uint, expression string) bool {
	expr := Expression{UserId: userId, Expression: expression, Status: StatusInProgress}
	dbExpr.Create(&expr)

	return true
}

// SetTaskResult устанавливает результат выполнения задачи
func SetTaskResult(userId, exprId int, status TaskStatus) {
	var expression Expression
	dbExpr.First(&expression, "user_id = ? AND id = ?", userId, exprId)
	expression.Status = status
	dbExpr.Save(&expression)
}

// GetAllTasks возвращает все задачи определённого юзера в базе данных
func GetAllTasks(userId uint) ([]Expression, error) {
	var expressions []Expression
	dbExpr.Find(&expressions, "user_id = ?", userId)
	return expressions, nil
}

// GetTaskStatus возвращает статус задачи определённого юзера в базе данных
func GetTaskStatus(userId, exprId int) string {

	var expression Expression
	dbExpr.First(&expression, "user_id = ? AND id = ?", userId, exprId)
	if expression.Status == StatusInProgress {
		return "In progress"
	} else if expression.Status == StatusCompleted {
		return "Completed"
	} else if expression.Status == StatusError {
		return "Error"
	}
	return "Unknown"
}

// GetTask возвращает задачу по её идентификатору
func GetTask(userId uint, exprId int) (Expression, error) {
	var expression Expression
	dbExpr.First(&expression, "user_id = ? AND id = ?", userId, exprId)
	return expression, nil
}
