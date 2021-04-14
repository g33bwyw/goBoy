package main

import "fmt"

//type person struct {
//	name string
//}
func main() {
	//p := new(person)
	//p.name = "wangyw"
	//fmt.Printf("%v\n", p)
	//fmt.Printf("%+v\n", p)
	//fmt.Printf("%#v\n", p)
	//
	//n:=3
	//fmt.Printf("+%d%%\n", n)
	//str := "wang'ya"
	//fmt.Printf("%q\n", str)
	//fmt.Printf("%s\n", str)

	var input1, input2 int
	//fmt.Println("请输入参数1：")
	//fmt.Scan(&input1)
	//fmt.Println("请输入参数2：")
	//fmt.Scan(&input2)
	//fmt.Println(input1 + input2)

	//fmt.Scanf("%d-%d", &input1, &input2)
	//fmt.Println(input1 + input2)


	fmt.Scanln(&input1,&input2)
	fmt.Println(input1 + input2)
}

