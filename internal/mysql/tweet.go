package mysql

import (
	"github.com/ghiac/social/internal/repositories"
	"github.com/jinzhu/gorm"
)

type TweetMysqlRepo struct {
	db *gorm.DB
}

func NewTweetMysqlRepo(db *gorm.DB) *TweetMysqlRepo {
	return &TweetMysqlRepo{db}
}

func (u *TweetMysqlRepo) Add(update *repositories.Tweet) error {
	if err := u.db.Create(update).Error; err != nil {
		return err
	}
	return nil
}

func (u *TweetMysqlRepo) GetAll() []repositories.Tweet {
	tweets := []repositories.Tweet{}
	u.db.Find(tweets)
	return tweets
}

func (u *TweetMysqlRepo) Get(i uint) (repositories.Tweet, bool) {
	session := repositories.Tweet{}
	notFound := u.db.Where("id = ? ", i).First(&session).RecordNotFound()
	return session, !notFound
}
