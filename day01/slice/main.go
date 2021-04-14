package main

import "fmt"

func deleteSlice(s []int, pos int)  []int {
	a := append(s[:pos], s[pos+1:]...)
	return  a
}

func main() {
	//a := []int{0,1,2,3,4,5,6,7}
	//s := a[2:6]  // s := a[low:high]2,3,4,5
	//s[1] = 0
	//fmt.Printf("s:%v len(s):%v cap(s):%v\n", s, len(s), cap(s)) //2,0,4,5--6
	//a = append(a, 8)
	//fmt.Println(a)
	//l:= s[1:3] //3,4
	//fmt.Printf("l:%v len(s):%v cap(s):%v\n", l, len(l), cap(l))//

	//a := [...]int{1,3,4,5}
	//b := []int{1,3,4,5}
	//fmt.Printf("%T", a)
	//fmt.Printf("%T", b)

	//切片的复制
	//a := []int{1,2,3,4,5}
	//b := make([]int, 5)
	//fmt.Println(b)
	//copy(b, a)
	//b[0] = -1
	//fmt.Println(b, a)

	//切片删除一个元素
	//d := deleteSlice(a, 1)
	//fmt.Println(d)
	//
	//
	//var c = make([]string, 5, 10)
	//for i := 0; i < 10; i++ {
	//	c = append(c, fmt.Sprintf("%v", i))
	//}
	//fmt.Printf("%#v",c)
	//fmt.Println(c)
	//
	//var e = []int{3, 7, 8, 9, 1}
	//sort.Ints(e)
	//fmt.Println(e)

	//a := []int{0,1,2,3,4,5,6}
	//a = append(a, 8)
	//t := a
	//fmt.Println(a,cap(a))
	//a = append(a[:1], 9,10,11,12,13,14,15,16,17,18,19)
	//fmt.Println(a,cap(a))
	//fmt.Println(t,cap(t))

	//a := [5]int{1, 2, 3, 4, 5}
	//s := a[1:3]  // s := a[low:high]
	//fmt.Printf("s:%v len(s):%v cap(s):%v\n", s, len(s), cap(s))
	//s2 := s[3:4]  // 索引的上限是cap(s)而不是len(s)
	//fmt.Printf("s2:%v len(s2):%v cap(s2):%v\n", s2, len(s2), cap(s2))


	//a := []int{1, 2, 3, 4, 5}
	//t := a[1:3]
	//fmt.Printf("t:%v len(t):%v cap(t):%v\n", t, len(t), cap(t))

	//test3 := make([]int, 3, 10)
	//test4 := test3
	//test4[0] = 100
	//fmt.Println(test3,test4)
	/**
	append进行扩容了，如果新申请的容量是旧的容量的2倍，直接进行更改为新的容量，否则旧的容量如果小于1024,直接2倍。
	如果大于1024，则是新的容量=新的容量+新的容量/4
	 */

	//test4 = append(test4, 200)
	//test4[1] = 500
	//fmt.Println(test3,test4)
	//fmt.Println(cap(test3),cap(test4))
	//s1 := []int{0,1,2,3,4,5,6,7}
	//s2 := append(s1[:2], s1[3:]...)
	//fmt.Println(s1,s2,cap(s1),cap(s2))
	//s3:=append(s2, 1000,1000)
	//fmt.Println(s1,s2,s3)

	s1 := make([]int, 3, 3)
	fmt.Println(s1)

}
