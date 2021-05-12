package main

import (
	"fmt"
	"sync"
)

type singleton struct {

}
var loadOnce sync.Once

func getInstance() *singleton {
	var instance *singleton
	loadOnce.Do(func() {
		instance = &singleton{}
	})
	return instance
}
func main() {
	var a int
	fmt.Printf("%p", &a)
}
