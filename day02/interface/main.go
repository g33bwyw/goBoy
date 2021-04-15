package main

import "fmt"

type Singer interface {
	say()
}

type Child struct {
}

type Work struct {
}

func (c Child) say() {
	fmt.Println("儿童唱儿歌")
}

func (w Work) say() {
	fmt.Println("工人唱共产国际歌")
}

func doSing(s Singer) {
	s.say()
}
func main() {
	p1 := Child{}
	doSing(p1)
	p2 := Work{}
	doSing(p2)
}
