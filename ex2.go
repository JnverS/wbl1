package main

import (
	"fmt"
	"sync"
)

func main() {
	numbers := [5]int{2, 4, 6, 8, 10}
	var wg sync.WaitGroup

	// добавляем количество задач, которые мы ожидаем
	wg.Add(len(numbers))
	for i := 0; i < 5; i++ {
		go func(num int) {
			// сообщаем, что задача закончена
			defer wg.Done()
			fmt.Println(num * num)
		}(numbers[i])
	}
	// ждем пока не выполнятся все задачи
	wg.Wait()
}
