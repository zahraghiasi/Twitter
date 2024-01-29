package mysql

import (
	"github.com/ghiac/social/internal/repositories"
	"github.com/jinzhu/gorm"
)

type HashtagMysqlRepo struct {
	db *gorm.DB
}

func NewHashtagMysqlRepo(db *gorm.DB) *HashtagMysqlRepo {
	return &HashtagMysqlRepo{db}
}

func (u *HashtagMysqlRepo) Add(reaction *repositories.Hashtag) error {
	if err := u.db.Create(reaction).Error; err != nil {
		return err
	}
	return nil
}

func (u *HashtagMysqlRepo) GetAll() []repositories.Hashtag {
	hashtags := []repositories.Hashtag{}
	u.db.Find(hashtags)
	return hashtags
}

func (u *HashtagMysqlRepo) Get(i uint) (repositories.Hashtag, bool) {
	hashtag := repositories.Hashtag{}
	notFound := u.db.Where("id = ? ", i).First(&hashtag).RecordNotFound()
	return hashtag, !notFound
}
