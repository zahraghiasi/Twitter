package http_endpoint

import (
	"fmt"
	"github.com/ghiac/social/internal/middlewares"
	"github.com/ghiac/social/internal/requests"
	"github.com/ghiac/social/internal/response"
	"github.com/labstack/echo/v4"
)

func (h *HttpServer) login(context echo.Context) error {
	cc := context.(*middlewares.CustomContext)
	u := new(requests.LoginRequest)
	if err := context.Bind(u); err != nil {
		return response.GetBadRequest(context, err.Error())
	}
	fmt.Println(u)
	resp, err := h.handler.Login(*u)
	if err != nil {
		return response.FromServerError(cc, *err)
	}
	return response.GetResult(cc, resp)
}

func (h *HttpServer) register(context echo.Context) error {
	cc := context.(*middlewares.CustomContext)
	u := new(requests.LoginRequest)
	if err := context.Bind(u); err != nil {
		return response.GetBadRequest(context, err.Error())
	}
	fmt.Println(u)

	resp, err := h.handler.Register(*u)
	if err != nil {
		return response.FromServerError(cc, *err)
	}
	return response.GetResult(cc, resp)
}

func (h *HttpServer) user(context echo.Context) error {
	cc := context.(*middlewares.CustomContext)
	u := new(requests.LoginRequest)
	if err := context.Bind(u); err != nil {
		return response.GetBadRequest(context, err.Error())
	}
	fmt.Println(u)

	resp, err := h.handler.UserInfo(int(cc.UserObj.ID))
	if err != nil {
		return response.FromServerError(cc, *err)
	}
	return response.GetResult(cc, resp)
}

func (h *HttpServer) updateProfile(c echo.Context) error {
	return h.NotImplement(c)
}

func (h *HttpServer) uploadPicture(c echo.Context) error {
	return h.NotImplement(c)
}
