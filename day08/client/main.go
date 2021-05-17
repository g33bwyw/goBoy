package main

import (
	"bufio"
	"fmt"
	"net"
	"os"
	"strings"
)

func main() {
	conn, err := net.Dial("tcp", "127.0.0.1:80")
	if err != nil {
		fmt.Println("client error", err)
		return
	}
	defer conn.Close()
	reader := bufio.NewReader(os.Stdin)
	for {
		str, _ := reader.ReadString('\n')
		str = strings.Trim(str, "\r\n")
		if strings.ToLower(str) == "q" {
			break
		}
		_, err := conn.Write([]byte(str))
		if err != nil {
			return
		}
		var buf [512]byte
		n, err := conn.Read(buf[:])
		if err != nil {
			fmt.Println("client read err ", err)
			return
		}
		fmt.Println(string(buf[:n]))

	}
}
