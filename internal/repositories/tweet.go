package repositories

import "github.com/jinzhu/gorm"

type Tweet struct {
	gorm.Model
	UserID    uint
	Message   []byte
	ReTweetID *uint
}

func (t Tweet) ToResponse() *TweetResponse {
	return &TweetResponse{
		User:      User{},
		Message:   string(t.Message),
		ReTweetId: t.ReTweetID,
	}
}

type TweetResponse struct {
	User      User
	Message   string
	ReTweetId *uint
}
