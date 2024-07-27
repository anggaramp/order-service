package order

import (
	"context"
	"errors"
	"github.com/google/uuid"
	"order-service/core/entity"
	"order-service/data_source/mysql_datasource"
)

func (m *module) GetAllOrder(ctx context.Context, request *entity.RequestGetList) (res *entity.ResponseGetAllOrder, err error) {
	queryOption := mysql_datasource.QueryOption{
		Limit:  request.Limit,
		Cursor: request.Cursor,
	}
	if request.Keyword != "" {
		queryOption.Filter = map[string]interface{}{
			"GoodsName": map[string]interface{}{
				"field":      "goods_name",
				"keyword":    request.Keyword,
				"searchType": "text",
				"match":      "contain",
			},
		}
	}
	result, err := m.repo.GetAllOrder(ctx, queryOption)

	if err != nil {
		return nil, err
	}

	return result, nil
}

func (m *module) GetOrder(ctx context.Context, uid *string) (res *entity.ResponseGetOrder, err error) {

	order, err := m.repo.GetOrderByUid(ctx, uid)

	if err != nil {
		return nil, err
	}

	return entity.ToResponseGetOrder(order), nil
}
func (m *module) CreateOrder(ctx context.Context, request *entity.RequestCreateOrder) error {

	customer, err := m.customerRepo.GetCustomerByUid(ctx, &request.CustomerUId)
	if err != nil {
		return errors.New("customer not exist")
	}

	order := &entity.Order{
		GoodsName:   request.GoodsName,
		Description: request.Description,
		Amount:      request.Amount,
		CustomerId:  customer.ID,
	}

	err = m.repo.CreateOrder(ctx, order)

	if err != nil {
		return err
	}

	return nil
}

func (m *module) UpdateOrder(ctx context.Context, uid *string, request *entity.RequestUpdateOrder) error {
	propertyMap := map[string]interface{}{
		"goods_name":  request.GoodsName,
		"description": request.Description,
		"amount":      request.Amount,
	}

	err := m.repo.UpdateOrder(ctx, &entity.Order{MetaData: entity.MetaData{UID: uuid.MustParse(*uid)}}, propertyMap)

	if err != nil {
		return err
	}

	return err
}

func (m *module) DeleteOrder(ctx context.Context, uid *string) error {
	propertyMap := []map[string]interface{}{
		{
			"key":      "uid",
			"operator": "=",
			"value":    *uid,
		},
	}

	err := m.repo.DeleteOrder(ctx, &entity.Order{}, propertyMap)

	if err != nil {
		return err
	}

	return nil
}
