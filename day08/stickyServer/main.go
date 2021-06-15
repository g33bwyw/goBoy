package main

import (
	"bufio"
	"fmt"
	"goBoy/day08/proto"
	"io"
	"net"
)

func process(conn net.Conn) {
	//defer func() {
	//	err := conn.Close()
	//	fmt.Println("sticky server close err：", err)
	//	return
	//}()

	defer conn.Close()

	reader := bufio.NewReader(conn)
	//buf := make([]byte, 1024)
	for {
		//n, err := reader.Read(buf)
		//if err == io.EOF {
		//	break
		//}
		//if err != nil {
		//	fmt.Println("read from client failed, err:", err)
		//	break
		//}
		//receiveStr := string(buf[:n])
		receiveStr, err := proto.Decode(reader)
		if err == io.EOF {
			return
		}
		if err != nil {
			fmt.Println("decode msg failed, err:", err)
			return
		}
		fmt.Println("收到client发来的数据：", receiveStr)
	}
}

func main() {
	listen, err := net.Listen("tcp", "127.0.0.1:30000")
	if err != nil {
		fmt.Println("sticky server listen err：", err)
		return
	}
	defer listen.Close()
	for {
		conn, acceptErr := listen.Accept()
		if acceptErr != nil {
			fmt.Println("sticky server accept err：", acceptErr)
			continue
		}
		go process(conn)
	}

}
