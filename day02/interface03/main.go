/**
该方法是实现接口嵌套
*/
package main

import "fmt"

type Singer interface {
	Say()
}

type Mover interface {
	move()
}

type Animaler interface {
	Singer
	Mover
}

type Animal struct {
	name string
}

func newAnimal(name string) *Animal {
	return &Animal{name: name}
}

func (a *Animal) say() {
	fmt.Printf("%s在叫\n", a.name)
}
func main() {
	dog := newAnimal("dog")
	dog.say()
}
