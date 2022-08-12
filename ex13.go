package main

import "fmt"

/*
	Поменять местами два числа без создания временной переменной.
*/

func main() {
	a := 5
	b := 15
	fmt.Println("before a =", a, "b =", b)

	a += b
	b = a - b
	a -=b
	fmt.Println("after  a =", a, "b =", b)
}
