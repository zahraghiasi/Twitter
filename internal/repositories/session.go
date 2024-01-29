package repositories

import "github.com/jinzhu/gorm"

type Session struct {
	gorm.Model
	UserID uint
	Token  string
}
