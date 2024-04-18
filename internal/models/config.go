package models

import (
	"time"
)

const (
	StatusInProgress TaskStatus = "In progress"
	StatusCompleted  TaskStatus = "Completed"
	StatusError      TaskStatus = "Error"
)

type TaskStatus string

// Константы таймаутов вычислений по умолчанию
const (
	ADDTIMEOUT      = 5 * time.Second
	SUBTRACTTIMEOUT = 3 * time.Second
	MULTIPLYTIMEOUT = 4 * time.Second
	DIVIDETIMEOUT   = 6 * time.Second
)
