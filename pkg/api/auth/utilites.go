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
	user, err := database.GetUser(login)
	if err != nil {
		return false, Error{Code: 500, Msg: gin.H{"error login db error": err.Error()}}
	}
	err = bcrypt.CompareHashAndPassword([]byte(user.HashPassword), []byte(password))
	if err != nil {
		return false, Error{Code: 401, Msg: gin.H{"error password": err.Error()}}
	}

	return true, Error{}
}
