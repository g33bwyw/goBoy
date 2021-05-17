package main

import (
	"bufio"
	"fmt"
	"net"
	"os"
)

func main() {
	//1.和服务端建立连接
	client, err := net.DialUDP("udp", nil, &net.UDPAddr{
		IP:   net.IPv4(127, 0, 0, 1),
		Port: 3000,
	})
	if err != nil {
		fmt.Println("client connect error:", err)
	}
	defer func() {
		err := client.Close()
		fmt.Println(err)
	}()
	for {
		content := make([]byte, 1024)
		fmt.Println("请输入英文字母：")
		reader := bufio.NewReader(os.Stdin)
		n, err := reader.Read(content)
		if err != nil {
			fmt.Println("client read error：", err)
			break
		}
		_, writeErr := client.Write(content[:n])
		if writeErr != nil {
			fmt.Println("client udp write error：", err)
			break
		}
		data := make([]byte, 1024)
		n, addr, readErr := client.ReadFromUDP(data)
		if readErr != nil {
			fmt.Println("server udp callback error：", err)
			break
		}
		fmt.Println(addr)
		fmt.Println("callback data:", string(data[:n]))
	}

}
