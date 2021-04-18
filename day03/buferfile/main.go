package main

import (
	"bufio"
	"fmt"
	"io"
	"os"
)

/**
buffer读取和写入文件
mod 1，通过读取分割符号读取文件内容
	2.通过字节切片来读取文件内容
*/
func readFile(filename string, mod int) interface{} {
	file, err := os.Open(filename)
	if err == io.EOF {
		fmt.Println("文件打开失败")
		return false
	}
	defer file.Close()
	var fileContent string
	r := bufio.NewReader(file)

	if mod == 1 {
		for {
			content, err := r.ReadString('\n')
			if err == io.EOF {
				if len(content) != 0 {
					fileContent += content
				}
				break
			}
			fileContent += content
		}
	} else {
		content := make([]byte, 1024)
		for {
			n, err := r.Read(content)
			if err == io.EOF {
				break
			}
			fileContent += string(content[:n])
		}
	}

	return fileContent

}

/**
以字符串的形式写入到文件中
 */
func writeFile(filename string, content string) interface{}{
	file, err := os.OpenFile(filename, os.O_CREATE|os.O_TRUNC|os.O_RDWR, 0666)
	if err == io.EOF {
		fmt.Println("文件打开失败")
		return false
	}
	defer file.Close()
	w:=bufio.NewWriter(file)
	w.WriteString(content)
	w.Flush()
	return true
}

/**
文件拷贝
 */
func copyFile(sourceFile,distFile string) (interface{}) {
	source, err := os.Open(sourceFile)
	if err == io.EOF {
		fmt.Println("读取文件打开失败")
		return false
	}
	dist, err := os.OpenFile(distFile,  os.O_CREATE|os.O_TRUNC|os.O_RDWR, 0666)
	if err == io.EOF {
		fmt.Println("写入文件打开失败")
		return false
	}
	defer source.Close()
	defer dist.Close()
	r := bufio.NewReader(source)
	w := bufio.NewWriter(dist)
	content := make([]byte,1024)
	for {
		n,err:=r.Read(content)
		if err == io.EOF {
			break
		}
		w.Write(content[:n])
	}
	w.Flush()
	return true
}

func main() {
	//fmt.Println(readFile("E:/go/src/goBoy/day03/docker-compose.yml", 2))

	//s:=readFile("E:/go/src/goBoy/day03/docker-compose.yml", 2)
	//writeFile("./c.txt", fmt.Sprintf("%s", s))

	copyFile("E:/go/src/goBoy/day03/docker-compose.yml", "./d.txt")
}
