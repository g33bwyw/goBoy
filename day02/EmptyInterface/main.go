package main

import fmt1 "fmt"

func show(x interface{}) {
	fmt1.Println(x)
}

func main() {
	var e interface{}
	a := 1
	e = a
	fmt1.Printf("%T----%#v\n", e, e)
	b := "hello word"
	e = b
	fmt1.Printf("%T----%#v\n", e, e)

	m := make(map[string]interface{})
	m["a"] = "lihao"
	m["b"] = 123
	fmt1.Println(m)

}
