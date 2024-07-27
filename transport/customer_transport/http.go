package customer_transport

import (
	"github.com/golang-jwt/jwt/v5"
	"github.com/labstack/echo/v4"
	"net/http"
	"order-service/core/entity"
	"order-service/core/service"
	"order-service/shared/middleware"
	"order-service/shared/validator"
)

type HttpHandlerCustomer struct {
	customerService service.CustomerService
	validator       *validator.Validator
}
type Opts struct {
	CustomerService service.CustomerService
	Validator       *validator.Validator
}

func (httpHandler *HttpHandlerCustomer) NewCustomerHttpHandler(e *echo.Group, o Opts) {
	httpHandler.customerService = o.CustomerService
	httpHandler.validator = o.Validator

	rest := e.Group("/customer")
	rest.Use(middleware.UseSecureHeaders())
	rest.GET("", httpHandler.GetAllCustomer)
	rest.GET("/:uid", httpHandler.GetCustomer)
	rest.PATCH("/:uid", httpHandler.UpdateCustomer)
	rest.DELETE("/:uid", httpHandler.DeleteCustomer)
	rest.PUT("", httpHandler.CreateCustomer)

}

func (httpHandler *HttpHandlerCustomer) GetAllCustomer(c echo.Context) error {
	var err error

	request := new(entity.RequestGetList)
	if err := c.Bind(request); err != nil {
		return c.JSON(http.StatusOK, middleware.MakeFailResponse(err.Error()))
	}

	if err := httpHandler.validator.ValidateStruct(c.Request().Context(), request); err != nil {
		return c.JSON(http.StatusOK, middleware.MakeFailResponse(err.Error()))
	}
	result, err := httpHandler.customerService.GetAllCustomer(c.Request().Context(), request)

	if err != nil {
		return c.JSON(http.StatusOK, middleware.MakeErrorResponse(err.Error()))
	}

	return c.JSON(http.StatusOK, middleware.MakeSuccessResponse(result))
}

func (httpHandler *HttpHandlerCustomer) GetCustomer(c echo.Context) error {
	var err error

	uid := c.Param("uid")
	if uid == "" {
		return c.JSON(http.StatusOK, middleware.MakeFailResponse("empty uid"))
	}

	result, err := httpHandler.customerService.GetCustomer(c.Request().Context(), &uid)

	if err != nil {
		return c.JSON(http.StatusOK, middleware.MakeErrorResponse(err.Error()))
	}

	return c.JSON(http.StatusOK, middleware.MakeSuccessResponse(result))
}

func (httpHandler *HttpHandlerCustomer) CreateCustomer(c echo.Context) error {

	user := c.Get("user").(*jwt.Token)
	claims := user.Claims.(*entity.JwtClaims)

	request := new(entity.RequestCreateCustomer)
	if err := c.Bind(request); err != nil {
		return c.JSON(http.StatusOK, middleware.MakeFailResponse(err.Error()))
	}

	if err := httpHandler.validator.ValidateStruct(c.Request().Context(), request); err != nil {
		return c.JSON(http.StatusOK, middleware.MakeFailResponse(err.Error()))
	}

	request.UserId = claims.UserId

	err := httpHandler.customerService.CreateCustomer(c.Request().Context(), request)
	if err != nil {
		return c.JSON(http.StatusOK, middleware.MakeErrorResponse(err.Error()))
	}

	return c.JSON(http.StatusOK, middleware.MakeSuccessResponse(nil))
}

func (httpHandler *HttpHandlerCustomer) UpdateCustomer(c echo.Context) error {
	var err error
	var customer *entity.Customer

	uid := c.Param("uid")
	if uid == "" {
		return c.JSON(http.StatusOK, middleware.MakeFailResponse("empty uid"))
	}
	request := new(entity.RequestUpdateCustomer)
	if err := c.Bind(request); err != nil {
		return c.JSON(http.StatusOK, middleware.MakeFailResponse(err.Error()))
	}

	if err := httpHandler.validator.ValidateStruct(c.Request().Context(), request); err != nil {
		return c.JSON(http.StatusOK, middleware.MakeFailResponse(err.Error()))
	}
	err = httpHandler.customerService.UpdateCustomer(c.Request().Context(), &uid, request)

	if err != nil {
		return c.JSON(http.StatusOK, middleware.MakeErrorResponse(err.Error()))
	}

	return c.JSON(http.StatusOK, middleware.MakeSuccessResponse(customer))
}
func (httpHandler *HttpHandlerCustomer) DeleteCustomer(c echo.Context) error {
	var err error

	uid := c.Param("uid")
	if uid == "" {
		return c.JSON(http.StatusOK, middleware.MakeFailResponse("empty uid"))
	}

	err = httpHandler.customerService.DeleteCustomer(c.Request().Context(), &uid)

	if err != nil {
		return c.JSON(http.StatusOK, middleware.MakeErrorResponse(err.Error()))
	}

	return c.JSON(http.StatusOK, middleware.MakeSuccessResponse(nil))
}
