package transport

import (
	"github.com/labstack/echo/v4"
	"go.uber.org/zap"
	"gorm.io/gorm"
	customerRepository "order-service/core/repositories/customer"
	orderRepository "order-service/core/repositories/order"
	userRepository "order-service/core/repositories/user"
	customerService "order-service/core/service/customer"
	orderService "order-service/core/service/order"
	userService "order-service/core/service/user"
	"order-service/data_source/mysql_datasource"
	"order-service/shared/validator"
	"order-service/transport/customer_transport"
	"order-service/transport/order_transport"
	"order-service/transport/user_transport"
)

type Transport struct {
	user_transport.HttpHandlerUser
	customer_transport.HttpHandlerCustomer
	order_transport.HttpHandlerOrder
}

func Setup(e *echo.Group, client *gorm.DB, logger *zap.Logger) (transport *Transport) {
	transport = &Transport{}
	//datasource
	mysqlDatasource := mysql_datasource.NewMysqlDatasource(client)

	userRepo := userRepository.New(userRepository.Opts{MysqlDatasource: mysqlDatasource})
	customerRepo := customerRepository.New(customerRepository.Opts{MysqlDatasource: mysqlDatasource})
	orderRepo := orderRepository.New(orderRepository.Opts{MysqlDatasource: mysqlDatasource})

	validate := validator.New()

	userSev := userService.New(userService.Opts{Repo: userRepo, Logger: logger})
	customerSev := customerService.New(customerService.Opts{Repo: customerRepo, Logger: logger})
	orderSev := orderService.New(orderService.Opts{Repo: orderRepo, CustomerRepo: customerRepo, Logger: logger})

	//setup transport

	transport.NewUserHttpHandler(e, user_transport.Opts{
		UserService: userSev,
		Validator:   validate,
	})

	transport.NewCustomerHttpHandler(e, customer_transport.Opts{
		CustomerService: customerSev,
		Validator:       validate,
	})

	transport.NewOrderHttpHandler(e, order_transport.Opts{
		OrderService: orderSev,
		Validator:    validate,
	})

	return
}
