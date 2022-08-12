package main

import "fmt"

/*
	Имеется последовательность строк - (cat, cat, dog, cat, tree) создать для нее
	собственное множество
*/

func main() {
	subsequence := []string{"cat", "cat", "dog", "cat", "tree"}

	subset := make(map[string]struct{})
	set := make([]string, 0)

	for _, v := range subsequence {
		if _, ok := subset[v]; !ok {
				subset[v] = struct{}{}
				set = append(set, v)
		}
	}
	fmt.Println(set)
}
