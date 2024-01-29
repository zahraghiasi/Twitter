package mysql

import (
	"github.com/ghiac/social/internal/repositories"
	"github.com/jinzhu/gorm"
)

type SessionMysqlRepo struct {
	db *gorm.DB
}

func NewSessionMysqlRepo(db *gorm.DB) *SessionMysqlRepo {
	return &SessionMysqlRepo{db}
}

func (u *SessionMysqlRepo) Add(update *repositories.Session) error {
	if err := u.db.Create(update).Error; err != nil {
		return err
	}
	return nil
}

func (u *SessionMysqlRepo) GetByToken(token string) (repositories.Session, bool) {
	session := repositories.Session{}
	notFound := u.db.Where("token = ? ", token).First(&session).RecordNotFound()
	return session, !notFound
}
