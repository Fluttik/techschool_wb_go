package main

import (
	"fmt"
	"os"
	"strconv"
	"time"
)

func worker(id int, messages <-chan int) {
	for msg := range messages {
		fmt.Printf("Воркер %d получил: %d\n", id, msg)
	}
}

func initializate(num_workers int, pool <-chan int) {
	for n := 0; n < num_workers; n++ {
		go worker(n, pool)
	}
}

func main() {
	numWorkers := 2
	if len(os.Args) == 2 {
		numWorkers, err := strconv.Atoi(os.Args[1])
		if err != nil || numWorkers <= 0 {
			fmt.Println("Ошибка: количество воркеров должно быть положительным целым числом")
			os.Exit(1)
		}
	}

	fmt.Printf("Запуск %d воркеров...\n", numWorkers)

	intCh := make(chan int)
	initializate(numWorkers, intCh)
	counter := 0
	for {
		fmt.Println("Отправляем число:", counter)
		intCh <- counter
		counter++
		time.Sleep(5000 * time.Millisecond)
	}
}
