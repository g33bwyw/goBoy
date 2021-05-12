package main

import (
	"fmt"
	"sync"
)
var wg sync.WaitGroup
func printC(c chan int) {
	//for {
	//	v,ok := <-c
	//	if !ok {
	//		break
	//	}
	//	fmt.Println(v)
	//}
	for v := range c{
		fmt.Println(v)
	}
	wg.Done()
}
func main() {
	c := make(chan int)
	wg.Add(1)
	go printC(c)
	for i:=1;i<=10;i++{
		c<-i
	}
	close(c)
	wg.Wait()
}
