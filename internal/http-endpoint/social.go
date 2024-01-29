package http_endpoint

import (
	"github.com/ghiac/social/internal/middlewares"
	"github.com/ghiac/social/internal/repositories"
	"github.com/ghiac/social/internal/requests"
	"github.com/ghiac/social/internal/response"
	"github.com/labstack/echo/v4"
)

func (h *HttpServer) follow(context echo.Context) error {
	cc := context.(*middlewares.CustomContext)
	u := new(requests.RelationRequest)
	if err := context.Bind(u); err != nil {
		return response.GetBadRequest(context, err.Error())
	}
	if !cc.LoggedIn {
		return response.GetUnAuthorized(context)
	}
	if cc.UserObj.ID == u.UserID {
		return response.GetTrueSuccess(context)
	}
	err := h.handler.FollowUser(cc.UserObj.ID, u.UserID)
	if err != nil {
		return response.FromServerError(cc, *err)
	}
	return response.GetTrueSuccess(cc)
}

func (h *HttpServer) unfollow(context echo.Context) error {
	cc := context.(*middlewares.CustomContext)
	u := new(requests.RelationRequest)
	if err := context.Bind(u); err != nil {
		return response.GetBadRequest(context, err.Error())
	}
	if !cc.LoggedIn {
		return response.GetUnAuthorized(context)
	}
	if cc.UserObj.ID == u.UserID {
		return response.GetTrueSuccess(context)
	}
	err := h.handler.FollowUser(cc.UserObj.ID, u.UserID)
	if err != nil {
		return response.FromServerError(cc, *err)
	}
	return response.GetTrueSuccess(cc)
}

func (h *HttpServer) tweets(context echo.Context) error {
	tweets := h.handler.GetTweets()
	return response.GetResult(context, tweets)
}

type Query struct {
	ID uint `uri:"id" binding:"id"`
}

func (h *HttpServer) createTweet(context echo.Context) error {
	cc := context.(*middlewares.CustomContext)
	u := new(repositories.Tweet)
	if err := context.Bind(u); err != nil {
		return response.GetBadRequest(context, err.Error())
	}
	if !cc.LoggedIn {
		return response.GetUnAuthorized(context)
	}
	err := h.handler.CreateTweet(cc.UserObj.ID, string(u.Message), u.ReTweetID)
	if err != nil {
		return response.FromServerError(cc, *err)
	}
	return response.GetTrueSuccess(cc)
}

func (h *HttpServer) getTweet(context echo.Context) error {
	cc := context.(*middlewares.CustomContext)
	var query Query
	err := context.Bind(&query)
	if err != nil {
		return response.GetBadRequest(context, err.Error())
	}
	tweet, sErr := h.handler.GetTweet(query.ID)
	if sErr != nil {
		return response.FromServerError(cc, *sErr)
	}
	return response.GetResult(cc, tweet)
}

func (h *HttpServer) editTweet(context echo.Context) error {
	cc := context.(*middlewares.CustomContext)
	var query Query
	err := context.Bind(&query)
	if err != nil {
		return response.GetBadRequest(context, err.Error())
	}
	return h.NotImplement(cc)
}

func (h *HttpServer) deleteTweet(context echo.Context) error {
	cc := context.(*middlewares.CustomContext)
	var query Query
	err := context.Bind(&query)
	if err != nil {
		return response.GetBadRequest(context, err.Error())
	}
	return h.NotImplement(cc)
}

func (h *HttpServer) likeReact(context echo.Context) error {
	cc := context.(*middlewares.CustomContext)
	var query Query
	err := context.Bind(&query)
	if err != nil {
		return response.GetBadRequest(context, err.Error())
	}
	return h.NotImplement(cc)
}

func (h *HttpServer) removeReact(context echo.Context) error {
	cc := context.(*middlewares.CustomContext)
	var query Query
	err := context.Bind(&query)
	if err != nil {
		return response.GetBadRequest(context, err.Error())
	}
	return h.NotImplement(cc)
}
