/**
文件file方式读取以及写入
*/
package main

import (
	"fmt"
	"io"
	"os"
)

/**
文件读取
*/
func readFile(filename string) interface{} {
	file, err := os.Open(filename)
	if err != nil {
		fmt.Println("读取的文件打开失败")
		return false
	}
	defer file.Close()
	content := make([]byte, 1024)
	fileContent := make([]byte, 0)
	for {
		n, err := file.Read(content)
		fileContent = append(fileContent, content[:n]...)
		if err == io.EOF {
			break
		}

	}
	return string(fileContent)
}

/**
文件写入
*/
func writeNewFile(filename string, content string) interface{} {
	file, err := os.OpenFile(filename, os.O_CREATE|os.O_TRUNC|os.O_WRONLY, 0666) //os.O_TRUNC擦除重写os.o_append追加
	if err != nil {
		fmt.Println("读取的文件打开失败")
		return false
	}
	defer file.Close()
	file.WriteString(content)
	return true
}

/**
文件复制
*/
func copyFile(sourceFile, distFile string) interface{} {
	source, err := os.Open(sourceFile)
	if err != nil {
		fmt.Println("读取的文件打开失败")
		return false
	}
	dist, err := os.OpenFile(distFile, os.O_CREATE|os.O_TRUNC|os.O_WRONLY, 0666)
	if err != nil {
		fmt.Println("创建文件失败")
		return false
	}
	defer source.Close()
	defer dist.Close()
	content := make([]byte, 1024)
	for {
		n, err := source.Read(content)
		if err == io.EOF {
			break
		}
		dist.Write(content[:n])
	}

	return true
}

/**
每间隔5秒打印一次日志信息
*/
func printLog(filename string, pos *int64) string {
	file, err := os.Open(filename)
	if err != nil {
		fmt.Println("读取的文件打开失败")
		return ""
	}
	defer file.Close()
	//获取文件末尾的位置
	//n,err := file.Seek(pos, 1024)
	content := make([]byte, 1024)
	n, _ := file.ReadAt(content, *pos)
	fmt.Println(string(content[:n]))
	return ""

}
func main() {
	s := readFile("E:/go/src/goBoy/day03/docker-compose.yml")
	switch s.(type) {
	case string:
		writeNewFile("./a.txt", fmt.Sprintf("%s", s))
	}
	//copyFile("E:/go/src/goBoy/day03/docker-compose.yml", "b.txt")

	//var pos int64 = 0
	//for {
	//	fmt.Println("----------------------")
	//	printLog("E:/go/src/goBoy/day03/docker-compose.yml", &pos)
	//	pos += 1024
	//	time.Sleep(time.Second * 3)
	//}

}
