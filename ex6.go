package main

/*
	Реализовать все возможные способы остановки выполнения горутины.
*/

import (
	"context"
	"fmt"
	"os"
	"os/signal"
	"sync"
	"syscall"
	"time"
)

func main() {
	wg := sync.WaitGroup{}
	wg.Add(1)
	// с помощью метода After
	go func() {
		defer wg.Done()
		// блокируется пока не пройдет время
		timer := time.After(3*time.Second)
		<-timer
		fmt.Println("Stop after")
	}()
	wg.Wait()
	// с помощью контекста
	// здесь завершение по тайм-ауту, также можно использовать дедлайн и cancel
	ctx, cancel := context.WithTimeout(context.Background(), 3*time.Second)
	defer cancel()

	wg.Add(1)
	go func() {
		defer wg.Done()
		for {
			select {
			case <-ctx.Done():
				fmt.Println("Stop context")
				return
			default:
				fmt.Println("with context")
				time.Sleep(time.Second)
			}
		}
	}()
	wg.Wait()
	// с помощью сигнала
	sigChan := make(chan os.Signal, 1)
	cleanup := make(chan bool)
	signal.Notify(sigChan, syscall.SIGINT)

	go func() {
		for {
			select {
			//ожидание сигнала прерывания ctrl+C
			case <-sigChan:
				fmt.Println("Stop signal")
				//передаем в мейн что горутина завершена и можно выходить
				cleanup <- true
				return
			default:
				fmt.Println("with signal")
				time.Sleep(time.Second)
			}

		}
	}()
	//ждем завершения горутины
	<- cleanup
}

