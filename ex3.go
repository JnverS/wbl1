package main

import (
	"fmt"
	"sync"
)

func main() {
	numbers := [5]int{2, 4, 6, 8, 10}
	var wg sync.WaitGroup
	// создаем буферизированный канал
	ch := make(chan int, 1)

	ch <- 0
	// добавляем количество задач, которые мы ожидаем
	wg.Add(len(numbers))
	for i := 0; i < 5; i++ {
		go func(num int) {
			// сообщаем, что задача закончена
			defer wg.Done()
			// записываем в канал сумму нового квадрата и предыдущих
			ch <- num * num + <- ch
		}(numbers[i])
	}
	// ждем пока не выполнятся все задачи
	wg.Wait()
	fmt.Println(<-ch)
	// закрываем канал
	close(ch)
}

