package order

import (
	"gorm.io/gorm"
	"order-service/core/repositories"
	"order-service/data_source/mysql_datasource"
)

type module struct {
	mysqlDatasource *mysql_datasource.MysqlDatasource
}
type Opts struct {
	MysqlDatasource *mysql_datasource.MysqlDatasource
}

func New(o Opts) repositories.OrderRepository {
	return &module{
		mysqlDatasource: o.MysqlDatasource,
	}
}

func (m *module) getDBMySql() (db *gorm.DB) {
	return m.mysqlDatasource.GetDB()
}
