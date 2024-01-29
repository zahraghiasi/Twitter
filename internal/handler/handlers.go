package handler

import (
	"github.com/ghiac/social/internal/mysql"
	"github.com/ghiac/social/internal/repositories"
)

type Handler struct {
	UserRepo     *mysql.UserMysqlRepo
	SessionRepo  *mysql.SessionMysqlRepo
	TweetRepo    *mysql.TweetMysqlRepo
	RelationRepo *mysql.RelationMysqlRepo
	ReactionRepo *mysql.ReactionMysqlRepo
	HashtagRepo  *mysql.HashtagMysqlRepo
	EventRepo    *mysql.EventMysqlRepo
}

func New(db *repositories.Database) *Handler {
	return &Handler{
		UserRepo:     mysql.NewUserMysqlRepo(db.Connection),
		SessionRepo:  mysql.NewSessionMysqlRepo(db.Connection),
		TweetRepo:    mysql.NewTweetMysqlRepo(db.Connection),
		RelationRepo: mysql.NewRelationMysqlRepo(db.Connection),
		ReactionRepo: mysql.NewReactionMysqlRepo(db.Connection),
		HashtagRepo:  mysql.NewHashtagMysqlRepo(db.Connection),
		EventRepo:    mysql.NewEventMysqlRepo(db.Connection),
	}
}

func (h *Handler) Start() {
}
