package database

const (
	StatusInProgress TaskStatus = "In progress"
	StatusCompleted  TaskStatus = "Completed"
	StatusError      TaskStatus = "Error"
)

type TaskStatus string
