package main

import (
	"fmt"
	"goBoy/day08/proto"
	"net"
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
		buf, _ := proto.Encode(message)
		conn.Write(buf)
	}

}
