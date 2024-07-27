package order

import (
	"go.uber.org/zap"
	"order-service/core/repositories"
	"order-service/core/service"
)

type module struct {
	repo         repositories.OrderRepository
	customerRepo repositories.CustomerRepository
	logger       *zap.Logger
}

type Opts struct {
	Repo         repositories.OrderRepository
	CustomerRepo repositories.CustomerRepository
	Logger       *zap.Logger
}

func New(o Opts) service.OrderService {
	return &module{
		repo:         o.Repo,
		customerRepo: o.CustomerRepo,
		logger:       o.Logger,
	}
}
