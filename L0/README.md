# Решение тестового задания L0.
#### Микросервис на Go с использованием базы данных (PostgreSQL) и очереди сообщений (Kafka). Сервис будет получать данные заказов из очереди, сохраняет их в базу данных и кэширует в памяти (map) для быстрого доступа.

### Инструкция по установке

##### Клонируем репозиторий:
```bash
git@github.com:Fluttik/techschool_wb_go.git
```

##### Переходим в директорию с проектом
```bash
cd techschool_wb_go/L0
```

##### Cоздаем .env файл
Необходимо создать .env файл по примеру .env.example

##### Запускаем билд
```bash
docker-compose up --build
```
### Инструкция по использованию
#### Запуск микросервиса Golang
##### Переходим в директорию с main файлом
```bash
cd techschool_wb_go/L0/cmd/server
```

##### Запускаем сервис
```bash
go run main.go
```

#### Запуск kafka и отправление сообщений в топик
##### Открываем терминал внутри контейрена kafka
```bash
docker exec -it kafka
```
##### Создать новый топик
Имя топика должно соответсвовать имени указанному в .env файле
```bash
kafka-topics --create --topic topic_name --bootstrap-server localhost:9092
```

##### Зайти в producer kafka
```bash
kafka-console-producer --topic topic_name --bootstrap-server localhost:9092
```
После чего в терминале можно будет отправлять сообщения в топик


