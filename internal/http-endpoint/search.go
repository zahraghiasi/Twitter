package http_endpoint

import (
	"github.com/ghiac/social/internal/response"
	"github.com/labstack/echo/v4"
)

func (h *HttpServer) search(c echo.Context) error {

	return response.GetTrueSuccess(c)
}
