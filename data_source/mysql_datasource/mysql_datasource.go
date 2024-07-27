package mysql_datasource

import (
	"fmt"
	"github.com/lib/pq"
	"gorm.io/gorm"
	"gorm.io/gorm/clause"
	"strings"
)

type MysqlDatasource struct {
	Client *gorm.DB
}

type QueryOptionOrder struct {
	Field     string `json:"field"`
	Direction string `json:"direction"`
}

type QueryOption struct {
	Filter map[string]interface{} `json:"filter"`
	Order  []QueryOptionOrder     `json:"order"`
	Limit  int                    `json:"limit"`
	Cursor string                 `json:"cursor"`
}

func NewMysqlDatasource(client *gorm.DB) *MysqlDatasource {
	return &MysqlDatasource{
		Client: client,
	}
}

func (r *MysqlDatasource) GetDB() (tx *gorm.DB) {
	tx = r.Client
	return
}

func (r *MysqlDatasource) Create(tx *gorm.DB, data interface{}) (err error) {
	tx = tx.Model(data)

	err = tx.Clauses(clause.Returning{}).Create(data).Error

	return
}

func (r *MysqlDatasource) Delete(tx *gorm.DB, data interface{}, conditions []map[string]interface{}) (err error) {

	for _, condition := range conditions {
		var key, operator, value string

		if _value, ok := condition["key"].(string); ok {
			key = _value
		}
		if _value, ok := condition["operator"].(string); ok {
			operator = _value
		}
		if _value, ok := condition["value"].(string); ok {
			value = _value
		}

		tx = tx.Where(fmt.Sprintf("%s %s '%v'", key, operator, value))
	}

	err = tx.Delete(data).Error

	return
}

func (r *MysqlDatasource) Update(tx *gorm.DB, uid *string, data interface{}, properties map[string]interface{}) (err error) {
	tx = tx.Model(data)

	err = tx.Clauses(clause.Returning{}).Where("uid=?", *uid).Updates(properties).Error

	return
}

func (r *MysqlDatasource) Get(tx *gorm.DB, entity interface{}, conditions []map[string]interface{}) (resp interface{}, err error) {
	tx = tx.Model(entity)

	for _, condition := range conditions {
		var key, operator, value string

		if _value, ok := condition["key"].(string); ok {
			key = _value
		}
		if _value, ok := condition["operator"].(string); ok {
			operator = _value
		}
		if _value, ok := condition["value"].(string); ok {
			value = _value
		}

		tx = tx.Where(fmt.Sprintf("%s %s '%s'", key, operator, value))
	}

	err = tx.First(entity).Error

	resp = entity

	return
}

func (r *MysqlDatasource) GetV2(tx *gorm.DB, entity interface{}, conditions []map[string]interface{}) (err error) {

	for _, condition := range conditions {
		var key, operator, value string

		if _value, ok := condition["key"].(string); ok {
			key = _value
		}
		if _value, ok := condition["operator"].(string); ok {
			operator = _value
		}
		if _value, ok := condition["value"].(string); ok {
			value = _value
		}

		tx = tx.Where(fmt.Sprintf("%s %s '%s'", key, operator, value))
	}

	err = tx.First(entity).Error

	return
}

func (r *MysqlDatasource) GetList(tx *gorm.DB, entity interface{}, data interface{}, queryOption QueryOption) (resp interface{}, err error) {
	tx = tx.Model(entity)

	// page

	if queryOption.Limit > 0 {
		tx = tx.Limit(int(queryOption.Limit) + 1)
	}
	if queryOption.Cursor != "" {
		tx = tx.Where("uid >= ?", queryOption.Cursor)
	}

	// filter
	err = r.processQueryOptionFilter(tx, queryOption.Filter)
	if nil != err {
		return
	}

	// order
	for _, order := range queryOption.Order {
		orderQuery := fmt.Sprintf("%s %s", order.Field, order.Direction)

		tx = tx.Order(orderQuery)
	}

	if len(queryOption.Order) < 1 {
		tx = tx.Order("created_timestamp desc")
	}

	err = tx.Find(data).Error

	resp = data

	return
}

func (r *MysqlDatasource) GetListWithRaw(tx *gorm.DB, data interface{}, query *string) (resp interface{}, err error) {
	if nil != query {
		tx = tx.Raw(*query)
	}

	err = tx.Find(data).Error

	resp = data

	return
}

