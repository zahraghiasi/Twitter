package middlewares

import (
	"github.com/ghiac/social/internal/mysql"
	"github.com/ghiac/social/internal/repositories"
	"github.com/labstack/echo/v4"
	"github.com/patrickmn/go-cache"
	"strings"
	"time"
)

type CustomContext struct {
	echo.Context
	Cache       *cache.Cache
	UserObj     *repositories.User
	SessionObj  *repositories.Session
	SessionRepo *mysql.SessionMysqlRepo
	UserRepo    *mysql.UserMysqlRepo
	LoggedIn    bool
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

func (c *CustomContext) IsLoggedInUser() (bool, repositories.User) {
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
