package main

import (
	"flag"
	"fmt"
)

func main() {
	//var name string
	//var age int
	//flag.IntVar(&age,"age", 0, "年龄")
	//flag.StringVar(&name,"name","","姓名")

	name := flag.String("name", "", "姓名")
	age := flag.Int("age", 0, "年龄")
	flag.Parse()

	fmt.Println(flag.Args())
	//命令行参数之后的其他参数
	fmt.Println(flag.NArg())
	//使用命令行参数
	fmt.Println(flag.NFlag())
	fmt.Println(*name, *age)

}
