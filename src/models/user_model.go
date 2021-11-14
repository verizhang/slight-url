package models

import (
	"gorm.io/gorm"
	"time"
)

type User struct {
	ID        uint           `json:"id" gorm:"primaryKey"`
	CreatedAt time.Time      `json:"createdAt"`
	UpdatedAt time.Time      `json:"updatedAt"`
	DeletedAt gorm.DeletedAt `json:"deletedAt"`
	Name      string         `json:"name"`
	Username  string         `json:"username" gorm:"index:idx_username,unique"`
	Password  string         `json:"-"`
	Token     string         `json:"token,omitempty" gorm:"-"`
}
