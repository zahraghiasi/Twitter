package repositories

import "github.com/jinzhu/gorm"

type Hashtag struct {
	gorm.Model
	UserID  int
	TweetID int
	Hashtag string
}
