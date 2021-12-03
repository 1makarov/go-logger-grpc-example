package services

import (
	"context"
	"github.com/1makarov/go-logger-grpc-example/internal/proto"
	"github.com/1makarov/go-logger-grpc-example/internal/repository"
)

type Logger interface {
	Add(ctx context.Context, v *proto.ActionRequest) (*proto.Empty, error)
}

type Service struct {
	Logger Logger
}

func New(repo *repository.Repository) *Service {
	return &Service{
		Logger: NewLoggerService(repo.Logger),
	}
}
