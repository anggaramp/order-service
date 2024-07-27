package customer

import (
	"context"
	"errors"
	"github.com/google/uuid"
	"order-service/core/entity"
	"order-service/data_source/mysql_datasource"
)

func (m *module) GetAllCustomer(ctx context.Context, request *entity.RequestGetList) (res *entity.ResponseGetAllCustomer, err error) {
	queryOption := mysql_datasource.QueryOption{
		Limit:  request.Limit,
		Cursor: request.Cursor,
	}

	if request.Keyword != "" {
		queryOption.Filter = map[string]interface{}{
			"Name": map[string]interface{}{
				"field":      "name",
				"keyword":    request.Keyword,
				"searchType": "text",
				"match":      "contain",
			},
		}
	}

	result, err := m.repo.GetAllCustomer(ctx, queryOption)

	if err != nil {
		return nil, err
	}

	return result, nil
}

func (m *module) GetCustomer(ctx context.Context, uid *string) (res *entity.ResponseGetCustomer, err error) {

	customer, err := m.repo.GetCustomerByUid(ctx, uid)

	if err != nil {
		return nil, err
	}

	return entity.ToResponseGetCustomer(customer), nil
}
func (m *module) CreateCustomer(ctx context.Context, request *entity.RequestCreateCustomer) error {

	_, err := m.repo.GetCustomerByEmail(ctx, &request.Email)
	if err == nil {
		return errors.New("email exist")
	}

	customer := &entity.Customer{
		Email:   request.Email,
		Name:    request.Name,
		Address: request.Address,
		Mobile:  request.Mobile,
		UserId:  request.UserId,
	}

	err = m.repo.CreateCustomer(ctx, customer)

	if err != nil {
		return err
	}

	return nil
}

func (m *module) UpdateCustomer(ctx context.Context, uid *string, request *entity.RequestUpdateCustomer) error {
	propertyMap := map[string]interface{}{
		"email":   request.Email,
		"name":    request.Name,
		"address": request.Address,
		"mobile":  request.Mobile,
	}

	err := m.repo.UpdateCustomer(ctx, &entity.Customer{MetaData: entity.MetaData{UID: uuid.MustParse(*uid)}}, propertyMap)

	if err != nil {
		return err
	}

	return err
}

func (m *module) DeleteCustomer(ctx context.Context, uid *string) error {
	propertyMap := []map[string]interface{}{
		{
			"key":      "uid",
			"operator": "=",
			"value":    *uid,
		},
	}

	err := m.repo.DeleteCustomer(ctx, &entity.Customer{}, propertyMap)

	if err != nil {
		return err
	}

	return nil
}
