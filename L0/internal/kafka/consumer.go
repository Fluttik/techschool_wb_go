package kafka

import (
	"context"
	"encoding/json"
	"log"

	"L0/internal/db"

	"github.com/segmentio/kafka-go"
	"gorm.io/gorm"
)

func StartConsumer(dbConn *gorm.DB, topic string) {
	go func() {
		reader := kafka.NewReader(kafka.ReaderConfig{
			Brokers:        []string{"localhost:9092"},
			Topic:          topic,
			GroupID:        "my-groupID",
			CommitInterval: 0,
		})
		defer reader.Close()

		for {

			msg, err := reader.ReadMessage(context.Background())
			if err != nil {
				log.Println("Ошибка при получении сообщения")
			}
			var order db.Order
			if err := json.Unmarshal(msg.Value, &order); err != nil {
				log.Printf("Ошибка парсинга сообщения '%s' из Kafka", msg.Value)
				continue
			}
			existingOrder, err := db.GetOrderFromDBByID(dbConn, order.OrderUID)
			if err != nil {
				db.CreateOrder(dbConn, &order)
				log.Printf("Заказ c id %s добавлен в БД", order.OrderUID)
			} else {
				log.Printf("Заказ c id %s уже существует", existingOrder.OrderUID)
			}

			err = reader.CommitMessages(context.Background(), msg)
			if err != nil {
				log.Printf("Ошибка при коммите сообщения: %v", err)
			}
		}
	}()
}
