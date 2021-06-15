package main

import (
	"bufio"
	"log"
	"net"
	"time"
)

func main() {
	listener, err := net.Listen("tcp", "127.0.0.1:3000")
	if err != nil {
		log.Fatalf("建立tcp server失败:%v\n", err)
	}
	defer listener.Close()

	//
	for {
		// 新建一个连接
		conn, err := listener.Accept()
		if err != nil {
			log.Printf("accept err: %v\n", err)
		}
		go worker(conn)
		log.Println("连接已退出...", conn.LocalAddr())
	}
}

func worker(conn net.Conn) {

	defer conn.Close()
	read := make(chan string)
	go func() {
		for {
			reader := bufio.NewReader(conn)
			var buf [128]byte
			n, err := reader.Read(buf[:]) // 读取数据
			if err != nil {
				log.Println("read from client failed, err:", err)
				break
			}
			recvStr := string(buf[:n])
			log.Println("收到client端发来的数据：", recvStr)
			read <- recvStr
		}
	}()
	for {
		select {
		case <-time.After(time.Second * 3):
			log.Println("长时间未收到消息,退出...")
			conn.Write([]byte("timeout..."))
			conn.Close()
			return
		case recvStr := <-read:
			conn.Write([]byte(recvStr)) // 发送数据
		}
	}
}
