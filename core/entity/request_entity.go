package entity

type RequestLoginUser struct {
	Email    string `json:"email" validate:"required"`
	Password string `json:"password" validate:"required"`
}

type RequestCreateUser struct {
	Email    string `json:"email" validate:"required,email"`
	Username string `json:"username" validate:"required"`
	Password string `json:"password" validate:"required"`
}

type RequestUpdateUser struct {
	Email    string `json:"email" validate:"required,email"`
	Username string `json:"username" validate:"required"`
	Password string `json:"password" validate:"required"`
}

type RequestGetList struct {
	Limit   int    `query:"limit" json:"limit" validate:"required"`
	Order   string `query:"order" json:"order"`
	Cursor  string `query:"cursor" json:"cursor"`
	Keyword string `query:"keyword" json:"keyword"`
}

type RequestCreateCustomer struct {
	Email   string `json:"email" validate:"required,email"`
	Name    string `json:"name" validate:"required"`
	Address string `json:"address" validate:"required"`
	Mobile  string `json:"mobile" validate:"required"`
	UserId  int64  `json:"user_id"`
}

type RequestUpdateCustomer struct {
	Email   string `json:"email" validate:"required,email"`
	Name    string `json:"name" validate:"required"`
	Address string `json:"address" validate:"required"`
	Mobile  string `json:"mobile" validate:"required"`
}

type RequestCreateOrder struct {
	GoodsName   string  `json:"goods_name" validate:"required"`
	Description string  `json:"description" validate:"required"`
	Amount      float64 `json:"amount" validate:"required"`
	CustomerUId string  `json:"customer_uid"`
}
type RequestUpdateOrder struct {
	GoodsName   string  `json:"goods_name" validate:"required"`
	Description string  `json:"description" validate:"required"`
	Amount      float64 `json:"amount" validate:"required"`
}
