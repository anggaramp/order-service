package customer

import (
	"go.uber.org/zap"
	"order-service/core/repositories"
	"order-service/core/service"
)

type module struct {
	repo   repositories.CustomerRepository
	logger *zap.Logger
}

type Opts struct {
	Repo   repositories.CustomerRepository
	Logger *zap.Logger
}

func New(o Opts) service.CustomerService {
	return &module{
		repo:   o.Repo,
		logger: o.Logger,
	}
}
