package main

/*
	Реализовать конкурентную запись данных в map.
*/

import (
	"fmt"
	"sync"
)

type Map struct {
	rmx sync.RWMutex
	m map[int]string
}

func NewMap() *Map {
	return &Map {
		m: make(map[int]string),
	}
}

func (m *Map) Write(key int, value string) {
	m.rmx.Lock()
	m.m[key] = value
	m.rmx.Unlock()
}

func main() {
	wg := sync.WaitGroup{}
	m := NewMap()

	wg.Add(5)

	for i:= 0; i < 5;i++ {
		k := i
		go func() {
			defer wg.Done()
			m.Write(k, "hello")
		}()
	}
	wg.Wait()
	fmt.Println(m.m)
}
