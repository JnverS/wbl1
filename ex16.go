package main

/*
	Реализовать быструю сортировку массива (quicksort) встроенными методами
	языка.
*/

import (
	"fmt"
	"math/rand"
	"sort"
	"time"
)

// создаем рандомный массив
func createarr(size int) []int {
	arr := make([]int, size, size)
	// рандомизация генератора случайных чисел
	rand.Seed(time.Now().UnixNano())
	// чтобы были отрицательные числа
	for i := 0; i < size; i++ {
		arr[i] = rand.Intn(999) - rand.Intn(999)
	}
	return arr
}

func partition (arr []int) int {
	// выбираем крайние токи
	left, right := 0, len(arr)-1
	// выбираем опорную точку
	mid := arr[len(arr)/2]

	// проходим массив: элементы которые меньше слева о.т, больше - справа
	for {
		for arr[left] < mid {
			left++
		}
		for arr[right] > mid {
			right--
		}
		if left >= right {
			return right
		}
		arr[left], arr[right] = arr[right], arr[left]
	}
}

func quicksort(arr []int) []int {
	if len(arr) < 2 {
		return arr
	}

	split := partition(arr)
	// повторяем для каждой части
	quicksort(arr[:split])
	quicksort(arr[split:])
	return arr
}

func main() {
	arr := createarr(15)
	fmt.Println(arr)
	fmt.Println(quicksort(arr))

	// можно еще использовать пакет sort
	arr2 := createarr(10)
	fmt.Println(arr2)
	sort.Ints(arr2)
	fmt.Println(arr2)
}
