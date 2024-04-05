package worker

import (
	"sync"
	"time"
)

// Константы таймаутов вычислений по умолчанию
const (
	ADDTIMEOUT      = 5 * time.Second
	SUBTRACTTIMEOUT = 3 * time.Second
	MULTIPLYTIMEOUT = 4 * time.Second
	DIVIDETIMEOUT   = 6 * time.Second
)

const (
	COUNTWORKERS     = 5
	COUNTWORKERSFREE = 5
)

// TaskCalculate - структура для формирования задачи
type TaskCalculate struct {
	ID         uint    `json:"id"`
	Expression []Token `json:"expression"`
}

// Token - структура для формирования польской нотации выражения
type Token struct {
	Value string
	IsOp  bool
}

// Result - структура для очереди ответа
type Result struct {
	ID        uint
	FlagError bool
	Result    float64
}

type Worker struct {
	Count           int
	CountFree       int
	Queue           []TaskCalculate
	ResultQueue     []Result
	taskChannel     chan TaskCalculate
	PingTimeoutCalc []time.Time
	AddTimeout      time.Duration
	SubtractTimeout time.Duration
	MultiplyTimeout time.Duration
	DivideTimeout   time.Duration
	Mu              sync.Mutex
}