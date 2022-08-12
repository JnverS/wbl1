package main

/*
	Реализовать пересечение двух неупорядоченных множеств
*/

import "fmt"

func main() {
	first := []string{"zero", "1", "2", "3", "4", "5"}
	second := []string{"1", "5", "0", "zero", "8"}

	fmt.Println(intersect(first, second))
}


func intersect(first, second []string) []string {
	m := make(map[string]int)
	nn := make([]string, 0)

	for _, v := range first {
		m[v]++
	}

	for _, v := range second {
		times, _ := m[v]
		if times > 0 {
			m[v]--
			nn = append(nn, v)
		}
	}
	return nn
}
