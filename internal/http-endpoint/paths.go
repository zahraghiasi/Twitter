package http_endpoint

import (
	"fmt"
	"github.com/ghiac/social/internal/config"
	"github.com/ghiac/social/internal/middlewares"
	"github.com/ghiac/social/internal/mysql"
	"github.com/labstack/echo/v4"
	"github.com/labstack/echo/v4/middleware"
	"github.com/patrickmn/go-cache"
	"strconv"
	"time"
)

func (h *HttpServer) AddPaths(userRepo *mysql.UserMysqlRepo, sessionRepo *mysql.SessionMysqlRepo) {
	h.addStaticRoutes()
	tokenCache := cache.New(time.Hour, time.Hour*24)
	h.echo.Use(func(next echo.HandlerFunc) echo.HandlerFunc {
		return func(c echo.Context) error {
			cc := &middlewares.CustomContext{
				Context:     c,
				Cache:       tokenCache,
				UserRepo:    userRepo,
				SessionRepo: sessionRepo,
			}
			cc.IsLoggedInCostumer()
			return next(cc)
		}
	})

	h.echo.POST("/Login", h.login)
	h.echo.POST("/Register", h.register)
	h.echo.POST("/UploadProfile", h.uploadProfile, h.middleWare.LoggedInMiddleware())
	h.echo.POST("/UpdateProfile", h.updateProfile, h.middleWare.LoggedInMiddleware())
	h.echo.POST("/User", h.user, h.middleWare.LoggedInCostumerMiddleware())
	h.echo.GET("/User", h.user, h.middleWare.LoggedInCostumerMiddleware())
	h.echo.POST("/Tweet", h.tweet, h.middleWare.LoggedInMiddleware())
	h.echo.GET("/Tweet/:id", h.tweet, h.middleWare.LoggedInMiddleware())
	h.echo.PUT("/Tweet/:id", h.editTweet, h.middleWare.LoggedInMiddleware())
	h.echo.DELETE("/Tweet/:id", h.deleteTweet, h.middleWare.LoggedInMiddleware())
	h.echo.GET("/React/Like/:id", h.likeReact, h.middleWare.LoggedInMiddleware())
	h.echo.GET("/React/Remove/:id", h.removeReact, h.middleWare.LoggedInMiddleware())
	h.echo.GET("/Search", h.search) // By hashtag and username
}

func (h *HttpServer) BindHttp() {
	h.echo.Debug = false
	h.echo.HideBanner = true
	h.echo.HidePort = true
	fmt.Println("Started ", "0.0.0.0", " on ", config.Conf.Core.Port)
	h.echo.Logger.Fatal(h.echo.Start(":" + strconv.Itoa(config.Conf.Core.Port)))
}

func (h *HttpServer) addStaticRoutes() {
	h.echo.Use(middlewares.GetCORSMiddleWare())
	h.echo.Use(middleware.BodyLimit("80M"))

	//h.echo.GET("/ostatic/:filename/:quality/:width/:height", h.handler.GeneralHandler.OptImgStatic)
}
