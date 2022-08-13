package main

import "fmt"

/*
	Удалить i-ый элемент из слайса
*/

func main() {
	var i int
	a := []string{"1", "2", "3", "4", "5", "6"}
	fmt.Println("Введите номер элемента для удаления:")
	fmt.Scan(&i)

	if i > 0 && i < len(a) {
	// вариант с перемещением последнего элемента в i индекс
	a[i] = a[len(a)-1]
	a[len(a)-1] = ""
	a = a[:len(a)-1]
	fmt.Println(a)
	// вариант с сохранением порядка
	// сдвигаем влево на один индекс
	if i < len(a) {
		copy(a[i:], a[i+1:])
		// удаляем последний элемент
		a[len(a)-1] = ""
		// усекаем срез
		a = a[:len(a)-1]
		fmt.Println(a)
	}
	} else {
		fmt.Println("Неверный индекс")
	}
}
