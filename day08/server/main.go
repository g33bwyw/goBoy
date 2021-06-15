package main

import (
	"bufio"
	"fmt"
	"net"
	"time"
)

func process(conn net.Conn) {
	defer conn.Close()
	for {
		var buf [128]byte
		reader := bufio.NewReader(conn)
		n, err := reader.Read(buf[:])
		if err != nil {
			fmt.Println("read client err :", err)
			break
		}
		content := string(buf[:n])
		fmt.Println("client 信息", content)
		cur := time.Now().Format("2006-01-02 15:04:05")
		callBack := cur + " " + content
		conn.Write([]byte(callBack))
	}
}

func main() {
	listen, err := net.Listen("tcp", "127.0.0.1:3000") //监听端口
	if err != nil {
		fmt.Println("server error:", err)
		return
	}
	for {
		conn, err := listen.Accept() //建立连接
		if err != nil {
			fmt.Println("server error:", err)
			return
		}
		//处理数据
		go process(conn)
	}

}
