package main

import "fmt"

type Teacher interface {
	teach(str string)
}

type Student interface {
	study(str string)
	play(str string)
}

type People struct {
	age     int
	student Student
	teacher Teacher
	name    string
}

func (p *People) study(str string) {
	fmt.Printf("%s的责任：%s\r\n", p.name, str)
}
func (p *People) play(str string) {
	fmt.Printf("%s的爱好：%s\r\n", p.name, str)
}
func main() {
	var x interface{}
	s := "Hello 沙河"
	x = s
	fmt.Printf("type:%T value:%v\n", x, x)
	i := 100
	x = i
	fmt.Printf("type:%T value:%v\n", x, x)
	b := true
	x = b
	fmt.Printf("type:%T value:%v\n", x, x)
}
