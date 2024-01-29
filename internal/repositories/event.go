package repositories

import "github.com/jinzhu/gorm"

type Event struct {
	gorm.Model
	UserID  int
	Message string
}
