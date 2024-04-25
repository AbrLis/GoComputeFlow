package auth

import (
	"github.com/gin-gonic/gin"
	"golang.org/x/crypto/bcrypt"

	"GoComputeFlow/internal/database"
)

type Error struct {
	Code int
	Msg  gin.H
}

// CheckUserExists проверяет, существует ли пользователь с указанным логином и паролем
func CheckUserExists(login, password string) (bool, Error) {
	user, err := database.GetUser(login)
	if err != nil {
		return false, Error{Code: 404, Msg: gin.H{"user not found": err.Error()}}
	}
	err = bcrypt.CompareHashAndPassword([]byte(user.HashPassword), []byte(password))
	if err != nil {
		return false, Error{Code: 401, Msg: gin.H{"error password": err.Error()}}
	}

	// Если всё в порядке, то в поле Code передаю ID пользователя
	return true, Error{Code: int(user.ID), Msg: nil}
}
