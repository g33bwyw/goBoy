# VScode安装

```
ctrl + shift + P
输入 go:install/update tools 安装go相关的提示包

设置自动保存

设置snippets
{
    "println":{
    "prefix": "pln",
    "body":"fmt.Println($0)",
    "description": "println"
    },
    "printf":{
    "prefix": "plf",
    "body": "fmt.Printf(\"$0\")",
    "description": "printf"
    }
}


```
## 编译


#  常量

const 关键词

iota使用：遇到const变为0,每增加一行值+1，iota初始值为0

``` 
1.初始化多个值
const (
	a = 100
	b
	c
)
结果： 100， 100， 100

2.iota用法
const (
	a = iota
	_
	b = iota
	c = 100
	g
	d = iota
	e
	f
)
结果：a = 0, b=2,c=100,g=100,d=5,e=6,f=7
```



# 基础运算符

``` 
/**
	在一个数组中存在唯一一个数字的值，其他数字个数全都是偶数
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
```



#  字符串相关

``` 
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

```

