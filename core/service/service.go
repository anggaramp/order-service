package service

import (
	"context"
	"order-service/core/entity"
)

type UserService interface {
	GetUser(ctx context.Context, uid *string) (res *entity.ResponseGetUser, err error)
	GetAllUser(ctx context.Context, request *entity.RequestGetList) (res *entity.ResponseGetAllUser, err error)
	CreateUser(ctx context.Context, request *entity.RequestCreateUser) error
	UpdateUser(ctx context.Context, uid *string, request *entity.RequestUpdateUser) error
	LoginUser(ctx context.Context, request *entity.RequestLoginUser) (res *entity.ResponseLogin, err error)
	DeleteUser(ctx context.Context, uid *string) error
	Migration(ctx context.Context) error
}

type CustomerService interface {
	GetCustomer(ctx context.Context, uid *string) (res *entity.ResponseGetCustomer, err error)
	GetAllCustomer(ctx context.Context, request *entity.RequestGetList) (res *entity.ResponseGetAllCustomer, err error)
	CreateCustomer(ctx context.Context, request *entity.RequestCreateCustomer) error
	UpdateCustomer(ctx context.Context, uid *string, request *entity.RequestUpdateCustomer) error
	DeleteCustomer(ctx context.Context, uid *string) error
}

type OrderService interface {
	GetOrder(ctx context.Context, uid *string) (res *entity.ResponseGetOrder, err error)
	GetAllOrder(ctx context.Context, request *entity.RequestGetList) (res *entity.ResponseGetAllOrder, err error)
	CreateOrder(ctx context.Context, request *entity.RequestCreateOrder) error
	UpdateOrder(ctx context.Context, uid *string, request *entity.RequestUpdateOrder) error
	DeleteOrder(ctx context.Context, uid *string) error
}
