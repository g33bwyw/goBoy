package main

import (
	"fmt"
	"goBoy/day08/proto"
	"net"
	"time"
)

func main() {
	conn, err := net.Dial("tcp", "127.0.0.1:30000")
	if err != nil {
		fmt.Println("client err:", err)
		return
	}
	defer conn.Close()
	for i := 0; i < 20; i++ {
		message := "Hello, Hello. How are you?"
		//conn.Write([]byte(message))
		buf, err := proto.Encode(message)
		if err != nil {
			fmt.Println("encode msg failed, err:", err)
			return
		}
		conn.Write(buf)
	}
	time.Sleep(time.Second * 60)
}
