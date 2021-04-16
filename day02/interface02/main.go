/**
类型嵌套,该方法是类型嵌套
*/
package main

import "fmt"

type WashMachine interface {
	dry()
	wash()
}

type dryer struct {
}

func (*dryer) dry() {
	fmt.Println("这是干洗功能")
}

type huadi struct {
	dryer
}

func main() {
	x := &huadi{}
	x.dry()
}
