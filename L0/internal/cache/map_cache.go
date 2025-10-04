package cache

import (
	"L0/internal/db"
	"log"

	"gorm.io/gorm"
)

var orders = make(map[string]*db.Order)

func InitCache(dbConn *gorm.DB) {
	// получаем послдние 100 заказов из БД для сохранения в кеше
	dbOrders, err := db.GetOrders(dbConn, 100)
	if err != nil {
		log.Println("Не удалось загрузить заказы из БД.")
	}

	// заполнение кеша
	for i := range dbOrders {
		order := &dbOrders[i]
		orders[order.OrderUID] = order
		log.Printf("Заказ с id %s загружен в кеш", order.OrderUID)
	}

	if len(orders) == 0 {
		log.Println("Кеш пуст — в БД пока нет заказов")
	} else {
		log.Printf("В кеш загружено %d заказа(ов)\n", len(orders))
	}
}

// Функкия для получения 1 заказа по его id
func GetOrderByID(id string) (*db.Order, bool) {
	order, exists := orders[id]
	return order, exists
}

// Функия для получения всех заказов (использовалась для тестов, в проекте не используется)
func GetOrders() []*db.Order {
	var result []*db.Order
	for _, order := range orders {
		result = append(result, order)
	}
	return result
}

// Функция для добавления заказа в кеш
func SetOrder(order *db.Order) {
	orders[order.OrderUID] = order
}
