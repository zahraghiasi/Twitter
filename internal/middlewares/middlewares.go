package middlewares

import (
	"github.com/labstack/echo/v4"
	"github.com/labstack/echo/v4/middleware"
)

type MiddleWare struct {
	loggedIn bool
}

func New() *MiddleWare {
	middleWare := MiddleWare{}
	return &middleWare
}

func (m *MiddleWare) LoggedInCostumerMiddleware() echo.MiddlewareFunc {
	return middleware.KeyAuth(func(key string, c echo.Context) (bool, error) {
		cc := c.(*CustomContext)
		loggedIn, costumerObj := cc.IsLoggedInCostumer()
		return loggedIn && costumerObj.ID > 0, nil
	})
}

func (m *MiddleWare) LoggedInMiddleware() echo.MiddlewareFunc {
	return middleware.KeyAuth(func(key string, c echo.Context) (bool, error) {
		return len(key) > 30, nil
	})
}

func GetCORSMiddleWare() echo.MiddlewareFunc {
	return middleware.CORSWithConfig(middleware.CORSConfig{
		AllowOrigins: []string{"*"},
		AllowHeaders: []string{echo.HeaderOrigin, echo.HeaderContentType, echo.HeaderAccept, echo.HeaderAuthorization, "range", "Range"},
	})
}
