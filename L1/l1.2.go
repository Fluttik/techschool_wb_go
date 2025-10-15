package main

import (
	"fmt"
	"sync"
)

var nums [5]int = [5]int{2, 4, 6, 8, 10}

func square(num int, wg *sync.WaitGroup, out chan<- int) {
	defer wg.Done()
	out <- num * num
}

func main() {
	var wg sync.WaitGroup
	results := make(chan int, len(nums))

	go func() {
		wg.Wait()
		close(results)
	}()

	for _, num := range nums {
		wg.Add(1)
		go square(num, &wg, results)
	}

	for res := range results {
		fmt.Println(res)
	}

}
