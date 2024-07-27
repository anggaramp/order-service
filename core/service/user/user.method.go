package user

import (
	"context"
	"errors"
	"github.com/golang-jwt/jwt/v5"
	"github.com/google/uuid"
	"order-service/core/entity"
	"order-service/data_source/mysql_datasource"
	"order-service/shared/util"
	"time"
)

func (m *module) Migration(ctx context.Context) error {
	err := m.repo.AutoMigration(ctx)
	if err != nil {
		return err
	}
	return nil
}

func (m *module) LoginUser(ctx context.Context, request *entity.RequestLoginUser) (res *entity.ResponseLogin, err error) {

	user, err := m.repo.GetUserByEmail(ctx, &request.Email)
	if err != nil {
		return nil, err
	}

	if !util.ValidatePassword(user.Password, request.Password) {
		return nil, errors.New("wrong password")
	}

	claims := &entity.JwtClaims{
		user.ID,
		jwt.RegisteredClaims{
			ExpiresAt: jwt.NewNumericDate(time.Now().Add(time.Hour * 72)),
		},
	}

	newToken := jwt.NewWithClaims(jwt.SigningMethodHS256, claims)

	result, err := newToken.SignedString([]byte("secret"))
	if err != nil {
		return nil, err
	}

	return &entity.ResponseLogin{Token: result}, nil
}

func (m *module) GetAllUser(ctx context.Context, request *entity.RequestGetList) (res *entity.ResponseGetAllUser, err error) {

	queryOption := mysql_datasource.QueryOption{
		Limit:  request.Limit,
		Cursor: request.Cursor,
	}
	if request.Keyword != "" {
		queryOption.Filter = map[string]interface{}{
			"Username": map[string]interface{}{
				"field":      "username",
				"keyword":    request.Keyword,
				"searchType": "text",
				"match":      "contain",
			},
		}
	}
	result, err := m.repo.GetAllUser(ctx, queryOption)

	if err != nil {
		return nil, err
	}

	return result, nil
}

func (m *module) GetUser(ctx context.Context, uid *string) (res *entity.ResponseGetUser, err error) {

	user, err := m.repo.GetUserByUid(ctx, uid)

	if err != nil {
		return nil, err
	}

	return entity.ToResponseGetUser(user), nil
}
func (m *module) CreateUser(ctx context.Context, request *entity.RequestCreateUser) error {

	_, err := m.repo.GetUserByEmail(ctx, &request.Email)
	if err == nil {
		return errors.New("email exist")
	}

	user := &entity.User{
		Email:    request.Email,
		Username: request.Username,
		Password: util.HashPassword(request.Password),
	}

	err = m.repo.CreateUser(ctx, user)

	if err != nil {
		return err
	}

	return nil
}

func (m *module) UpdateUser(ctx context.Context, uid *string, request *entity.RequestUpdateUser) error {
	propertyMap := map[string]interface{}{
		"email":    request.Email,
		"username": request.Username,
		"password": request.Password,
	}

	err := m.repo.UpdateUser(ctx, &entity.User{MetaData: entity.MetaData{UID: uuid.MustParse(*uid)}}, propertyMap)

	if err != nil {
		return err
	}

	return err
}

func (m *module) DeleteUser(ctx context.Context, uid *string) error {
	propertyMap := []map[string]interface{}{
		{
			"key":      "uid",
			"operator": "=",
			"value":    *uid,
		},
	}

	err := m.repo.DeleteUser(ctx, &entity.User{}, propertyMap)

	if err != nil {
		return err
	}

	return nil
}
