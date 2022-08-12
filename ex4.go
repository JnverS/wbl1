package main

/*
	Реализовать постоянную запись данных в канал (главный поток). Реализовать
	набор из N воркеров, которые читают произвольные данные из канала и
	выводят в stdout. Необходима возможность выбора количества воркеров при
	старте.
	Программа должна завершаться по нажатию Ctrl+C. Выбрать и обосновать
	способ завершения работы всех воркеров.
*/

import (
	"context"
	"fmt"
	"os/signal"
	"sync"
	"syscall"
	"time"
)

func main() {
	str := "Тебе осталось 7 дней"
	var N int
	// получаем количество воркеров, которых будем создавать
	fmt.Println("Введите количество врокеров.")
	fmt.Scan(&N)

	ch := make(chan string, N)
	// создаем контекст и передаем в него ожидание сигнала прерывания
	ctx, cancel := signal.NotifyContext(context.Background(), syscall.SIGINT)
	defer cancel()
	wg := sync.WaitGroup{}

	wg.Add(N)
	for i := 0; i < N; i++ {
		k := i
		go func () {
			defer wg.Done()
			for {
				select {
				case <- ctx.Done():
					fmt.Println(k, "worker завершен")
					return
				case task, ok := <-ch:
					if !ok {
						return
					}
					fmt.Println(k, task)

				}
			}
		}()
	}

	for {
		select {
		case <-ctx.Done():
			wg.Wait()
			close(ch)
			fmt.Println("Main завершен")
			return
		case ch <- str:
			time.Sleep(time.Second)
		}
	}
}

