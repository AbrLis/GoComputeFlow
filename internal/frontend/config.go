package frontend

import "errors"

const (
	APIPath         = "http://localhost:3000/api/v1"
	CountExpression = "10" // Количество выражений на главной странице
)

var errorsTimeout = errors.New("timeouts parsing error")
