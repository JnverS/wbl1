package main

import "fmt"

/*
	Разработать конвейер чисел. Даны два канала: в первый пишутся числа (x) из
	массива, во второй — результат операции x*2, после чего данные из второго
	канала должны выводиться в stdout.
*/

func main() {
	inChan, outChan := make(chan int, 1), make(chan int, 1)
	arr := [10]int{0,1,2,3,4,5,6,7,8,9}


	// горутина для записи в первый канал из массива
	go func() {
		for x := 0; x < len(arr); x++ {
			inChan <- arr[x]
		}
		close(inChan)
	}()
	// горутина для записи во второй канал х*2
	 go func() {
		for x := range inChan {
			outChan <- x * 2
		}
		close(outChan)
	 }()
	// распечатываем данные из второго канала
	for x := range outChan {
		fmt.Println(x)
	}
}
