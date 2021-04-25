package main

import (
	"fmt"
	"math/rand"
	"sync"
	"time"
)

var wg sync.WaitGroup

func f1(i int) {
	defer wg.Done()
	time.Sleep(time.Second * time.Duration(rand.Intn(3)))
	fmt.Println(i)
}

func f2() {
	defer wg.Done()
	fmt.Println("f2")
}
func f3() {
	defer wg.Done()
	fmt.Println("f3")
}
func main() {
	//rand.Seed(time.Now().UnixNano())
	//for i:=0;i<10;i++ {
	//	wg.Add(1)
	//	go f1(i)
	//}
	//wg.Wait()

	wg.Add(2)
	go f2()
	go f3()
	wg.Wait()

}
