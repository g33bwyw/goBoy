package main

import (
	"fmt"
	"unsafe"
)

type Person struct {
	name string
	age int
	sex byte
}
func main01()  {
	//初始化
	//var p Person
	//p.name = "wangyw"
	//p.age = 32
	//p.sex = '1'
	//fmt.Printf("%p\n", &p)
	//fmt.Printf("%#v", p)

	//p := Person{"wangyw", 32, '1'}
	//fmt.Printf("%#v\n", p)

	//创建指针变量
	//p := new(Person)
	//fmt.Printf("%#v\n", p)

	p := &Person{"wangyw", 32, '1'}
	fmt.Printf("%#v\n", p)
	fmt.Println(unsafe.Sizeof(p))



}
