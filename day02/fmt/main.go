package main

import "fmt"

//type person struct {
//	name string
//}
func main() {
	//p := new(person)
	//p.name = "wangyw"
	//fmt.printf("%v\n", p)
	//fmt.printf("%+v\n", p)
	//fmt.printf("%#v\n", p)
	//
	//n:=3
	//fmt.printf("+%d%%\n", n)
	//str := "wang'ya"
	//fmt.printf("%q\n", str)
	//fmt.printf("%s\n", str)

	//var input1, input2 int
	//fmt.Println("请输入参数1：")
	//fmt.Scan(&input1)
	//fmt.Println("请输入参数2：")
	//fmt.Scan(&input2)
	//fmt.Println(input1 + input2)

	//fmt.Scanf("%d-%d", &input1, &input2)
	//fmt.Println(input1 + input2)


	//fmt.Scanln(&input1,&input2)
	//fmt.Println(input1 + input2)

	s := "小王子11112ee"

	a:=fmt.Sprintf("%s", s)
	b:=fmt.Sprintf("%5s", s)
	c:=fmt.Sprintf("%-5s", s)
	d:=fmt.Sprintf("%5.7s", s)
	e:=fmt.Sprintf("%-5.7s", s)
	f:=fmt.Sprintf("%5.2s", s)
	g:=fmt.Sprintf("%05s", s)

	fmt.Printf("%#v\n",a)
	fmt.Printf("%#v\n",b)
	fmt.Printf("%#v\n",c)
	fmt.Printf("%#v\n",d)
	fmt.Printf("%#v\n",e)
	fmt.Printf("%#v\n",f)
	fmt.Printf("%#v\n",g)
}

