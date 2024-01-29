package handler

import (
	"github.com/ghiac/social/internal/repositories"
	"github.com/ghiac/social/pkg/access"
	"github.com/ghiac/social/pkg/text"
)

func (h *Handler) CreateTweet(userID uint, message string, reTweetID *uint) *access.ServerError {
	if reTweetID == nil && len(message) < 6 {
		return access.GetError(text.NotEnoughChar)
	}

	if reTweetID != nil {
		tweetObj, found := h.TweetRepo.Get(*reTweetID)
		if !found {
			return access.GetError(text.NotFound)
		}
		retweetID := tweetObj.ID
		if tweetObj.ReTweetID != nil {
			retweetID = *tweetObj.ReTweetID
		}
		err := h.TweetRepo.Add(&repositories.Tweet{
			UserID:    userID,
			ReTweetID: &retweetID,
		})
		if err != nil {
			return access.GetInternalError("Create tweet", err)
		}
		return nil
	}
	err := h.TweetRepo.Add(&repositories.Tweet{
		UserID:  userID,
		Message: []byte(message),
	})
	if err != nil {
		return access.GetInternalError("Create tweet", err)
	}
	return nil
}

func (h *Handler) GetTweet(id uint) (*repositories.TweetResponse, *access.ServerError) {
	tweetObj, found := h.TweetRepo.Get(id)
	if !found {
		return nil, access.GetError(text.NotFound)
	}
	return tweetObj.ToResponse(), nil
}

func (h *Handler) GetTweets() []repositories.TweetResponse {
	responses := []repositories.TweetResponse{}
	tweets := h.TweetRepo.GetAll()
	for _, tweet := range tweets {
		responses = append(responses, *tweet.ToResponse())
	}
	return responses
}

func (h *Handler) EditTweet(userID, id uint, message string) *access.ServerError {
	tweet, found := h.TweetRepo.Get(id)
	if !found {
		return access.GetError(text.NotFound)
	}
	if tweet.UserID != userID {
		return access.GetError(text.Forbidden)
	}
	err := h.TweetRepo.EditTweet(tweet.ID, []byte(message))
	if err != nil {
		return access.GetInternalError("Edit Tweet", err)
	}
	return nil
}

func (h *Handler) DeleteTweet(userID, id uint) *access.ServerError {
	tweet, found := h.TweetRepo.Get(id)
	if !found {
		return access.GetError(text.NotFound)
	}
	if tweet.UserID != userID {
		return access.GetError(text.Forbidden)
	}
	err := h.TweetRepo.DeleteTweet(tweet.ID)
	if err != nil {
		return access.GetInternalError("Delete Tweet", err)
	}
	return nil
}
