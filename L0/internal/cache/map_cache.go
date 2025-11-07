package cache

import (
	"log"
	"sync"

	"gorm.io/gorm"

	"L0/internal/db"
)

type Cache struct {
	mu     sync.RWMutex
	orders map[string]*db.Order
}

var cache = &Cache{
	orders: make(map[string]*db.Order),
}

func InitCache(dbConn *gorm.DB) {
	// получаем послдние 100 заказов из БД для сохранения в кеше
	dbOrders, err := db.GetOrders(dbConn, 100)
	if err != nil {
		log.Println("Не удалось загрузить заказы из БД.")
	}

	cache.mu.Lock()
	defer cache.mu.Unlock()

	// заполнение кеша
	for i := range dbOrders {
		order := &dbOrders[i]
		cache.orders[order.OrderUID] = order
		log.Printf("Заказ с id %s загружен в кеш", order.OrderUID)
	}

	if len(cache.orders) == 0 {
		log.Println("Кеш пуст — в БД пока нет заказов")
	} else {
		log.Printf("В кеш загружено %d заказа(ов)\n", len(cache.orders))
	}
}

// Функкия для получения 1 заказа по его id
func GetOrderByID(id string) (*db.Order, bool) {
	cache.mu.RLock()
	defer cache.mu.RUnlock()
	order, exists := cache.orders[id]
	return order, exists
}

// Функия для получения всех заказов (использовалась для тестов, в проекте не используется)
func GetOrders() []*db.Order {
	cache.mu.RLock()
	defer cache.mu.RUnlock()

	var result []*db.Order
	for _, order := range cache.orders {
		result = append(result, order)
	}
	return result
}

// Функция для добавления заказа в кеш
func SetOrder(order *db.Order) {
	cache.mu.Lock()
	defer cache.mu.Unlock()
	cache.orders[order.OrderUID] = order
}
