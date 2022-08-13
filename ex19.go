package main

import (
	"fmt"
)

/*
	Разработать программу, которая переворачивает подаваемую на ход строку
	(например: «главрыба — абырвалг»). Символы могут быть unicode.
*/

func Reverse(str string) string {
	// перекиыдываем в руны
	r := []rune(str)
	// и меняем первый с последним, второй с пред..и тд
	for i, j := 0, len(r) - 1; i < j; i, j = i+1, j - 1 {
		r[i], r[j] = r[j], r[i]
	}
	return string(r)
}

func main() {
	str := "В чёрном цилиндре, в наряде старинном."
	fmt.Println(str)
	fmt.Println(Reverse(str))
}
