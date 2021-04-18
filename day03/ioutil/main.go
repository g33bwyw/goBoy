package main

import (
	"fmt"
	"io/ioutil"
)

func main() {
	source := "E:/go/src/goBoy/day03/docker-compose.yml"
	content, _ := ioutil.ReadFile(source)
	fmt.Println(string(content))

	ioutil.WriteFile("./e.txt", content, 0666)
}
