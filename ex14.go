package main

import (
	"fmt"
	"reflect"
)

/*
	Разработать программу, которая в рантайме способна определить тип
	переменной: int, string, bool, channel из переменной типа interface{}.
*/

func main() {
	var i, str, b interface{}

	i = 42
	str = "Hello"
	b = true
	ch := make(chan int)

	// определение типа с использованием пакета reflect
	// kind возвращает константу, указывающую какой примитивный тип хранится
	fmt.Println(reflect.TypeOf(i).Kind())
	fmt.Println(reflect.TypeOf(str).Kind())
	fmt.Println(reflect.TypeOf(b).Kind())
	fmt.Println(reflect.TypeOf(ch).Kind())

	// определение типа с использованием пакета fmt
	fmt.Printf("Type: %T\n", i)
	fmt.Printf("Type: %T\n", str)
	fmt.Printf("Type: %T\n", b)
	fmt.Printf("Type: %T\n", ch)

	// определение типа с использованием switch
	switch i.(type) {
	case int:
		fmt.Println("Type: int")
	case string:
		fmt.Println("Type: string")
	case bool:
		fmt.Println("Type: bool")
	default:
		fmt.Println("unexpected type")
	}
}
