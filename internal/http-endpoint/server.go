package http_endpoint

import (
	"github.com/ghiac/go-commons/log"
	"github.com/ghiac/social/internal/handler"
	"github.com/ghiac/social/internal/middlewares"
	"github.com/ghiac/social/internal/mysql"
	"github.com/ghiac/social/internal/repositories"
	"github.com/labstack/echo/v4"
	"go.mongodb.org/mongo-driver/mongo"
)

type HttpServer struct {
	echo       *echo.Echo
	handler    *handler.Handler
	database   *repositories.Database
	middleWare *middlewares.MiddleWare
}

func (h *HttpServer) startToKillListen(signalKill chan bool) {
	<-signalKill
	err := h.echo.Close()
	if err != nil {
		log.Logger.Error(err)
	}
	signalKill <- true
}

func (h *HttpServer) Start() {
	h.handler.Start()
}

func New(sigKill chan bool, database *repositories.Database, mnDb *mongo.Database) HttpServer {
	middleWare := middlewares.New()
	httpServer := HttpServer{
		echo:       echo.New(),
		handler:    handler.New(database),
		database:   database,
		middleWare: middleWare,
	}

	httpServer.AddPaths(mysql.NewUserMysqlRepo(database.Connection), mysql.NewSessionMysqlRepo(database.Connection))
	go httpServer.BindHttp()

	go httpServer.startToKillListen(sigKill)
	return httpServer
}
