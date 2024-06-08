package models

import "time"

type Privilege struct {
	ID          int    `gorm:"primaryKey"`
	Description string `gorm:"type:varchar(256);not null"`
	Name        string `gorm:"unique;not null;type:varchar(256)"`
	CreatedAt   time.Time
	UpdatedAt   time.Time
}
