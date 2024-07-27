package order_transport

import (
	"github.com/labstack/echo/v4"
	"net/http"
	"order-service/core/entity"
	"order-service/core/service"
	"order-service/shared/middleware"
	"order-service/shared/validator"
)

type HttpHandlerOrder struct {
	orderService service.OrderService
	validator    *validator.Validator
}
type Opts struct {
	OrderService service.OrderService
	Validator    *validator.Validator
}

func (httpHandler *HttpHandlerOrder) NewOrderHttpHandler(e *echo.Group, o Opts) {
	httpHandler.orderService = o.OrderService
	httpHandler.validator = o.Validator

	rest := e.Group("/order")
	rest.Use(middleware.UseSecureHeaders())
	rest.GET("", httpHandler.GetAllOrder)
	rest.GET("/:uid", httpHandler.GetOrder)
	rest.PATCH("/:uid", httpHandler.UpdateOrder)
	rest.DELETE("/:uid", httpHandler.DeleteOrder)
	rest.PUT("", httpHandler.CreateOrder)

}

func (httpHandler *HttpHandlerOrder) GetAllOrder(c echo.Context) error {
	var err error

	request := new(entity.RequestGetList)
	if err := c.Bind(request); err != nil {
		return c.JSON(http.StatusOK, middleware.MakeFailResponse(err.Error()))
	}

	if err := httpHandler.validator.ValidateStruct(c.Request().Context(), request); err != nil {
		return c.JSON(http.StatusOK, middleware.MakeFailResponse(err.Error()))
	}
	result, err := httpHandler.orderService.GetAllOrder(c.Request().Context(), request)

	if err != nil {
		return c.JSON(http.StatusOK, middleware.MakeErrorResponse(err.Error()))
	}

	return c.JSON(http.StatusOK, middleware.MakeSuccessResponse(result))
}

func (httpHandler *HttpHandlerOrder) GetOrder(c echo.Context) error {
	var err error

	uid := c.Param("uid")
	if uid == "" {
		return c.JSON(http.StatusOK, middleware.MakeFailResponse("empty uid"))
	}

	result, err := httpHandler.orderService.GetOrder(c.Request().Context(), &uid)

	if err != nil {
		return c.JSON(http.StatusOK, middleware.MakeErrorResponse(err.Error()))
	}

	return c.JSON(http.StatusOK, middleware.MakeSuccessResponse(result))
}

func (httpHandler *HttpHandlerOrder) CreateOrder(c echo.Context) error {

	request := new(entity.RequestCreateOrder)
	if err := c.Bind(request); err != nil {
		return c.JSON(http.StatusOK, middleware.MakeFailResponse(err.Error()))
	}

	if err := httpHandler.validator.ValidateStruct(c.Request().Context(), request); err != nil {
		return c.JSON(http.StatusOK, middleware.MakeFailResponse(err.Error()))
	}

	err := httpHandler.orderService.CreateOrder(c.Request().Context(), request)
	if err != nil {
		return c.JSON(http.StatusOK, middleware.MakeErrorResponse(err.Error()))
	}

	return c.JSON(http.StatusOK, middleware.MakeSuccessResponse(nil))
}

func (httpHandler *HttpHandlerOrder) UpdateOrder(c echo.Context) error {
	var err error
	var order *entity.Order

	uid := c.Param("uid")
	if uid == "" {
		return c.JSON(http.StatusOK, middleware.MakeFailResponse("empty uid"))
	}
	request := new(entity.RequestUpdateOrder)
	if err := c.Bind(request); err != nil {
		return c.JSON(http.StatusOK, middleware.MakeFailResponse(err.Error()))
	}

	if err := httpHandler.validator.ValidateStruct(c.Request().Context(), request); err != nil {
		return c.JSON(http.StatusOK, middleware.MakeFailResponse(err.Error()))
	}
	err = httpHandler.orderService.UpdateOrder(c.Request().Context(), &uid, request)

	if err != nil {
		return c.JSON(http.StatusOK, middleware.MakeErrorResponse(err.Error()))
	}

	return c.JSON(http.StatusOK, middleware.MakeSuccessResponse(order))
}
func (httpHandler *HttpHandlerOrder) DeleteOrder(c echo.Context) error {
	var err error

	uid := c.Param("uid")
	if uid == "" {
		return c.JSON(http.StatusOK, middleware.MakeFailResponse("empty uid"))
	}

	err = httpHandler.orderService.DeleteOrder(c.Request().Context(), &uid)

	if err != nil {
		return c.JSON(http.StatusOK, middleware.MakeErrorResponse(err.Error()))
	}

	return c.JSON(http.StatusOK, middleware.MakeSuccessResponse(nil))
}
