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
	return nil
}

func (h *Handler) AddPost(cc *middlewares.CustomContext, costumerID uint64) *access.ServerError {
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
	return response.GetResult(c, true)
}
