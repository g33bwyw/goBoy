package main

import (
	"fmt"
	"time"
)

func main() {
	//加载时区
	local, _ := time.LoadLocation("Asia/Shanghai")
	timeObj := time.Now()
	//第一种格式化方式
	fmt.Printf("%d-%02d-%02d %02d:%02d:%02d \n",
		timeObj.Year(), timeObj.Month(), timeObj.Day(), timeObj.Hour(), timeObj.Minute(), timeObj.Second())
	//第二种格式化方式
	fmt.Println(timeObj.Format("2006-01-02 15:04:05"))
	//打印时间带时区
	fmt.Println(timeObj.Local())
	fmt.Println(timeObj.Location())
	//获取当前时间的时间戳
	fmt.Println(timeObj.Unix())
	//获取指定时间的时间戳
	obj, _ := time.ParseInLocation("2006-01-02 15:04:05", "2021-03-20 16:08:26", local)
	fmt.Println(obj.Unix())
	//时间定时器
	//t := time.Tick(time.Second * 3)
	//for i := range t {
	//	fmt.Println(i)
	//}

	//时间添加一个小时
	newTime := timeObj.Add(time.Minute * 30)
	fmt.Println(newTime.Unix())
	fmt.Println(newTime.Format("2006-01-02 15:04:05"))

}
