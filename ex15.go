package main

import (
	"fmt"
	"strings"
)

/*
	К каким негативным последствиям может привести данный фрагмент кода, и как
	это исправить? Приведите корректный пример реализации.
	var justString string
	func someFunc() {
	v := createHugeString(1 << 10)
	justString = v[:100]
	}
	func main() {
	someFunc()
	}
*/
// глобальная переменная = доступ из любого места программы
// если при создании строки использовать символы unicode,
// то строка скопируется некорректно, необходимо добавить преобразование в тип rune

func createHugeString(num int) string{
	return strings.Repeat("o", num)
}

func someFunc() {

	v := createHugeString(1 << 10) // 10 000 000 000
	r := []rune(v)
	str := string(r[:100])
	fmt.Println(str)
}

func main() {
	someFunc()
}
