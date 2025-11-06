package db

import (
	"time"
)

type Order struct {
	OrderUID          string    `json:"order_uid" gorm:"primaryKey" validate:"required,min=1"`
	TrackNumber       string    `json:"track_number" validate:"required,min=1"`
	Entry             string    `json:"entry" validate:"required,min=1"`
	Locale            string    `json:"locale" validate:"required,oneof=ru en ch"`
	InternalSignature string    `json:"internal_signature"`
	CustomerID        string    `json:"customer_id" validate:"required,min=1"`
	DeliveryService   string    `json:"delivery_service" validate:"required,min=1"`
	ShardKey          string    `json:"shardkey" validate:"required,min=1"`
	SmID              int64     `json:"sm_id" validate:"required,gte=0"`
	DateCreated       time.Time `json:"date_created" validate:"required"`
	OofShard          string    `json:"oof_shard" validate:"required"`

	Delivery Delivery `json:"delivery" gorm:"foreignKey:OrderUID;references:OrderUID" validate:"required"`
	Payment  Payment  `json:"payment" gorm:"foreignKey:Transaction;references:OrderUID" validate:"required"`
	Items    []Item   `json:"items" gorm:"foreignKey:OrderUID;references:OrderUID" validate:"required,min=1"`
}

type Delivery struct {
	OrderUID string `json:"-" gorm:"primaryKey" validate:"required,min=1"`
	Name     string `json:"name" validate:"required,min=1"`
	Phone    string `json:"phone" validate:"required"`
	Zip      string `json:"zip" validate:"required"`
	City     string `json:"city" validate:"required"`
	Address  string `json:"address" validate:"required"`
	Region   string `json:"region" validate:"required"`
	Email    string `json:"email" validate:"required"`
}

type Payment struct {
	Transaction  string `json:"transaction" gorm:"primaryKey" validate:"required"`
	RequestID    string `json:"request_id"`
	Currency     string `json:"currency" validate:"required,len=3"`
	Provider     string `json:"provider" validate:"required"`
	Amount       int64  `json:"amount" validate:"required,gte=0"`
	PaymentDT    int64  `json:"payment_dt" validate:"required"`
	Bank         string `json:"bank" validate:"required"`
	DeliveryCost int64  `json:"delivery_cost" validate:"required,gte=0"`
	GoodsTotal   int64  `json:"goods_total" validate:"required,gte=0"`
	CustomFee    int64  `json:"custom_fee" validate:"required,gte=0"`

	OrderUID string `json:"-" gorm:"index" validate:"required"`
}

type Item struct {
	ChrtID      int64  `json:"chrt_id" gorm:"primaryKey" validate:"required"`
	OrderUID    string `json:"-" gorm:"index" validate:"required"`
	TrackNumber string `json:"track_number" validate:"required"`
	Price       int64  `json:"price" validate:"required,gte=0"`
	Rid         string `json:"rid" validate:"required"`
	Name        string `json:"name" validate:"required"`
	Sale        int64  `json:"sale" validate:"gte=0,lte=100"`
	Size        string `json:"size" validate:"required,gte=0"`
	TotalPrice  int64  `json:"total_price" validate:"required,gte=0"`
	NmID        int64  `json:"nm_id" validate:"required,gte=0"`
	Brand       string `json:"brand" validate:"required"`
	Status      int64  `json:"status" validate:"required"`
}
