package main

/*
	Разработать программу, которая будет последовательно отправлять значения в
	канал, а с другой стороны канала — читать. По истечению N секунд программа
	должна завершаться.
*/

import (
	"context"
	"fmt"
	"sync"
	"time"
)

func main() {
	var N int
	fmt.Println("Введите количество секунд:")
	fmt.Scan(&N)

	// устанавливаем время, через которое программа завершится
	withDeadline, cancel := context.WithDeadline(context.Background(), time.Now().Add(time.Second*time.Duration(N)))
	defer cancel()

	var i int
	ch := make(chan int, 1)
	wg := sync.WaitGroup{}
	wg.Add(2)
	go func() {
		defer wg.Done()
		for {
			select {
				// если время закончилось
				case <-withDeadline.Done():
					close(ch)
					return
				// пишем в канал
				case ch <- i:
					i++
			}
		}
	}()
	go func() {
		defer wg.Done()
		for {
			select {
				case <-withDeadline.Done():
					return
				// читаем из канала + проверяем что он открыт
				case v, ok := <-ch:
					if !ok {
						return
					}
					fmt.Println(v)

			}
		}
	}()
	// ждем завершения всех
	wg.Wait()
}
