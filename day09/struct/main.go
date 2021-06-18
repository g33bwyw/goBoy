package main

import (
	"encoding/json"
	"fmt"
)

//Student 学生
type Student struct {
	ID     int
	Gender string
	Name   string
}

//Class 班级
type Class struct {
	Title    string     `json:"title"`
	Students []*Student `json:"students"`
}

func main() {
	c := &Class{
		Title:    "101",
		Students: make([]*Student, 0),
	}
	for i := 1; i <= 20; i++ {
		sex := "女"
		if i/2 == 0 {
			sex = "男"
		}
		s := &Student{
			ID:     i,
			Gender: sex,
			Name:   fmt.Sprintf("stu%d", i),
		}
		c.Students = append(c.Students, s)
	}
	jsonInfo, _ := json.Marshal(c)
	jsonStr := string(jsonInfo)
	fmt.Println("json信息：", jsonStr)
	c1 := &Class{}
	json.Unmarshal(jsonInfo, c1)
	fmt.Printf("json结构体%v", c1)

}
