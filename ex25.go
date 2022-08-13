package main

import (
	"context"
	"fmt"
	"time"
)

/*
	Реализовать собственную функцию sleep.
*/

func sleep(d time.Duration) {
	<- time.After(d)
}

func sleepCtx(d time.Duration) {
	ctx, cancel := context.WithDeadline(context.Background(), time.Now().Add(d*time.Second))
	defer cancel()
	<-ctx.Done()
}

func main() {
	fmt.Println("Start")
	sleep(5 * time.Second)
	fmt.Println("Second sleep")
	sleepCtx(3)
	fmt.Println("Finish")
}
