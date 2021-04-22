package main

import (
	"fmt"
	"runtime"
)

func log() {
	pc, file, line, ok := runtime.Caller(1)

	fmt.Println(runtime.FuncForPC(pc).Name())
	fmt.Println(file)
	fmt.Println(line)
	fmt.Println(ok)
}

func main() {
	log()
}
