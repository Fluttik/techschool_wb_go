package db

import (
	"log"

	"gorm.io/driver/postgres"
	"gorm.io/gorm"
)

func ConnectToDB(db_url string) *gorm.DB {
	log.Println("URL:")
	log.Println(db_url)

	db, err := gorm.Open(postgres.Open(db_url), &gorm.Config{})
	if err != nil {
		log.Fatalf("Не удалось подключится к БД: %v", err)
	}
	log.Println("Подключение установлено.")
	return db
}

func InitDB(db *gorm.DB) error {
	err := db.AutoMigrate(&Order{}, &Delivery{}, &Payment{}, &Item{})
	return err
}

func CreateOrder(db *gorm.DB, order *Order) error {
	return db.Create(order).Error
}

func GetOrderFromDBByID(db *gorm.DB, uid string) (*Order, error) {
	var order Order
	err := db.Preload("Delivery").Preload("Payment").Preload("Items").First(&order, "order_uid = ?", uid).Error
	if err != nil {
		return nil, err
	}
	return &order, nil
}

func GetOrders(db *gorm.DB, limit int) ([]Order, error) {
	var orders []Order
	err := db.Preload("Delivery").Preload("Payment").Preload("Items").
		Order("date_created DESC").Limit(limit).Find(&orders).Error
	return orders, err
}
