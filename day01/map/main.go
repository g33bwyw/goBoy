package main

import (
	"fmt"
	"math/rand"
	"strings"
	"time"
)

//切片
func mapSlice() {
	a := make([]map[string]string, 3)
	a[0] = make(map[string]string)
	a[0]["安徽"] = "合肥"
	a[0]["上海"] = "徐汇"
	a[1] = map[string]string{"name": "wyw", "age": "18"}
	a[2] = a[1]
	a[2]["name"] = "test"
	fmt.Printf("%v", a)
}

func productStudent() {
	fmt.Println("------------------")
	rand.Seed(time.Now().UnixNano())
	student := make(map[string]int)
	for i := 0; i < 100; i++ {
		k := fmt.Sprintf("stu%d", i)
		student[k] = rand.Intn(100)
	}
	fmt.Println(student)
}

/**
统计单词数目
**/
func statWord(s string) map[string]int {
	sSlice := strings.Split(s, " ")
	statRes := make(map[string]int)
	for _, v := range sSlice {
		_, ok := statRes[v]
		if ok {
			statRes[v] = statRes[v] + 1
		} else {
			statRes[v] = 1
		}
	}
	return statRes
}

func main() {

	//map的声明
	//a := make(map[string]string, 1)
	//a["安徽"] = "合肥"
	//a["上海"] = "徐汇"
	//fmt.Println(a)
	//v, ok := a["安徽"]
	//if ok {
	//	delete(a, "安徽")
	//	fmt.Println("ok3", v)
	//} else {
	//	fmt.Println("error1")
	//}
	//fmt.Println(a)
	//mapSlice()
	//productStudent()
	//fmt.Println("统计单词结果")
	//statWord("how do you do")
	//
	//
	//
	type Map map[string][]int
	m := make(Map)
	s := []int{1, 2}
	s = append(s, 3)
	fmt.Printf("%p\n", s)
	m["q1mi"] = s
	fmt.Printf("%p\n", s)
	s = append(s[:1], s[2:]...)
	fmt.Println("容量：",cap(s))
	fmt.Printf("%p", s)
	fmt.Printf("%+v\n", s)
	fmt.Printf("%+v\n", m["q1mi"])
}
