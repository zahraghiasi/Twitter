package http_endpoint

import (
	"github.com/ghiac/social/internal/models"
	"github.com/labstack/echo/v4"
	"net/http"
)

func (h *HttpServer) NotImplement(c echo.Context) error {
	return c.JSON(http.StatusNotImplemented, models.GeneralResponse{
		Ok:          false,
		Description: "Not Implement",
	})
}

func (h *HttpServer) Forbidden(c echo.Context) error {
	return c.JSON(http.StatusForbidden, models.GeneralResponse{
		Ok:          false,
		Description: "Forbidden",
	})
}
