//1 задание

package main

import (
	"fmt"
	"sync"
)

func count(ch <-chan int, wg *sync.WaitGroup) {
	defer wg.Done()
	for num := range ch {
		fmt.Println("Processed:", num*num) // Возведение числа в квадрат
	}
}

func main() {
	ch := make(chan int)
	var wg sync.WaitGroup

	wg.Add(1)
	go count(ch, &wg)

	// Отправляем числа в канал
	for i := 1; i <= 5; i++ {
		ch <- i
	}
	close(ch) // Закрываем канал

	wg.Wait() // Ждем завершения горутины
	fmt.Println("All tasks completed.")
}
