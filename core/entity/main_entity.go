package entity

import (
	"github.com/golang-jwt/jwt/v5"
	"github.com/google/uuid"
	"gorm.io/gorm"
	"time"
)

type MetaData struct {
	ID               int64     `gorm:"primary_key;column:id;AUTO_INCREMENT;index" json:"id"`
	UID              uuid.UUID `gorm:"type:char(36);column:uid;index"  json:"uid"`
	CreatedTimestamp time.Time `gorm:"type:timestamp;default:CURRENT_TIMESTAMP" json:"created_timestamp"`
	UpdatedTimestamp time.Time `gorm:"type:timestamp;default:CURRENT_TIMESTAMP ON UPDATE CURRENT_TIMESTAMP" json:"updated_timestamp"`
}
type User struct {
	MetaData
	Email    string `gorm:"email"`
	Username string `gorm:"username"`
	Password string `gorm:"password"`
}
type Customer struct {
	MetaData
	Name    string  `gorm:"name"`
	Email   string  `gorm:"email"`
	Address string  `gorm:"address"`
	Mobile  string  `gorm:"mobile"`
	Orders  []Order `gorm:"orders"`
	UserId  int64   `gorm:"index;column:user_id"`
	User    User    `gorm:"foreignKey:UserId"` // Relationship
}

type Order struct {
	MetaData
	GoodsName   string   `gorm:"goods_name"`
	Description string   `gorm:"description"`
	Amount      float64  `gorm:"amount"`
	CustomerId  int64    `gorm:"index;column:customer_id"` // Foreign key
	Customer    Customer `gorm:"foreignKey:CustomerId"`    // Relationship
}

type JwtClaims struct {
	UserId int64 `json:"user_id"`
	jwt.RegisteredClaims
}

func (c *Customer) BeforeCreate(tx *gorm.DB) (err error) {
	c.UID = uuid.New()
	c.CreatedTimestamp = time.Now()
	c.UpdatedTimestamp = time.Now()
	return
}

func (c *Customer) BeforeUpdate(tx *gorm.DB) (err error) {
	c.UpdatedTimestamp = time.Now()
	return
}

func (o *Order) BeforeCreate(tx *gorm.DB) (err error) {
	o.UID = uuid.New()
	o.CreatedTimestamp = time.Now()
	o.UpdatedTimestamp = time.Now()
	return
}
func (o *Order) BeforeUpdate(tx *gorm.DB) (err error) {
	o.UpdatedTimestamp = time.Now()
	return
}

func (u *User) BeforeCreate(tx *gorm.DB) (err error) {
	u.UID = uuid.New()
	u.CreatedTimestamp = time.Now()
	u.UpdatedTimestamp = time.Now()
	return
}
func (u *User) BeforeUpdate(tx *gorm.DB) (err error) {
	u.UpdatedTimestamp = time.Now()
	return
}
