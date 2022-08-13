package main

import (
	"fmt"
	"strings"
)

/*
	Разработать программу, которая проверяет, что все символы в строке
	уникальные (true — если уникальные, false etc). Функция проверки должна быть
	регистронезависимой.
	Например:
	abcd — true
	abCdefAaf — false
	aabcd — false
*/

func ChkUniq(s string) bool {
	str := []rune(strings.ToLower(s))
	buffer := make(map[rune]bool)

	for _, v := range str {
		if _, ok := buffer[v]; ok {
			return false
		}
		buffer[v] = true
	}
	return true
}

func main() {
	str := "see you"
	if ChkUniq(str) {
		fmt.Println("Все символы уникальные")
	} else {
		fmt.Println("Есть повторяющиеся символы")
	}
}
