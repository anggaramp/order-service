package entity

import (
	"github.com/google/uuid"
	"time"
)

type ResponseLogin struct {
	Token string `json:"token"`
}

type ResponseGetUser struct {
	UID              uuid.UUID `json:"uid"`
	Email            string    `json:"email"`
	Username         string    `json:"username"`
	CreatedTimestamp time.Time `json:"created_timestamp"`
	UpdatedTimestamp time.Time `json:"updated_timestamp"`
}

type Pagination struct {
	Limit      int    `json:"limit"`
	NextCursor string `json:"next_cursor"`
	PrevCursor string `json:"prev_cursor"`
	HasNext    bool   `json:"has_next"`
}

type ResponseGetAllUser struct {
	Pagination
	Users interface{} `json:"users"`
}

type ResponseGetCustomer struct {
	UID              uuid.UUID     `json:"uid"`
	Email            string        `json:"email"`
	Name             string        `json:"name"`
	Mobile           string        `json:"mobile"`
	Address          string        `json:"address"`
	Orders           []DetailOrder `json:"orders"`
	CreatedTimestamp time.Time     `json:"created_timestamp"`
	UpdatedTimestamp time.Time     `json:"updated_timestamp"`
}

type ResponseGetAllCustomer struct {
	Pagination
	Customers interface{} `json:"orders"`
}

type DetailOrder struct {
	UID              uuid.UUID `json:"uid"`
	GoodsName        string    `json:"goods_name"`
	Description      string    `json:"description"`
	Amount           float64   `json:"amount"`
	CreatedTimestamp time.Time `json:"created_timestamp"`
	UpdatedTimestamp time.Time `json:"updated_timestamp"`
}

type ResponseGetOrder struct {
	UID              uuid.UUID           `json:"uid"`
	GoodsName        string              `json:"goods_name"`
	Description      string              `json:"description"`
	Amount           float64             `json:"amount"`
	Customer         ResponseGetCustomer `json:"customer,omitempty"`
	CreatedTimestamp time.Time           `json:"created_timestamp"`
	UpdatedTimestamp time.Time           `json:"updated_timestamp"`
}

type ResponseGetAllOrder struct {
	Pagination
	Orders interface{} `json:"orders"`
}

func ToResponseGetUser(user *User) *ResponseGetUser {
	return &ResponseGetUser{
		UID:              user.UID,
		Email:            user.Email,
		Username:         user.Username,
		CreatedTimestamp: user.CreatedTimestamp,
		UpdatedTimestamp: user.UpdatedTimestamp,
	}
}

func ToDataGetAllUser(user *[]User) []ResponseGetUser {
	result := make([]ResponseGetUser, 0)
	for _, u := range *user {
		temp := ResponseGetUser{
			UID:              u.UID,
			Email:            u.Email,
			Username:         u.Username,
			CreatedTimestamp: u.CreatedTimestamp,
			UpdatedTimestamp: u.UpdatedTimestamp,
		}
		result = append(result, temp)
	}
	return result
}

func ToResponseGetCustomer(customer *Customer) *ResponseGetCustomer {
	result := ResponseGetCustomer{
		UID:              customer.UID,
		Email:            customer.Email,
		Name:             customer.Name,
		Mobile:           customer.Mobile,
		Address:          customer.Address,
		CreatedTimestamp: customer.CreatedTimestamp,
		UpdatedTimestamp: customer.UpdatedTimestamp,
	}
	for _, v := range customer.Orders {
		result.Orders = append(result.Orders, DetailOrder{
			UID:              v.UID,
			GoodsName:        v.GoodsName,
			Description:      v.Description,
			Amount:           v.Amount,
			CreatedTimestamp: v.CreatedTimestamp,
			UpdatedTimestamp: v.UpdatedTimestamp,
		})
	}
	return &result
}

func ToDataGetAllCustomer(customer *[]Customer) []ResponseGetCustomer {
	result := make([]ResponseGetCustomer, 0)
	for _, c := range *customer {
		temp := ResponseGetCustomer{
			UID:              c.UID,
			Email:            c.Email,
			Name:             c.Name,
			Mobile:           c.Mobile,
			Address:          c.Address,
			CreatedTimestamp: c.CreatedTimestamp,
			UpdatedTimestamp: c.UpdatedTimestamp,
		}
		result = append(result, temp)
	}
	return result
}

func ToResponseGetOrder(order *Order) *ResponseGetOrder {
	return &ResponseGetOrder{
		UID:         order.UID,
		GoodsName:   order.GoodsName,
		Description: order.Description,
		Amount:      order.Amount,
		Customer: ResponseGetCustomer{
			UID:              order.Customer.UID,
			Email:            order.Customer.Email,
			Name:             order.Customer.Name,
			Mobile:           order.Customer.Mobile,
			Address:          order.Customer.Address,
			CreatedTimestamp: order.Customer.CreatedTimestamp,
			UpdatedTimestamp: order.Customer.UpdatedTimestamp,
		},
		CreatedTimestamp: order.CreatedTimestamp,
		UpdatedTimestamp: order.UpdatedTimestamp,
	}
}

func ToDataGetAllOrder(order *[]Order) []ResponseGetOrder {
	result := make([]ResponseGetOrder, 0)
	for _, o := range *order {
		temp := ResponseGetOrder{
			UID:         o.UID,
			GoodsName:   o.GoodsName,
			Description: o.Description,
			Amount:      o.Amount,
			Customer: ResponseGetCustomer{
				UID:              o.Customer.UID,
				Email:            o.Customer.Email,
				Name:             o.Customer.Name,
				Mobile:           o.Customer.Mobile,
				Address:          o.Customer.Address,
				CreatedTimestamp: o.Customer.CreatedTimestamp,
				UpdatedTimestamp: o.Customer.UpdatedTimestamp,
			},
			CreatedTimestamp: o.CreatedTimestamp,
			UpdatedTimestamp: o.UpdatedTimestamp,
		}
		result = append(result, temp)
	}
	return result
}
