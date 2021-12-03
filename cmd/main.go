package main

import (
	"context"
	"github.com/1makarov/go-logger-grpc-example/internal/config"
	"github.com/1makarov/go-logger-grpc-example/internal/mongodb"
	"github.com/1makarov/go-logger-grpc-example/internal/repository"
	"github.com/1makarov/go-logger-grpc-example/internal/server"
	"github.com/1makarov/go-logger-grpc-example/internal/services"
	"github.com/1makarov/go-logger-grpc-example/pkg/signaler"
	"github.com/sirupsen/logrus"
)

func init() {
	logrus.SetFormatter(new(logrus.JSONFormatter))
}

func main() {
	cfg := config.Init()
	ctx := context.Background()

	mongoClient, err := mongodb.Open(context.Background(), cfg.DB)
	if err != nil {
		logrus.Errorln(err)
		return
	}
	defer func() {
		if err = mongoClient.Disconnect(ctx); err != nil {
			logrus.Errorln(err)
		}
	}()

	db := mongoClient.Database(cfg.DB.Name)

	repo := repository.New(db)
	service := services.New(repo)

	srv := server.New(service.Logger)
	defer srv.Stop()

	go func() {
		if err = srv.Start(cfg.Port); err != nil {
			logrus.Errorln(err)
		}
	}()

	logrus.Infoln("started")

	signaler.Wait()
}
