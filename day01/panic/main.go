package main

import "fmt"

func funcA() {
	fmt.Println("this is A")
}

func funcB() {
	defer func() {
		err := recover()
		if err != nil {
			fmt.Println("this is recover")
		}
	}()
	panic("this is B")
}


func funcC() {
	fmt.Println("this is C")
}



func main() {
	funcA()
	funcB()
	funcC()
}
