package order

import (
	"context"
	"fmt"
	"order-service/core/entity"
	"order-service/data_source/mysql_datasource"
)

func (m *module) GetAllOrder(ctx context.Context, queryOption mysql_datasource.QueryOption) (res *entity.ResponseGetAllOrder, err error) {
	var resp interface{}

	resp, err = m.mysqlDatasource.GetList(m.getDBMySql().Preload("Customer"), entity.Order{}, &[]entity.Order{}, queryOption)

	order, ok := resp.(*[]entity.Order)
	if !ok {
		return nil, fmt.Errorf("invalid parsing data to struct")
	}

	resultData := entity.ToDataGetAllOrder(order)

	result := entity.ResponseGetAllOrder{
		Pagination: entity.Pagination{
			PrevCursor: queryOption.Cursor,
			Limit:      queryOption.Limit,
			HasNext:    false,
		},
	}

	if len(resultData) > queryOption.Limit {
		result.NextCursor = resultData[queryOption.Limit].UID.String()
		result.Orders = resultData[:queryOption.Limit]
		result.HasNext = true
	}
	if len(resultData) <= queryOption.Limit {
		result.Orders = resultData
	}

	return &result, nil
}

func (m *module) GetOrderByUid(ctx context.Context, uid *string) (order *entity.Order, err error) {
	var resp interface{}
	condition := []map[string]interface{}{
		{
			"key":      "uid",
			"operator": "=",
			"value":    *uid,
		},
	}

	resp, err = m.mysqlDatasource.Get(m.getDBMySql().Preload("Customer"), &entity.Order{}, condition)
	order, ok := resp.(*entity.Order)
	if !ok {
		return nil, fmt.Errorf("invalid parsing data to struct")
	}

	return
}
