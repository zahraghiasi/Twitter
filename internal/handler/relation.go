package handler

import (
	"github.com/ghiac/social/internal/middlewares"
	"github.com/ghiac/social/internal/response"
	"github.com/ghiac/social/pkg/access"
	"github.com/labstack/echo/v4"
)

func (h *Handler) UnFollowUser(c echo.Context) error {
	//if err := h.repos.FollowCostumerAndSequenceScogramRepo.Delete(costumer.ID, u.UserID); err != nil {
	//	return response.GetTryAgain(c, err)
	//}
	return response.GetTrueSuccess(c)
}

func (h *Handler) FollowUser(follower, following uint) *access.ServerError {
	// TODO increase speed
	//_, exist := h.repos.FollowCostumerAndSequenceScogramRepo.IsExist(follower, following)
	//if exist {
	//	return nil
	//}
	//_, costumerFollowedFound := h.jhonClient.GetUser(u.UserID)
	//if costumerFollowedFound != nil {
	//	return access.GetError(text.UserNotFound)
	//}
	//if !exist {
	//	h.repos.FollowCostumerAndSequenceScogramRepo.Add(scogram.FollowCostumerAndSequence{
	//		FollowerCostumerID: cc.UserObj.ID,
	//		FollowedCostumerID: u.UserID,
	//		LastSeq:            0,
	//	})
	//} else {
	//	h.repos.FollowCostumerAndSequenceScogramRepo.Add(scogram.FollowCostumerAndSequence{
	//		FollowerCostumerID: cc.UserObj.ID,
	//		FollowedCostumerID: u.UserID,
	//		LastSeq:            0,
	//	})
	//}
	return nil
}

func (h *Handler) AddPost(cc *middlewares.CustomContext, costumerID uint64) *access.ServerError {
	//kind := new(requests.PostRequest)
	//post := repositories.Tweet{
	//	UserID:    0,
	//	Message:   nil,
	//	ReTweetID: nil,
	//}
	//err := h.repos.PostsScogramRepo.Add(&post)
	//if err != nil {
	//	return err
	//}
	return nil
}

func (h *Handler) RemovePost(cc *middlewares.CustomContext) error {
	//post, found := h.repos.BlogPostRepo.Get(u.ID)
	//if !found || (post.UserID != cc.UserObj.ID && !hasAdminPermission) {
	//	post, found := h.repos.PostsScogramRepo.Get(u.ID)
	//	if !found || (post.UserID != cc.UserObj.ID && !hasAdminPermission) {
	//		return response.GetErrorNotFound(cc, ".پست پیدا نشد")
	//	}
	//	err := h.repos.PostsScogramRepo.Delete(post.ID)
	//	if err != nil {
	//		return response.GetTryAgain(cc, err)
	//	}
	//	return response.GetTrueSuccess(cc)
	//}
	//// TODO Use Peer
	//err := h.repos.BlogPostRepo.Delete(post.ID)
	////err := h.repos.PostsScogramRepo.Delete(post.ID)
	//if err != nil {
	//	return response.FromServerError(cc, *err)
	//}
	return response.GetTrueSuccess(cc)
}

// feed: true
// username
// hashtag
// most views
// most likes
func (h *Handler) GetPosts(c echo.Context) error {
	//cc := c.(*middlewares.CustomContext)
	//loggedIn, costumerObj := cc.IsLoggedInUser()
	//if !loggedIn {
	//	return response.GetUnAuthorized(c)
	//}

	//var postResponse []scogram.ScogramPost

	return response.GetResult(c, true)
}