func (r *MysqlDatasource) Query(tx *gorm.DB, data interface{}, query *string) (resp interface{}, err error) {

	if nil != query {
		tx = tx.Raw(*query)
	}

	err = tx.Scan(data).Error

	resp = data

	return
}

func (r *MysqlDatasource) processQueryOptionFilter(tx *gorm.DB, filter map[string]interface{}) (err error) {
	if nil == filter {
		return
	}

	for key, value := range filter {
		if 0 == strings.Index(key, "orSet") {
			queries, args := r.getWhereOrSet(value.([]interface{}))
			*tx = *tx.Where(strings.Join(queries, " OR "), args...)
		} else {
			query, arg := r.getWhere(value.(map[string]interface{}))
			*tx = *tx.Where(query, arg)
		}
	}

	return
}

func (r *MysqlDatasource) getWhereOrSet(items []interface{}) (queries []string, args []interface{}) {
	for _, item := range items {
		query, arg := r.getWhere(item.(map[string]interface{}))
		queries = append(queries, query)
		args = append(args, arg)
	}

	return
}

func (r *MysqlDatasource) getWhere(item map[string]interface{}) (query string, arg interface{}) {

	field := item["field"].(string)
	searchType := item["searchType"].(string)
	match := item["match"].(string)
	keyword := item["keyword"]

	field = strings.ReplaceAll(field, "\"", "'")

	//log.Debug("getWhere %s %s %s %v", field, searchType, match, keyword)

	switch searchType {
	case "text":
		switch match {
		case "contain":
			query = fmt.Sprintf("LOWER(%s) LIKE ?", field)
			arg = fmt.Sprintf("%%%s%%", strings.ToLower(keyword.(string)))
		case "startWith":
			query = fmt.Sprintf("LOWER(%s) LIKE ?", field)
			arg = fmt.Sprintf("%s%%", strings.ToLower(keyword.(string)))
		case "endWith":
			query = fmt.Sprintf("LOWER(%s) LIKE ?", field)
			arg = fmt.Sprintf("%%%s", strings.ToLower(keyword.(string)))
		case "exact":
			query = fmt.Sprintf("LOWER(%s) = ?", field)
			arg = strings.ToLower(keyword.(string))
		case "notEqual":
			query = fmt.Sprintf("LOWER(%s) != ?", field)
			arg = strings.ToLower(keyword.(string))
		case "gt":
			query = fmt.Sprintf("%s > ?", field)
			arg = keyword
		case "gte":
			query = fmt.Sprintf("%s >= ?", field)
			arg = keyword
		case "lt":
			query = fmt.Sprintf("%s < ?", field)
			arg = keyword
		case "lte":
			query = fmt.Sprintf("%s <= ?", field)
			arg = keyword
		}
	case "number":
		switch match {
		case "exact":
			query = fmt.Sprintf("%s = ?", field)
			arg = keyword.(uint64)
		case "notEqual":
			query = fmt.Sprintf("%s != ?", field)
			arg = keyword.(uint64)
		}
	case "list":
		switch match {
		case "contain", "overlap":
			if s, ok := keyword.(string); ok {
				query = fmt.Sprintf("%s = ?", field)
				arg = fmt.Sprintf("%v", s)
			} else if ss, ok := keyword.([]interface{}); ok {
				var symbol string
				switch match {
				case "contain":
					symbol = "@>"
				case "overlap":
					symbol = "&&"
				}

				query = fmt.Sprintf("%s %s ?", field, symbol)
				var ssString []string
				for _, s := range ss {
					ssString = append(ssString, s.(string))
				}
				arg = pq.Array(ssString)
			}
		}
	case "date":

		dateTypeName := "day"

		dateType, ok := item["dateType"]
		if ok {
			dateTypeName = dateType.(string)
		}

		switch match {
		case "gt":
			query = fmt.Sprintf("date_trunc('%s', %s) > ?", dateTypeName, field)
			arg = keyword
		case "gte":
			query = fmt.Sprintf("date_trunc('%s', %s) >= ?", dateTypeName, field)
			arg = keyword
		case "lt":
			query = fmt.Sprintf("date_trunc('%s', %s) < ?", dateTypeName, field)
			arg = keyword
		case "lte":
			query = fmt.Sprintf("date_trunc('%s', %s) <= ?", dateTypeName, field)
			arg = keyword
		}
	case "bool":
		switch match {
		case "exact":
			query = fmt.Sprintf("%s = ?", field)
			arg = keyword.(bool)
		}
	}

	//log.Debug("getWhere %s %v", query, arg)

	return
}
