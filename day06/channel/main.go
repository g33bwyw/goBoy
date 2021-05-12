package main

import (
	"fmt"
	"math/rand"
	"strconv"
	"sync"
	"time"
)

var wg sync.WaitGroup
var once sync.Once
var num int

//循环生成一个随机数
func makeSeedNum(jobChan chan int64) {
	rand.Seed(time.Now().UnixNano())
	for i := 0; i < 50; i++ {
		jobChan <- rand.Int63n(100000)
	}
	return
}
func sumSeedNum(seedNum int64, resultChan chan string) {
	var sum int64
	seedStr := strconv.FormatInt(seedNum, 10)
	seedSlice := []rune(seedStr)
	for _, v := range seedSlice {
		num, _ := strconv.ParseInt(string(v), 10, 64)
		sum += num
	}
	resultChan <- fmt.Sprintf("值：%d,总和：%d", seedNum, sum)
	num++
	if (num == 50) {
		close(resultChan)
	}
	//t := func() {
	//	close(resultChan)
	//}
	//once.Do(t)
	return
}

func main() {
	jobChan := make(chan int64, 20)
	resultChan := make(chan string, 50)
	go makeSeedNum(jobChan)
	go func() {
		for seedNum := range jobChan {
			sumSeedNum(seedNum, resultChan)
		}
	}()

	for {
		r, ok := <-resultChan
		if !ok {
			break
		}
		fmt.Println(r)


	}

	//次写法打印出来的是字符串
	//for _,v := range <-resultChan{
	//	fmt.Println(string(v))
	//}

}
