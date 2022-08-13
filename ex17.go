package main

import (
	"fmt"
	"sort"
)

/*
	Реализовать бинарный поиск встроенными методами языка.
*/

func BinarySearch(arr []int, x int) int {
	// сортируем массив
	sort.Ints(arr)
	start := 0
	end := len(arr) - 1
	for start <= end {
		// находим середину
		mid := start + (end - start)/2
		//сравниваем средний элемент с тем, что ищем
		if arr[mid] == x {
			return mid
		}
		// если не нашли проверяем в какой половине искать наш элемент
		if arr[mid] > x {
			end = mid - 1
		} else {
			start = mid + 1
		}
	}
	// если совсем не нашли вернем -1
	return -1
}

func main() {
	arr := []int{1,2,3,4,5,6,7,8,9}
	x := 0

	index := BinarySearch(arr, x)
	if index == -1 {
		fmt.Println("Result BinarySearch: Not found")
	} else {
		fmt.Println("Result BinarySearch:", index)
	}
	// поиск с использованием пакета sort
	i := sort.Search(len(arr), func(i int) bool {return arr[i] >= x})
	if i < len(arr) && arr[i] == x {
		fmt.Println("Result sort.Search:", i)
	} else {
		fmt.Println("Result sort.Search: Not found")
	}
}
