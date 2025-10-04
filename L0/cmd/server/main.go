package main

import (
	"L0/internal/cache"
	"L0/internal/config"
	"L0/internal/db"
	"L0/internal/kafka"

	"log"
	"net/http"

	"github.com/gin-gonic/gin"
	"github.com/joho/godotenv"
	"gorm.io/gorm"
)

func init() {
	if err := godotenv.Load(); err != nil {
		log.Print(".env файл не найден.")
	}
}

// Основная логика по поиску заказа
func getOrder(dbConn *gorm.DB) gin.HandlerFunc {
	return func(c *gin.Context) {
		id := c.Param("id")

		// Проверяем есть ли в мапе
		if order, found := cache.GetOrderByID(id); found {
			c.IndentedJSON(http.StatusOK, order)
			return
		}

		// Ищем в БД
		order, err := db.GetOrderFromDBByID(dbConn, id)
		if err != nil {
			c.IndentedJSON(http.StatusNotFound, gin.H{"error": "Заказ не найден ни в кеше, ни в БД"})
			return
		}

		// сохраняем
		cache.SetOrder(order)

		c.IndentedJSON(http.StatusOK, order)
	}
}

func main() {
	conf := config.New()
	if conf.PostgresURL != "" {
		log.Println(conf.PostgresURL)
	} else {
		log.Println("URL для PostgeSQL не задан.")
	}

	dbConn := db.ConnectToDB(conf.PostgresURL)
	if err := db.InitDB(dbConn); err != nil {
		log.Fatal("Ошибка инициализации БД:", err)
	}

	// Иницилизация мапы, подключения к БД и консьюмера
	cache.InitCache(dbConn)
	db.InitDB(dbConn)
	kafka.StartConsumer(dbConn, conf.KafkaTopic)

	router := gin.Default()
	router.GET("/order/:id", getOrder(dbConn))

	router.LoadHTMLGlob("web/*")
	router.GET("/", func(c *gin.Context) {
		c.HTML(http.StatusOK, "index.html", gin.H{})
	})

	router.Run("localhost:8080")
}
