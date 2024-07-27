package order

import (
	"context"
	"gorm.io/gorm"
	"order-service/core/entity"
)

func (m *module) AutoMigration() error {
	err := m.getDBMySql().AutoMigrate(
		&entity.User{},
		&entity.Customer{},
		&entity.Order{},
	)
	return err
}
func (m *module) CreateCustomer(ctx context.Context, order *entity.Customer) error {
	return m.mysqlDatasource.Create(m.getDBMySql(), order)
}

func (m *module) UpdateCustomer(ctx context.Context, order *entity.Customer, updateValue map[string]interface{}) error {
	uid := order.UID.String()
	err := m.getDBMySql().Transaction(func(tx *gorm.DB) error {
		return m.mysqlDatasource.Update(m.getDBMySql(), &uid, order, updateValue)
	})
	return err
}

func (m *module) DeleteCustomer(ctx context.Context, order *entity.Customer, condition []map[string]interface{}) error {
	err := m.getDBMySql().Transaction(func(tx *gorm.DB) error {
		return m.mysqlDatasource.Delete(tx, order, condition)
	})
	return err
}
