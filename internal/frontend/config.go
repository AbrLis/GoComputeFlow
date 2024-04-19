package frontend

import (
	"GoComputeFlow/internal/models"
	"errors"
)

const (
	APIPath         = "http://localhost:3000/api/v1"
	CountExpression = 11 // Количество выражений на главной странице
)

var (
	errorAPI            = errors.New("API error")
	errorsTimeout       = errors.New("timeouts parsing error")
	errorParsingTimeout = errors.New("parsing timeouts error")
	errorUnauthorized   = errors.New("errorUnauthorized")
)
var timeLifeCookie = 2

type indexPageData struct {
	Expressions []models.Expression
	Message     string
	IsNext      bool
	MyPage      int
}

type indexPageMonitoring struct {
	Message string
	Add     float32
	Sub     float32
	Mul     float32
	Div     float32
}
