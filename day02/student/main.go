package main

import (
	"fmt"
	"os"
)
type student struct {
	id int
	name string
}
var allStudent map[int]*student

func newStudent(id int, name string) *student{
	return &student{
		id:id,
		name:name,
	}
}
func searchStudent() {
	for i,v:=range allStudent {
		fmt.Printf("学号:%d,姓名：%s\n", i, v.name)
	}
}

func addStudent() {
	fmt.Println("请输入学生的学号")
	var (
		id int
		name string
	)
	fmt.Scanln(&id)
	fmt.Println("请输入学生的姓名")
	fmt.Scanln(&name)
	stu := newStudent(id,name)
	allStudent[id] = stu
	fmt.Println("学生添加成功")
}

func deleteStudent() {
	fmt.Println("请输入删除学生的学号")
	var (
		id int
	)
	delete(allStudent, id)
	fmt.Println("学生删除成功")
}
func main() {
	allStudent = make(map[int]*student)
	for{
		//欢迎菜单,以及菜单选项
		fmt.Printf("欢迎来到学生管理!\n" +
			"请选择菜单："+`
	1.查询学生信息 
	2.添加学生信息 
	3.删除学生信息 
	4.退出
`+"您选择的菜单：")
		var choice int
		fmt.Scanln(&choice)
		switch choice {
		case 1:
			searchStudent()
		case 2:
			addStudent()
		case 3:
			deleteStudent()
		case 4:
			os.Exit(1)
		default:
			fmt.Println("参数错误")
		}
	}
}
