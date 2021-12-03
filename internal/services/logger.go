package services

import (
	"context"
	"github.com/1makarov/go-logger-grpc-example/internal/proto"
	"github.com/1makarov/go-logger-grpc-example/internal/repository"
	"github.com/1makarov/go-logger-grpc-example/internal/types"
)

type LoggerService struct {
	repo repository.Logger
}

func NewLoggerService(repo repository.Logger) *LoggerService {
	return &LoggerService{
		repo: repo,
	}
}

func (s *LoggerService) Add(ctx context.Context, v *proto.ActionRequest) (*proto.Empty, error) {
	item := types.LogItem{
		Entity:    v.Entity.String(),
		Action:    v.Action.String(),
		UserID:    int64(v.UserId),
		Timestamp: v.Timestamp.AsTime(),
	}

	return &proto.Empty{}, s.repo.Add(ctx, item)
}
