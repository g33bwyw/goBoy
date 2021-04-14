package main

import fmt2 "fmt"

func main() {
	//a := [3]int{1,2,3}
	//b := a
	//b[1] = 5
	//fmt.Println(a)

	a := [...]int{1,2,4}
	b := a
	b[2] = 5
	fmt2.Printf("%T", a)
	fmt2.Printf("%T", b)
	fmt2.Println(a)
	fmt2.Println(b)
}
