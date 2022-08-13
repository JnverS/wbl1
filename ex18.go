package main

import (
	"fmt"
	"sync"
)

/*
	Реализовать структуру-счетчик, которая будет инкрементироваться в
	конкурентной среде. По завершению программа должна выводить итоговое
	значение счетчика.
*/

type Counter struct {
	num int
	sync.Mutex
}

func (c *Counter) Increment() {
	c.Lock()
	defer c.Unlock()
	c.num++
}

func (c *Counter) Get() int{
	return c.num
}

func main() {
	counter := &Counter{
		num :0,
	}
	stop := make(chan struct{})
	go Start(counter, stop)

	select {
	case <-stop:
		fmt.Println("Итоговое значение счетчика:", counter.Get())
	}
}

func Start(counter *Counter, stop chan struct{}) {
	wg := sync.WaitGroup{}

	for i := 0; i < 15; i++ {
		wg.Add(1)
		go func (num int, counter *Counter, wg *sync.WaitGroup) {
			defer wg.Done()
			fmt.Println(num, "увеличил счетчик на 1")
			counter.Increment()
		}(i, counter, &wg)
	}
	wg.Wait()
	stop <- struct{}{}
	close(stop)
}
