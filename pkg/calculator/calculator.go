package calculator

import (
	"GoComputeFlow/pkg/database"
	"fmt"
	"log"
	"sync"
	"time"
)

var Calc *FreeCalculators

// CreateCalculators создает новый экземпляр структуры счётчика свободных вычислителей
func CreateCalculators() {
	Calc = &FreeCalculators{
		Count:           5,
		CountFree:       5,
		Queue:           []TaskCalculate{},
		queueInProcess:  map[string]TaskCalculate{},
		taskChannel:     make(chan TaskCalculate),
		AddTimeout:      ADDTIMEOUT,
		SubtractTimeout: SUBTRACTTIMEOUT,
		MultiplyTimeout: MULTIPLYTIMEOUT,
		DivideTimeout:   DIVIDETIMEOUT,
		mu:              sync.Mutex{},
	}
	Calc.PingTimeoutCalc = make([]time.Time, Calc.Count)

	// TODO: Добавить запуск распределителя вычислений в своём потоке, которой будет следить за очередью задач и
	// передавать задачи вычислителям, так же будет получать от них ответы и заносить результаты в бд
}

// AddExpressionToQueue добавляет выражение в очередь задач
func AddExpressionToQueue(expression string, userId uint) bool {
	// Парсим выражение
	tokens, err := ParseExpression(expression)
	if err != nil {
		log.Println("Error parsing expression: ", err)
		return false
	}

	// Добавляю задачу в очередь
	Calc.mu.Lock()
	Calc.Queue = append(Calc.Queue, TaskCalculate{ID: userId, Expression: tokens})
	Calc.mu.Unlock()

	// Добавляю задачу в список вычислений юзера в базу данных
	if ok := database.AddExprssion(userId, expression); !ok {
		return false
	}

	return true
}

// GetTimeoutsOperations возвращает время вычислений для каждой из операций
func GetTimeoutsOperations() map[string]string {
	return map[string]string{
		"+": fmt.Sprintf("%.2f sec", Calc.AddTimeout.Seconds()),
		"-": fmt.Sprintf("%.2f sec", Calc.SubtractTimeout.Seconds()),
		"*": fmt.Sprintf("%.2f sec", Calc.MultiplyTimeout.Seconds()),
		"/": fmt.Sprintf("%.2f sec", Calc.DivideTimeout.Seconds()),
	}
}
