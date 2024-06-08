package models

import (
	"time"
)

type User struct {
	ID         int `gorm:"primaryKey"`
	CreatedAt  time.Time
	UpdatedAt  time.Time
	Name       string      `gorm:"type:varchar(100);not null"`
	Email      string      `gorm:"type:varchar(100);unique_index;not null"`
	Age        int         `gorm:"type:int;not null"`
	Password   string      `gorm:"type:varchar(512);not null"`
	Privileges []Privilege `gorm:"many2many:user_privileges;"`
}
