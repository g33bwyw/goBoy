package main

import "fmt"

func main() {
	var a = new(int)
	*a   = 100
	fmt.Println(*a)
	//b := 100
	//a = &b
	//*a = 200
	//fmt.Println(*a, b)

}
