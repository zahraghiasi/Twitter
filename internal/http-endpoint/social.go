package http_endpoint

import (
	"github.com/ghiac/social/internal/middlewares"
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

func (h *HttpServer) tweets(c echo.Context) error {
	return h.NotImplement(c)
}

func (h *HttpServer) tweet(c echo.Context) error {
	return h.NotImplement(c)
}

func (h *HttpServer) editTweet(c echo.Context) error {
	return h.NotImplement(c)
}

func (h *HttpServer) deleteTweet(c echo.Context) error {
	return h.NotImplement(c)
}

func (h *HttpServer) likeReact(c echo.Context) error {
	return h.NotImplement(c)
}

func (h *HttpServer) removeReact(c echo.Context) error {
	return h.NotImplement(c)
}
