package main

import (
	"context"
	"fmt"
	"github.com/1makarov/go-logger-rpc-example/config"
	"github.com/1makarov/go-logger-rpc-example/internal/db/mongo"
	"github.com/1makarov/go-logger-rpc-example/internal/pkg/signaler"
	"github.com/1makarov/go-logger-rpc-example/internal/repository"
	"github.com/1makarov/go-logger-rpc-example/internal/server"
	"github.com/1makarov/go-logger-rpc-example/internal/service"
	"github.com/sirupsen/logrus"
)

func init() {
	logrus.SetFormatter(new(logrus.JSONFormatter))
}

func main() {
	cfg := config.Init()

	db, err := mongo.Open(context.Background(), cfg.DB)
	if err != nil {
		logrus.Fatalln(err)
	}

	repo := repository.New(db.Connect())
	services := service.New(repo)
	srv := server.New(services)

	go func() {
		if err = srv.Start(cfg.Port); err != nil {
			logrus.Errorln(err)
		}
	}()

	fmt.Println("logger started")

	signaler.Wait()

	srv.Stop()

	if err = db.Close(context.Background()); err != nil {
		logrus.Errorln(err)
	}
}
