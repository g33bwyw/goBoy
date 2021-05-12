package main

import (
	"fmt"
	"sync"
)

var x int
var wg sync.WaitGroup
var lock sync.Mutex
func readX() {
	for i:=0;i<10000;i++ {
		lock.Lock()
		x += 1
		lock.Unlock()
	}
	wg.Done()
}
func main() {
	wg.Add(2)
	go readX()
	go readX()
	wg.Wait()
	fmt.Println(x)
}
