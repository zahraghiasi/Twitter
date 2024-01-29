package repositories

import "github.com/jinzhu/gorm"

type Tweet struct {
	gorm.Model
	UserID    int
	Message   []byte
	ReTweetID *int
}
