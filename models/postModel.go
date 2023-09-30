package models

import (
	"time"

	"gorm.io/gorm"
)

type User struct {
	gorm.Model

	ID        uint    `gorm:"primaryKey;required"`
	Username  string  `gorm:"required"`
	Email     string  `gorm:"unique;required"`
	Password  string  `gorm:"required;minLength:6"`
	Photos    []Photo `gorm:"foreignKey:UserID"`
	CreatedAt time.Time
	UpdatedAt time.Time
	DeletedAt gorm.DeletedAt `gorm:"index"`
}

type Photo struct {
	gorm.Model

	ID       uint   `gorm:"primaryKey"`
	Title    string `gorm:"required"`
	Caption  string
	PhotoURL string `gorm:"required"`
	UserID   uint
}
