package main

import "fmt"

const (
	a = iota
	_
	b
	c
)

func main() {
	fmt.Println(a)
	fmt.Println(b)
	fmt.Println(c)
}
