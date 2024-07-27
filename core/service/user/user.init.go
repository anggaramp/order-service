package user

import (
	"go.uber.org/zap"
	"order-service/core/repositories"
	"order-service/core/service"
)

type module struct {
	repo   repositories.UserRepository
	logger *zap.Logger
}

type Opts struct {
	Repo   repositories.UserRepository
	Logger *zap.Logger
}

func New(o Opts) service.UserService {
	return &module{
		repo:   o.Repo,
		logger: o.Logger,
	}
}
