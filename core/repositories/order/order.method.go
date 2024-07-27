package order

import (
	"context"
	"gorm.io/gorm"
	"order-service/core/entity"
)

func (m *module) AutoMigration() error {
	err := m.getDBMySql().AutoMigrate(
		&entity.User{},
		&entity.Order{},
		&entity.Order{},
	)
	return err
}
func (m *module) CreateOrder(ctx context.Context, order *entity.Order) error {
	return m.mysqlDatasource.Create(m.getDBMySql(), order)
}

func (m *module) UpdateOrder(ctx context.Context, order *entity.Order, updateValue map[string]interface{}) error {
	uid := order.UID.String()
	err := m.getDBMySql().Transaction(func(tx *gorm.DB) error {
		return m.mysqlDatasource.Update(m.getDBMySql(), &uid, order, updateValue)
	})
	return err
}

func (m *module) DeleteOrder(ctx context.Context, order *entity.Order, condition []map[string]interface{}) error {
	err := m.getDBMySql().Transaction(func(tx *gorm.DB) error {
		return m.mysqlDatasource.Delete(tx, order, condition)
	})
	return err
}
