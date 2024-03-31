package database

const (
	StatusInProgress TaskStatus = iota
	StatusCompleted
	StatusError
)

type TaskStatus int
