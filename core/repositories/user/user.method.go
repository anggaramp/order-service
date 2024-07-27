package user

import (
	"context"
	"gorm.io/gorm"
	"order-service/core/entity"
)

func (m *module) AutoMigration(ctx context.Context) error {
	err := m.getDBMySql().AutoMigrate(
		&entity.User{},
		&entity.Customer{},
		&entity.Order{},
	)
	return err
}
func (m *module) CreateUser(ctx context.Context, user *entity.User) error {
	return m.mysqlDatasource.Create(m.getDBMySql(), user)
}

func (m *module) UpdateUser(ctx context.Context, user *entity.User, updateValue map[string]interface{}) error {
	uid := user.UID.String()
	err := m.getDBMySql().Transaction(func(tx *gorm.DB) error {
		return m.mysqlDatasource.Update(m.getDBMySql(), &uid, user, updateValue)
	})
	return err
}

func (m *module) DeleteUser(ctx context.Context, user *entity.User, condition []map[string]interface{}) error {
	err := m.getDBMySql().Transaction(func(tx *gorm.DB) error {
		return m.mysqlDatasource.Delete(tx, user, condition)
	})
	return err
}
