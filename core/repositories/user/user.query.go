package user

import (
	"context"
	"fmt"
	"order-service/core/entity"
	"order-service/data_source/mysql_datasource"
)

func (m *module) GetAllUser(ctx context.Context, queryOption mysql_datasource.QueryOption) (res *entity.ResponseGetAllUser, err error) {
	var resp interface{}

	resp, err = m.mysqlDatasource.GetList(m.getDBMySql(), entity.User{}, &[]entity.User{}, queryOption)

	user, ok := resp.(*[]entity.User)
	if !ok {
		return nil, fmt.Errorf("invalid parsing data to struct")
	}

	resultData := entity.ToDataGetAllUser(user)

	result := entity.ResponseGetAllUser{
		Pagination: entity.Pagination{
			PrevCursor: queryOption.Cursor,
			Limit:      queryOption.Limit,
			HasNext:    false,
		},
	}

	if len(resultData) > queryOption.Limit {
		result.NextCursor = resultData[queryOption.Limit].UID.String()
		result.Users = resultData[:queryOption.Limit]
		result.HasNext = true
	}
	if len(resultData) <= queryOption.Limit {
		result.Users = resultData
	}

	return &result, nil
}

func (m *module) GetUserByUid(ctx context.Context, uid *string) (user *entity.User, err error) {
	var resp interface{}
	condition := []map[string]interface{}{
		{
			"key":      "uid",
			"operator": "=",
			"value":    *uid,
		},
	}

	resp, err = m.mysqlDatasource.Get(m.getDBMySql(), &entity.User{}, condition)
	user, ok := resp.(*entity.User)
	if !ok {
		return nil, fmt.Errorf("invalid parsing data to struct")
	}

	return
}
func (m *module) GetUserByEmail(ctx context.Context, email *string) (user *entity.User, err error) {
	var resp interface{}
	condition := []map[string]interface{}{
		{
			"key":      "email",
			"operator": "=",
			"value":    *email,
		},
	}

	resp, err = m.mysqlDatasource.Get(m.getDBMySql(), &entity.User{}, condition)
	user, ok := resp.(*entity.User)
	if !ok {
		return nil, fmt.Errorf("invalid parsing data to struct")
	}

	return
}
