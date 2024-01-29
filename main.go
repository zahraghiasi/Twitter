package main

import (
	"context"
	"fmt"
	"github.com/ghiac/go-commons/log"
	"github.com/ghiac/go-commons/signal"
	"github.com/ghiac/social/internal/config"
	http_endpoint "github.com/ghiac/social/internal/http-endpoint"
	"github.com/ghiac/social/internal/mysql"
	"github.com/sirupsen/logrus"
	"go.mongodb.org/mongo-driver/mongo"
	"go.mongodb.org/mongo-driver/mongo/options"
	"time"
)

func initialize(sigKill chan bool) {
	fmt.Println("time:", time.Now())

	config.Initialize("")
	log.Initialize(config.Conf.Log.Level)
	if config.Conf.Core.Mode == "release" {
		log.Logger.SetFormatter(&logrus.JSONFormatter{})
	} else {
		config.Initialize("./conf/config.yaml")
	}
	database := mysql.New(mysql.Option{
		Host: config.Conf.Mysql.Host,
		Port: config.Conf.Mysql.Port,
		User: config.Conf.Mysql.User,
		Pass: config.Conf.Mysql.Pass,
		Db:   config.Conf.Mysql.DB,
	})

	client, err := mongo.NewClient(options.Client().ApplyURI("mongodb://"+config.Conf.Mongo.Address), &options.ClientOptions{Auth: &options.Credential{
		Username: config.Conf.Mongo.Username,
		Password: config.Conf.Mongo.Password,
	}})
	if err != nil {
		panic(err)
	}
	ctx, cancel := context.WithTimeout(context.Background(), 60*time.Second)
	defer cancel()
	err = client.Connect(ctx)
	if err != nil {
		panic("MonogoDb connection error:" + err.Error())
	}
	mongoDatabase := client.Database("Social")

	if err != nil {
		panic(err)
	}

	httpEndpoint := http_endpoint.New(sigKill, database, mongoDatabase)
	httpEndpoint.Start()
}

func main() {
	killSession := make(chan bool)
	initialize(killSession)
	signal.Signal.Wait()
	log.Logger.Info("Start To Kill")
	killSession <- true
	<-killSession
	log.Logger.Info("Dead!")
}
