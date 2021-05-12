package main

import (
	"fmt"
	"sync"
	"time"
)

var x int
var wg sync.WaitGroup
var lock sync.Mutex
var rwlock sync.RWMutex
func write() {
	lock.Lock()
	//rwlock.Lock()
	x++
	time.Sleep(10 * time.Millisecond)
	//rwlock.Unlock()
	lock.Unlock()
	wg.Done()
}

func read() {
	lock.Lock()
	//rwlock.RLock()
	fmt.Sprintf("%d", x)
	time.Sleep(time.Millisecond)
	//rwlock.RUnlock()
	lock.Unlock()
	wg.Done()
}

func main() {
	start := time.Now()
	for i:=1;i<=1000;i++ {
		wg.Add(1)
		go read()
	}
	for i:=1;i<=10;i++ {
		wg.Add(1)
		go write()
	}
	wg.Wait()
	fmt.Println(time.Now().Sub(start))

}
