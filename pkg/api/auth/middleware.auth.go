package auth

import (
	"os"
	"strings"

	"github.com/gin-gonic/gin"
	"github.com/golang-jwt/jwt/v5"
)

var SECRETKEY = os.Getenv("SECRETKEY")

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

		_, err := jwt.Parse(
			tokenJWT, func(token *jwt.Token) (interface{}, error) {
				if _, ok := token.Method.(*jwt.SigningMethodHMAC); !ok {
					return nil, jwt.ErrSignatureInvalid
				}
				return []byte(SECRETKEY), nil
			},
		)

		if err != nil {
			c.JSON(401, gin.H{"message": "Unauthorized: " + err.Error()})
			c.Abort()
			return
		}
	}
}
