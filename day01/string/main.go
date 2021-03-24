package main

import (
	"fmt"
	"strings"
)
func statChineseTotal(s string) int {
	a1 := len(s)
	a2 := len([]rune(s))
	return (a1-a2)/2

}

/**
字符串的相关使用
 */
func main() {
	var str string = "hello word china 中国"
	//打印字符串长度
	fmt.Println(len(str))
	//字符串的切割
	strArr := strings.Split(str," ")
	fmt.Println(strArr)
	//字符串的连接
	fmt.Println(strings.Join(strArr, "-"))
	//字符串的替换(最后一个参数表示替换多少个，-1全部替换)
	fmt.Println(strings.Replace(str," ", "--", -1))
	//查找字符串出现的位置
	fmt.Println(strings.Index(str, "ch"))
	//查找字符串是否存在其他元素
	fmt.Println(strings.Contains(str,"ch"))
	//查找字符串是否包含前缀
	fmt.Println(strings.HasPrefix(str, "he"))
	//查找字符串是否包含后缀
	fmt.Println(strings.HasSuffix(str,"dd"))

	//字符串循环打印(中文字符会做成三个字节打印)
	for i:=0;i<len(str);i++ {
		fmt.Printf("%v(%c) ", str[i], str[i])
	}
	fmt.Println()
	//range 直接打印出来字符串
	for _,v := range str {
		fmt.Printf(" %+v(%c) ", v, v)
	}
	fmt.Println()
	//中文字符串添加字符串，字符串底层使用rune
	runeArr := []rune(str)
	runeArr[2] = '美'
	fmt.Println(string(runeArr))

	//字符串底层使用byte
	str3 := "test"
	strArr2 := []byte(str3)
	strArr2[0] = 'T'
	fmt.Println(string(strArr2))

	fmt.Println("中文字数:",statChineseTotal("hello沙河小王子"))
}
