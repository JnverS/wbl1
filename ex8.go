package main

/*
	Дана переменная int64. Разработать программу которая устанавливает i-й бит в
	1 или 0
*/

import (
	"fmt"
	"strconv"
)

func main() {
	var num int64

	num = 100
	fmt.Println("before:", strconv.FormatInt(num, 2))

	// поменяем 2 бит с 1 на 0
	// операция И НЕ (если 1, то 0, если 0 то неизменяется)
	num &^= 1 << 2
	fmt.Println("after :", strconv.FormatInt(num, 2))

	// поменяем 3 бит с 0 на 1
	// операция ИЛИ (1 если хотя бы один 1)
	num |= 1 << 3
	fmt.Println("after :", strconv.FormatInt(num, 2))
}


/* другие поразрядные операции:
	операция И (1 если оба 1, 0 если хотя бы один 0)
	операция исключающее или (1 если один из 1)
*/

