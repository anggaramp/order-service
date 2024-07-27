package user_transport

import (
	"github.com/labstack/echo/v4"
	"net/http"
	"order-service/core/entity"
	"order-service/core/service"
	"order-service/shared/middleware"
	"order-service/shared/validator"
)

type HttpHandlerUser struct {
	userService service.UserService
	validator   *validator.Validator
}
type Opts struct {
	UserService service.UserService
	Validator   *validator.Validator
}

func (httpHandler *HttpHandlerUser) NewUserHttpHandler(e *echo.Group, o Opts) {
	httpHandler.userService = o.UserService
	httpHandler.validator = o.Validator

	e.POST("/migration", httpHandler.Migration)
	e.GET("/user", httpHandler.GetAllUser)
	e.GET("/user/:uid", httpHandler.GetUser)
	e.PATCH("/user/:uid", httpHandler.UpdateUser)
	e.DELETE("/user/:uid", httpHandler.DeleteUser)
	e.PUT("/user", httpHandler.CreateUser)
	e.POST("/login", httpHandler.Login)

}

func (httpHandler *HttpHandlerUser) Migration(c echo.Context) error {

	err := httpHandler.userService.Migration(c.Request().Context())

	if err != nil {
		return c.JSON(http.StatusOK, middleware.MakeErrorResponse(err.Error()))
	}

	return c.JSON(http.StatusOK, middleware.MakeSuccessResponse(nil))
}

func (httpHandler *HttpHandlerUser) GetAllUser(c echo.Context) error {

	request := new(entity.RequestGetList)
	if err := c.Bind(request); err != nil {
		return c.JSON(http.StatusOK, middleware.MakeFailResponse(err.Error()))
	}

	if err := httpHandler.validator.ValidateStruct(c.Request().Context(), request); err != nil {
		return c.JSON(http.StatusOK, middleware.MakeFailResponse(err.Error()))
	}
	result, err := httpHandler.userService.GetAllUser(c.Request().Context(), request)

	if err != nil {
		return c.JSON(http.StatusOK, middleware.MakeErrorResponse(err.Error()))
	}

	return c.JSON(http.StatusOK, middleware.MakeSuccessResponse(result))
}

func (httpHandler *HttpHandlerUser) GetUser(c echo.Context) error {
	var err error

	uid := c.Param("uid")
	if uid == "" {
		return c.JSON(http.StatusOK, middleware.MakeFailResponse("empty uid"))
	}

	result, err := httpHandler.userService.GetUser(c.Request().Context(), &uid)

	if err != nil {
		return c.JSON(http.StatusOK, middleware.MakeErrorResponse(err.Error()))
	}

	return c.JSON(http.StatusOK, middleware.MakeSuccessResponse(result))
}

func (httpHandler *HttpHandlerUser) CreateUser(c echo.Context) error {

	request := new(entity.RequestCreateUser)
	if err := c.Bind(request); err != nil {
		return c.JSON(http.StatusOK, middleware.MakeFailResponse(err.Error()))
	}

	if err := httpHandler.validator.ValidateStruct(c.Request().Context(), request); err != nil {
		return c.JSON(http.StatusOK, middleware.MakeFailResponse(err.Error()))
	}

	err := httpHandler.userService.CreateUser(c.Request().Context(), request)
	if err != nil {
		return c.JSON(http.StatusOK, middleware.MakeErrorResponse(err.Error()))
	}

	return c.JSON(http.StatusOK, middleware.MakeSuccessResponse(nil))
}

func (httpHandler *HttpHandlerUser) Login(c echo.Context) error {

	request := new(entity.RequestLoginUser)
	if err := c.Bind(request); err != nil {
		return c.JSON(http.StatusOK, middleware.MakeFailResponse(err.Error()))
	}

	if err := httpHandler.validator.ValidateStruct(c.Request().Context(), request); err != nil {
		return c.JSON(http.StatusOK, middleware.MakeFailResponse(err.Error()))
	}

	result, err := httpHandler.userService.LoginUser(c.Request().Context(), request)
	if err != nil {
		return c.JSON(http.StatusOK, middleware.MakeErrorResponse(err.Error()))
	}

	return c.JSON(http.StatusOK, middleware.MakeSuccessResponse(result))
}

func (httpHandler *HttpHandlerUser) UpdateUser(c echo.Context) error {
	var err error
	var user *entity.User

	uid := c.Param("uid")
	if uid == "" {
		return c.JSON(http.StatusOK, middleware.MakeFailResponse("empty uid"))
	}
	request := new(entity.RequestUpdateUser)
	if err := c.Bind(request); err != nil {
		return c.JSON(http.StatusOK, middleware.MakeFailResponse(err.Error()))
	}

	if err := httpHandler.validator.ValidateStruct(c.Request().Context(), request); err != nil {
		return c.JSON(http.StatusOK, middleware.MakeFailResponse(err.Error()))
	}
	err = httpHandler.userService.UpdateUser(c.Request().Context(), &uid, request)

	if err != nil {
		return c.JSON(http.StatusOK, middleware.MakeErrorResponse(err.Error()))
	}

	return c.JSON(http.StatusOK, middleware.MakeSuccessResponse(user))
}
func (httpHandler *HttpHandlerUser) DeleteUser(c echo.Context) error {
	var err error

	uid := c.Param("uid")
	if uid == "" {
		return c.JSON(http.StatusOK, middleware.MakeFailResponse("empty uid"))
	}

	err = httpHandler.userService.DeleteUser(c.Request().Context(), &uid)

	if err != nil {
		return c.JSON(http.StatusOK, middleware.MakeErrorResponse(err.Error()))
	}

	return c.JSON(http.StatusOK, middleware.MakeSuccessResponse(nil))
}
