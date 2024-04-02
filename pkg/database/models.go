package database

import (
	"gorm.io/gorm"
)

type User struct {
	gorm.Model
	Login        string
	HashPassword string
	Expressions  []Expression `gorm:"constraint:OnUpdate:CASCADE,OnDelete:CASCADE;"`
}

type Expression struct {
	gorm.Model `json:"-"`
	UserId     uint `json:"-"`
	Expression string
	Result     string
	Status     TaskStatus
}
