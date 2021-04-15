package main

import (
	"encoding/json"
	"fmt"
)

type Student struct {
	Id   int    `json:"id"`
	Name string `json:"name"`
}

type Class struct {
	Room    int        `json:"room"`
	Student []*Student `json:"student"`
}

/**
结构体的实例化与反实例化
*/
func main() {
	c := Class{
		Room:    1,
		Student: make([]*Student, 60),
	}
	for i := 0; i < 60; i++ {
		c.Student[i] = &Student{
			Id:   i,
			Name: fmt.Sprintf("stu%d", i),
		}
	}
	slice, _ := json.Marshal(c)
	str := string(slice)
	fmt.Println(str)

	slice2 := []byte(str)
	c2 := &Class{}
	json.Unmarshal(slice2, c2)
	fmt.Println(c2)
}
