package main

import (
	"fmt"
	"strings"
)

type person struct {
	name string
	age byte
	fav []string
}

//实例化函数,默认以new 开头
func newPerson(name string, age byte) *person{
	return &person{
		name: name,
		age: age,
	}
}

//设置用户爱好
func (p *person)setFav(fav []string) {
	p.fav = fav
}

//打印用户信息
func (p *person)printPerson() {
	if len(p.fav) == 0 {
		fmt.Printf("用户：%s,年龄：%c,无爱好\n", p.name, p.age)
	} else {
		fmt.Printf("用户：%s,年龄：%c,爱好：%s\n", p.name, p.age, strings.Join(p.fav, ","))
	}

}

type myint int
func main() {
	p := newPerson("wangyw", 'M')
	p.setFav([]string{"basketball","football"})
	p.printPerson()

	a := myint(10)
	fmt.Printf("%d", a)
}