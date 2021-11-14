package models

import (
	"gorm.io/gorm"
	"time"
)

type Url struct {
	ID          uint           `json:"id" gorm:"primaryKey"`
	CreatedAt   time.Time      `json:"createdAt"`
	UpdatedAt   time.Time      `json:"updatedAt"`
	DeletedAt   gorm.DeletedAt `json:"deletedAt"`
	Key         string         `json:"key"`
	Destination string         `json:"destination"`
	UserID      uint           `json:"userId"`
	User        *User          `json:"user"`
}
