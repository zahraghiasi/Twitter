package mysql

import (
	"github.com/ghiac/social/internal/repositories"
	"github.com/jinzhu/gorm"
)

type ReactionMysqlRepo struct {
	db *gorm.DB
}

func NewReactionMysqlRepo(db *gorm.DB) *ReactionMysqlRepo {
	return &ReactionMysqlRepo{db}
}

func (u *ReactionMysqlRepo) Add(reaction *repositories.Reaction) error {
	if err := u.db.Create(reaction).Error; err != nil {
		return err
	}
	return nil
}

func (u *ReactionMysqlRepo) GetAll() []repositories.Reaction {
	reactions := []repositories.Reaction{}
	u.db.Find(reactions)
	return reactions
}

func (u *ReactionMysqlRepo) Get(i uint) (repositories.Reaction, bool) {
	reaction := repositories.Reaction{}
	notFound := u.db.Where("id = ? ", i).First(&reaction).RecordNotFound()
	return reaction, !notFound
}
