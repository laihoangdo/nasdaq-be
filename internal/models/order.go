package models

import "time"

type OrderStatus string

func (or OrderStatus) String() string {
	return string(or)
}

type Order struct {
	ID             int       `gorm:"column:id" json:"id"`
	UserID         int       `gorm:"column:user_id" json:"user_id"`
	OrderPackageID int       `gorm:"column:order_package_id" json:"order_package_id"`
	SKU            int       `gorm:"column:sku" json:"sku"`
	CoinCode       string    `gorm:"column:coin_code" json:"coin_code"`
	Time           int       `gorm:"column:time" json:"time"`
	Date           time.Time `gorm:"column:date" json:"date"`
	IsInprogress   bool      `gorm:"column:is_inprogress" json:"is_inprogress"`
	Status         bool      `gorm:"column:status" json:"status"`
	Balance        int       `gorm:"column:balance" json:"balance"`
	CreatedAt      time.Time `gorm:"column:created_at" json:"created_at"`
	UpdatedAt      time.Time `gorm:"column:updated_at" json:"updated_at"`
}

func (*Order) TableName() string {
	return "orders"
}
