package mysql

import (
	"github.com/ghiac/social/internal/repositories"
	"github.com/jinzhu/gorm"
)

type EventMysqlRepo struct {
	db *gorm.DB
}

func NewEventMysqlRepo(db *gorm.DB) *EventMysqlRepo {
	return &EventMysqlRepo{db}
}

func (u *EventMysqlRepo) Add(event *repositories.Event) error {
	if err := u.db.Create(event).Error; err != nil {
		return err
	}
	return nil
}

func (u *EventMysqlRepo) GetAll() []repositories.Event {
	events := []repositories.Event{}
	u.db.Find(events)
	return events
}

func (u *EventMysqlRepo) Get(i uint) (repositories.Event, bool) {
	event := repositories.Event{}
	notFound := u.db.Where("id = ? ", i).First(&event).RecordNotFound()
	return event, !notFound
}
