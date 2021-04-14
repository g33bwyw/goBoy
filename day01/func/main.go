package main

import "fmt"

func add(i,v int) int{
	return i + v
}
func min(i,v int) int {
	return i-v
}
func simpleCalc(m string) func(i,v int) int {
	switch m {
	case "add" :
		return add
	case "min":
		return min
	}
	return nil
}

func f1() int {
	x := 5
	defer func() {
		x++
	}()
	return x
}

func f2() (x int) {
	defer func() {
		x++
	}()
	return 5
}

func f3() (y int) {
	x := 5
	defer func() {
		x++
	}()
	return x
}
func f4() (x int) {
	defer func(x int) {
		x++
	}(x)
	return 5
}

func calc(index string, a, b int) int {
	ret := a + b
	fmt.Println(index, a, b, ret)
	return ret
}

func main() {
	//定义函数作为参数
	//type calc func(m string) func(int, int) int
	//var a calc
	//a = simpleCalc
	//doAdd := a("add")
	//fmt.Println(doAdd(1,3))
	//doMin := a("min")
	//fmt.Println(doMin(7,2))

	//a := []int{0, 1, 2, 3, 4, 5, 7}
	//fmt.Println(moreArgs(a[2:]...))
	//fmt.Println(f1())
	//fmt.Println(f2())
	//fmt.Println(f3())
	//fmt.Println(f4())

	a := 1
	b := 2
	defer calc("1", a, calc("10", a, b))
	a = 0
	defer calc("2", a, calc("20", a, b))
	b = 1
}