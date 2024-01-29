package repositories

import (
	"github.com/jinzhu/gorm"
	"time"
)

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

type TweetRequest struct {
	ID        uint      `param:"ID" query:"ID" form:"ID" json:"ID" xml:"ID"`
	ReTweetID *uint     `param:"ReTweetID" query:"ReTweetID" form:"ReTweetID" json:"ReTweetID" xml:"ReTweetID"`
	Message   string    `param:"Message" query:"Message" form:"Message" json:"Message" xml:"Message"`
	CreatedAt time.Time `param:"CreatedAt" query:"CreatedAt" form:"CreatedAt" json:"CreatedAt" xml:"CreatedAt"`
}
