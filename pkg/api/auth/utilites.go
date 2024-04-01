package auth

import (
	"GoComputeFlow/pkg/database"
	"github.com/gin-gonic/gin"
	"golang.org/x/crypto/bcrypt"
)

type Error struct {
	Code int
	Msg  gin.H
}

// CheckUserExists проверяет, существует ли пользователь с указанным логином и паролем
func CheckUserExists(login, password string) (bool, Error) {
	hashedPasswordLogin, err := bcrypt.GenerateFromPassword([]byte(password), bcrypt.DefaultCost)
	if err != nil {
		return false, Error{Code: 500, Msg: gin.H{"error login": err.Error()}}
	}
	user, err := database.GetUser(string(hashedPasswordLogin))

	// Проверка, что такой пользователь зарегистрирован
	if user.Login != login {
		return false, Error{Code: 401, Msg: gin.H{"error": "Пользователь не зарегистрирован"}}
	}

	// Проверка пароля
	if string(hashedPasswordLogin) != user.HashPassword {
		return false, Error{Code: 401, Msg: gin.H{"error": "Неверный пароль"}}
	}

	return true, Error{}
}
