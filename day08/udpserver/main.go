package main

import (
	"fmt"
	"net"
	"strings"
)

func main() {
	//1.开启服务端口
	socket, err := net.ListenUDP("udp", &net.UDPAddr{
		IP:   net.IPv4(127, 0, 0, 1),
		Port: 3000,
	})
	if err != nil {
		fmt.Println("udp server err", err)
		return
	}
	defer func() {
		err := socket.Close()
		fmt.Println(err)
	}()
	//2.读取客户端发送过来的数据
	for {
		buf := make([]byte, 1024)
		n, clientAddr, err := socket.ReadFromUDP(buf)
		if err != nil {
			fmt.Println("udp server read err", err)
			break
		}
		content := strings.ToUpper(string(buf[:n]))
		fmt.Println("客户端发送的数据：", content)
		_, err2 := socket.WriteToUDP([]byte(content), clientAddr)
		if err2 != nil {
			fmt.Println("udp server write err", err)
			break
		}
	}
}
