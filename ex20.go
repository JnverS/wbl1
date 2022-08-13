package main

import (
	"fmt"
	"strings"
)

/*
	Разработать программу, которая переворачивает слова в строке.
	Пример: «snow dog sun — sun dog snow».
*/

func Reverse(str string) string {
	// разбиваем на слова
	split := strings.Fields(str)
	//меняем местами слова в массиве
	for i, j := 0, len(split) - 1; i < j; i, j = i+1, j - 1 {
		split[i], split[j] = split[j], split[i]
	}
	return strings.Join(split, " ")
}

func main() {
	str := "В город на праздник путник очень спешил."
	fmt.Println(str)
	fmt.Println(Reverse(str))
}
