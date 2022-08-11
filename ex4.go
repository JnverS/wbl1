package main

import (
	"fmt"
	"os"
	"os/signal"
	"sync"
	"syscall"
)

type Worker struct {
	id int
	taskChan chan string
	quit chan bool
}

func (w Worker) Start(wg *sync.WaitGroup) {
	for {
		select {
		case task := <- w.taskChan:
				fmt.Println(w.id, task)
			case <- w.quit:
				return
			}
		}
}

func (w Worker) Stop() {
	go func() {
		w.quit <- true
	}()
}

func NewWorker(channel chan string, Id int) *Worker {
	return &Worker{
		id: Id,
		taskChan: channel,
		quit: make(chan bool),
	}
}

func main() {
	str := "Тебе осталось 7 дней"
	var wg sync.WaitGroup
	var N int
	fmt.Println("Введите количество врокеров.")
	fmt.Scan(&N)

	ch := make(chan string)


	for i := 0; i < N; i++ {
		worker := NewWorker(ch, i)
		worker.Start(&wg)
	}

	signalChan := make(chan os.Signal, 1)
	signal.Notify(signalChan, os.Interrupt, syscall.SIGINT)
	<-signalChan

	for {
		ch <- str
	}


}

