package main

import "fmt"

func main() {
	const (
		a = iota
		_
		b = iota
		c = 100
		g
		d = iota
		e
		f
	)
	fmt.Println(a)
	fmt.Println(b)
	fmt.Println(c)
	fmt.Println(d)
	fmt.Println(e)
	fmt.Println(f)
	fmt.Println(g)
}
