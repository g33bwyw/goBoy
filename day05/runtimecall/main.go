package main

import (
	"fmt"
	"runtime"
)

func log() {
	//主要用来获取函数，文件，行号
	pc, file, line, ok := runtime.Caller(1)

	fmt.Println(runtime.FuncForPC(pc).Name())
	fmt.Println(file)
	fmt.Println(line)
	fmt.Println(ok)
}

func main() {
	log()
}
