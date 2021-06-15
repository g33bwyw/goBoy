package main

import (
	"log"
	"os"
)

const (
	a = 1 << iota
	b
	c
	d
	e
	f
	g = a | b
)

func main() {
	file, _ := os.OpenFile("./test.txt", os.O_CREATE|os.O_APPEND|os.O_RDWR, 0777)
	log.SetOutput(file)
	log.SetFlags(log.Llongfile | log.Lmicroseconds | log.Ldate)
	log.SetPrefix("[prod]")
	log.Println(a, b, c, d, e, f, g)
}
