package database

import (
	"gorm.io/gorm"
	"time"
)

type InfoModel struct {
	ID        uint           `gorm:"primarykey"`
	CreatedAt time.Time      `json:"-"`
	UpdatedAt time.Time      `json:"-"`
	DeletedAt gorm.DeletedAt `gorm:"index" json:"-"`
}

type User struct {
	InfoModel
	Login        string
	HashPassword string
	Expressions  []Expression `gorm:"constraint:OnUpdate:CASCADE,OnDelete:CASCADE;"`
}

type Expression struct {
	InfoModel
	UserId     uint `json:"-"`
	Expression string
	Result     string
	Status     TaskStatus
}
