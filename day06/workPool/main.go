package main

import (
	"fmt"
	"sync"
	"time"
)

var wg sync.WaitGroup

func worker(id int, jobs <-chan int, results chan<- int) {
	for j := range jobs {
		//fmt.Printf("worker:%d start job:%d\n", id, j)
		//time.Sleep(time.Second)
		fmt.Printf("worker:%d end job:%d\n", id, j)
		results <- j * 2
		time.Sleep(time.Second)
		wg.Done()
	}
}

var s []int

func main() {
	jobs := make(chan int, 100)
	results := make(chan int, 100)

	wg.Add(13)
	// 开启3个goroutine
	for w := 1; w <= 3; w++ {
		go worker(w, jobs, results)
	}

	// 5个任务
	go func() {
		for j := 1; j <= 5; j++ {
			jobs <- j
			wg.Done()
		}
	}()

	// 输出结果
	go func() {
		for a := 1; a <= 5; a++ {
			s = append(s, <-results)
			//fmt.Println(<-results)
			time.Sleep(time.Second * 2)
			wg.Done()
		}
	}()
	wg.Wait()

	fmt.Println(s)
}
