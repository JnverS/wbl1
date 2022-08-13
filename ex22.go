package main

import (
	"fmt"
	"math/big"
)

/*
	Разработать программу, которая перемножает, делит, складывает, вычитает две
	числовых переменных a,b, значение которых > 2^20.
*/

// для работы с большими числами есть пакет math/big
func main() {
	// задаем значение двум большим числам
	first := big.NewInt(25000000)
	second := big.NewInt(10000000)
	// выделяем память
	addResult := new(big.Int)
	// складываем
	addResult.Add(first, second)
	divResult := new(big.Int)
	divResult.Div(first, second)
	mulResult := new(big.Int)
	mulResult.Mul(first, second)
	subResult := new(big.Int)
	subResult.Sub(first, second)
	fmt.Println("Сложение:", addResult, "\nДеление:", divResult, "\nУмножение:", mulResult, "\nВычитание:", subResult)
}
