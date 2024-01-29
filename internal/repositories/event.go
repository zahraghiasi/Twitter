package repositories

import "github.com/jinzhu/gorm"

type Event struct {
	gorm.Model
	UserID    int
	EventType EventType
	Message   string
}

type EventType int

const (
	NotDefinedEvent EventType = iota
	Followed
)
