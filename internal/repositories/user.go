package repositories

import (
	"github.com/jinzhu/gorm"
)

type User struct {
	gorm.Model
	Email    string `gorm:"index"`
	Name     string
	LastName string
	Username string
	Picture  string
	Password string `json:"-"`
}

func (User) TableName() string {
	return "user"
}
