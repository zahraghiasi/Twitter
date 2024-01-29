package repositories

import "github.com/jinzhu/gorm"

type Database struct {
	Connection *gorm.DB
}
