package models

import (
	"time"

	"gorm.io/gorm"
)

type InfoModel struct {
	ID        uint           `gorm:"primarykey"`
	CreatedAt time.Time      `json:"-"`
	UpdatedAt time.Time      `json:"-"`
	DeletedAt gorm.DeletedAt `gorm:"index" json:"-"`
}

// User Таблица пользователя
type User struct {
	InfoModel
	Login        string
	Token        string
	HashPassword string
	Expressions  []Expression `gorm:"constraint:OnUpdate:CASCADE,OnDelete:CASCADE;"`
}

// Expression Таблица представления выражений
type Expression struct {
	InfoModel
	UserId     uint `json:"-"`
	Expression string
	Result     string
	Status     TaskStatus
}

// Timeouts таблица для сохранения таймаутов воркеров между перезапусками
type Timeouts struct {
	InfoModel
	AddTimeout      time.Duration
	SubtractTimeout time.Duration
	MutiplyTimeout  time.Duration
	DivideTimeout   time.Duration
}
