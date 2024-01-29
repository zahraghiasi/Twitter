package mysql

import (
	"github.com/ghiac/social/internal/repositories"
	"github.com/jinzhu/gorm"
)

type RelationMysqlRepo struct {
	db *gorm.DB
}

func NewRelationMysqlRepo(db *gorm.DB) *RelationMysqlRepo {
	return &RelationMysqlRepo{db}
}

func (u *RelationMysqlRepo) Add(relation *repositories.Relation) error {
	if err := u.db.Create(relation).Error; err != nil {
		return err
	}
	return nil
}

func (u *RelationMysqlRepo) GetAll() []repositories.Relation {
	relations := []repositories.Relation{}
	u.db.Find(relations)
	return relations
}

func (u *RelationMysqlRepo) Get(i uint) (repositories.Relation, bool) {
	relation := repositories.Relation{}
	notFound := u.db.Where("id = ? ", i).First(&relation).RecordNotFound()
	return relation, !notFound
}
