package auth

import (
	"strings"

	"github.com/gin-gonic/gin"
	"github.com/golang-jwt/jwt/v5"

	"GoComputeFlow/internal/api/apiConfig"
	"GoComputeFlow/internal/database"
)

func EnsureAuth() gin.HandlerFunc {
	return func(c *gin.Context) {
		token := c.Request.Header.Get("Authorization")
		jwtTokenString := strings.Split(token, " ")
		if token == "" || len(jwtTokenString) != 2 {
			c.JSON(401, gin.H{"message": "Unauthorized, required jwt token"})
			c.Abort()
			return
		}

		// Валидация токена
		tokenJWT := jwtTokenString[1]

		tokenFromString, err := jwt.Parse(
			tokenJWT, func(token *jwt.Token) (interface{}, error) {
				if _, ok := token.Method.(*jwt.SigningMethodHMAC); !ok {
					return nil, jwt.ErrSignatureInvalid
				}
				return []byte(apiConfig.SECRETKEY), nil
			},
		)

		if err != nil || !database.TokenExists(tokenJWT) {
			c.JSON(401, gin.H{"message": "Unauthorized: invalid token"})
			c.Abort()
			return
		}

		// Добавление id пользователя в контекст
		if claims, ok := tokenFromString.Claims.(jwt.MapClaims); ok && tokenFromString.Valid {
			if userId, ok := claims["user_id"]; ok {
				temp, ok := userId.(float64)
				if ok {
					c.Set("user_id", uint(temp))
					return
				}
			}
		}
		c.JSON(500, gin.H{"error": "Ошибка обработки jwt токена"})
		c.Abort()
	}
}
