package repositories

import "github.com/jinzhu/gorm"

type Reaction struct {
	gorm.Model
	UserID    int
	TweetID   int
	ReactType ReactType
}

type ReactType int

const (
	Like RelationType = iota
	Dislike
)
