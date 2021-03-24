/**
	查找一组数中只有唯一一个数组的元素
**/
package main

import "fmt"

func main() {
	var a []int = []int{3, 3, 9, 8, 7, 9, 8}
	var num int
	for i := 0; i < len(a); i++ {
		num ^= a[i]
	}
	fmt.Println(num)
}
