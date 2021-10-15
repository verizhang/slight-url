package user

import "gorm.io/gorm"

type User struct {
	gorm.Model
	Name     string
	Username string `gorm:"index:idx_username,unique"`
	Password string
}
