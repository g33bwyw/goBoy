package main

import "sync"

type Counter interface {
	Inc()
	Load() int64
}

type CommonCounter struct {
	x int64
}

func (c CommonCounter)Inc() {
	c.x++
}

func (c CommonCounter) Load() int64{
	return c.x
}


type MutexCounter struct {
	x int64
	lock sync.Mutex
}

func (m MutexCounter)Inc() {
	
	c.x++
}

func (c CommonCounter) Load() int64{
	return c.x
}

func main() {

}
