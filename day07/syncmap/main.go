package main

import (
	"fmt"
	"strconv"
	"sync"
)

var m = map[string]int{}
var wg sync.WaitGroup
var m1 sync.Map

func get(k string) int{
	val, ok := m1.Load(k)
	if !ok {
		return 0
	}
	z, ok := val.(int)
	if !ok {
		return 0
	}
	return z
}

func set(k string, v int) {
	m1.Store(k,v)
	wg.Done()
}

func main() {

	for i:=1;i<=2;i++ {
		wg.Add(1)
		go set(strconv.Itoa(i), i)
	}
	wg.Wait()
	fmt.Println(get("1"))
}
