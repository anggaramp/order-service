package order

import (
	"context"
	"fmt"
	"order-service/core/entity"
	"order-service/data_source/mysql_datasource"
)

func (m *module) GetAllCustomer(ctx context.Context, queryOption mysql_datasource.QueryOption) (res *entity.ResponseGetAllCustomer, err error) {
	var resp interface{}

	resp, err = m.mysqlDatasource.GetList(m.getDBMySql(), entity.Customer{}, &[]entity.Customer{}, queryOption)

	order, ok := resp.(*[]entity.Customer)
	if !ok {
		return nil, fmt.Errorf("invalid parsing data to struct")
	}

	resultData := entity.ToDataGetAllCustomer(order)

	result := entity.ResponseGetAllCustomer{
		Pagination: entity.Pagination{
			PrevCursor: queryOption.Cursor,
			Limit:      queryOption.Limit,
			HasNext:    false,
		},
	}

	if len(resultData) > queryOption.Limit {
		result.NextCursor = resultData[queryOption.Limit].UID.String()
		result.Customers = resultData[:queryOption.Limit]
		result.HasNext = true
	}
	if len(resultData) <= queryOption.Limit {
		result.Customers = resultData
	}

	return &result, nil
}

func (m *module) GetCustomerByUid(ctx context.Context, uid *string) (order *entity.Customer, err error) {
	var resp interface{}
	condition := []map[string]interface{}{
		{
			"key":      "uid",
			"operator": "=",
			"value":    *uid,
		},
	}

	resp, err = m.mysqlDatasource.Get(m.getDBMySql().Preload("Orders"), &entity.Customer{}, condition)
	order, ok := resp.(*entity.Customer)
	if !ok {
		return nil, fmt.Errorf("invalid parsing data to struct")
	}

	return
}
func (m *module) GetCustomerByEmail(ctx context.Context, email *string) (order *entity.Customer, err error) {
	var resp interface{}
	condition := []map[string]interface{}{
		{
			"key":      "email",
			"operator": "=",
			"value":    *email,
		},
	}

	resp, err = m.mysqlDatasource.Get(m.getDBMySql(), &entity.Customer{}, condition)
	order, ok := resp.(*entity.Customer)
	if !ok {
		return nil, fmt.Errorf("invalid parsing data to struct")
	}

	return
}
