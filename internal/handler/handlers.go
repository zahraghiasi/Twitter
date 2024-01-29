package handler

import (
	"github.com/ghiac/social/internal/mysql"
	"github.com/ghiac/social/internal/repositories"
)

type Handler struct {
	UserRepo    *mysql.UserMysqlRepo
	SessionRepo *mysql.SessionMysqlRepo
	TweetRepo   *mysql.TweetMysqlRepo
}

func New(db *repositories.Database) *Handler {
	return &Handler{
		UserRepo:    mysql.NewUserMysqlRepo(db.Connection),
		SessionRepo: mysql.NewSessionMysqlRepo(db.Connection),
		TweetRepo:   mysql.NewTweetMysqlRepo(db.Connection),
	}
}

func (h *Handler) Start() {
}
