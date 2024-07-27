package repositories

import (
	"context"
	"order-service/core/entity"
	"order-service/data_source/mysql_datasource"
)

type UserRepository interface {
	AutoMigration(ctx context.Context) error
	UpdateUser(ctx context.Context, user *entity.User, updateValue map[string]interface{}) error
	CreateUser(ctx context.Context, user *entity.User) error
	DeleteUser(ctx context.Context, user *entity.User, condition []map[string]interface{}) error
	GetUserByUid(ctx context.Context, uid *string) (user *entity.User, err error)
	GetUserByEmail(ctx context.Context, email *string) (user *entity.User, err error)
	GetAllUser(ctx context.Context, queryOption mysql_datasource.QueryOption) (res *entity.ResponseGetAllUser, err error)
}

type CustomerRepository interface {
	UpdateCustomer(ctx context.Context, customer *entity.Customer, updateValue map[string]interface{}) error
	CreateCustomer(ctx context.Context, customer *entity.Customer) error
	DeleteCustomer(ctx context.Context, customer *entity.Customer, condition []map[string]interface{}) error
	GetCustomerByUid(ctx context.Context, uid *string) (order *entity.Customer, err error)
	GetCustomerByEmail(ctx context.Context, email *string) (order *entity.Customer, err error)
	GetAllCustomer(ctx context.Context, queryOption mysql_datasource.QueryOption) (res *entity.ResponseGetAllCustomer, err error)
}

type OrderRepository interface {
	UpdateOrder(ctx context.Context, order *entity.Order, updateValue map[string]interface{}) error
	CreateOrder(ctx context.Context, order *entity.Order) error
	DeleteOrder(ctx context.Context, order *entity.Order, condition []map[string]interface{}) error
	GetOrderByUid(ctx context.Context, uid *string) (order *entity.Order, err error)
	GetAllOrder(ctx context.Context, queryOption mysql_datasource.QueryOption) (res *entity.ResponseGetAllOrder, err error)
}
