package main

import (
	"fmt"
	"math/rand"
	"strconv"
	"sync"
	"time"
)

var wg sync.WaitGroup

//循环生成一个随机数
func makeSeedNum(jobChan chan int64) {
	defer wg.Done()
	rand.Seed(time.Now().UnixNano())
	for i := 0; i < 20; i++ {
		jobChan <- rand.Int63n(100000)
	}
	close(jobChan)
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
	return
}

func main() {
	jobChan := make(chan int64, 20)
	resultChan := make(chan string, 20)
	wg.Add(25)
	go makeSeedNum(jobChan)
	for i := 0; i < 24; i++ {
		go func() {
			defer wg.Done()
			for seedNum := range jobChan {
				sumSeedNum(seedNum, resultChan)
			}
		}()
	}
	wg.Wait()
	close(resultChan)
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
