package middlewares

import (
	"github.com/ghiac/social/internal/models"
	"github.com/ghiac/social/internal/mysql"
	"github.com/ghiac/social/internal/repositories"
	"github.com/labstack/echo/v4"
	"github.com/patrickmn/go-cache"
	"net/http"
	"strings"
	"time"
)

type HFunc func(c *CustomContext) error

func NextContext(fun HFunc, loggedIn bool) echo.HandlerFunc {
	return func(context echo.Context) error {
		cc := context.(*CustomContext)
		if loggedIn {
			if !cc.LoggedIn {
				return context.JSON(http.StatusUnauthorized, models.GeneralResponse{
					Ok:          false,
					Description: "UnAuthorized",
				})
			}
		}
		return fun(cc)
	}
}

type CustomContext struct {
	echo.Context
	Cache       *cache.Cache
	UserObj     *repositories.User
	SessionObj  *repositories.Session
	LoggedIn    bool
	SessionRepo *mysql.SessionMysqlRepo
	UserRepo    *mysql.UserMysqlRepo
}

func (c *CustomContext) ClearCache() {
	bar := strings.Split(c.Request().Header.Get("Authorization"), " ")
	if len(bar) == 2 {
		token := bar[1]
		c.ClearCacheByToken(token)
	}
}

func (c *CustomContext) ClearCacheByToken(token string) {
	sessionCacheKey := "CheckTokenResponse" + token
	c.Cache.Delete(sessionCacheKey)
}

func (c *CustomContext) IsLoggedInCostumer() (bool, repositories.User) {
	bar := strings.Split(c.Request().Header.Get("Authorization"), " ")
	if len(bar) == 2 {
		token := bar[1]
		response := repositories.User{}
		sessionCacheKey := "CheckTokenResponse" + token
		respObj, found := c.Cache.Get(sessionCacheKey)
		if found {
			response = respObj.(repositories.User)
		} else {
			resp, found := c.SessionRepo.GetByToken(token)
			if !found {
				return false, repositories.User{}
			}
			user, found := c.UserRepo.Get(int(resp.UserID))
			if !found {
				return false, repositories.User{}
			}
			c.Cache.Set(sessionCacheKey, user, time.Second)
			response = user
		}
		c.LoggedIn = true
		c.UserObj = &response
		return true, response
	} else {
		return false, repositories.User{}
	}
}
