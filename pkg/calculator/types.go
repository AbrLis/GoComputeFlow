package calculator

import (
	"sync"
	"time"
)

var done chan struct{} // Канал завершения вычислительных операций

// Token - структура для формирования польской нотации выражения
type Token struct {
	Value string
	IsOp  bool
}

// TaskCalculate - структура для формирования задачи
type TaskCalculate struct {
	ID         uint    `json:"id"`
	Expression []Token `json:"expression"`
}

// FreeCalculators - Структура счётчика свободных вычислителей
type FreeCalculators struct {
	Count           int                      // Количество вычислителей
	CountFree       int                      // Свободные вычислители
	PingTimeoutCalc []time.Time              // Таймауты пингов вычислителей
	Queue           []TaskCalculate          // Очередь исполнения задач
	queueInProcess  map[string]TaskCalculate // Задачи находящиеся на обработке
	taskChannel     chan TaskCalculate       // Канал задач
	AddTimeout      time.Duration            // Таймауты операций
	SubtractTimeout time.Duration
	MultiplyTimeout time.Duration
	DivideTimeout   time.Duration
	mu              sync.Mutex
}
