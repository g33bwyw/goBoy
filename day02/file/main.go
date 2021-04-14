package main

import (
	"bufio"
	"fmt"
	"io"
	"math"
	"os"
	"strconv"
	"strings"
)

func main() {
	var filename,statFile string
	fmt.Println("请输入原始Excel地址：")
	fmt.Scan(&filename)
	fmt.Println("请输入输出Excvel地址")
	fmt.Scan(&statFile)

	//filename := "D:/goboy/src/goBoy/day02/file/daily_4_video.csv"
	newFileContent := make(map[int64]int64)

	file,err:= os.Open(filename)
	if err != nil {
		fmt.Printf("文件打开失败")
		return
	}
	//循环读取文件
	r := bufio.NewReader(file)
	i := 1
	for  {
		buf,err:=r.ReadBytes('\n')
		if err != nil {
			if err == io.EOF {
				break
			}
		}
		if i == 1 {
			i++
			continue
		}
		content := string(buf)
		content = strings.Trim(content,"\r\n")
		//fmt.Printf("%#v", content)
		if strings.Contains(content, ",") {
			sslice := strings.Split(content,",")

			videoFloatLongTime,_ := strconv.ParseFloat(sslice[0],64)
			videoLongTime := int64(math.Ceil(videoFloatLongTime))
			videoNum,_ := strconv.ParseInt(sslice[1], 10, 32)

			v,ok := newFileContent[videoLongTime]
			if ok  {
				newFileContent[videoLongTime] = videoNum + v
			} else {
				newFileContent[videoLongTime] = videoNum
			}
		} else {
			i++
			continue
		}
		i++
	}

	//statFile := "d:/output.csv"
	file2,err := os.Create(statFile)
	if err != nil {
		fmt.Printf("文件创建失败")
		return
	}
	file2.WriteString("视频时长,视频数量\r\n")
	for i,v := range newFileContent {
		file2.WriteString(fmt.Sprintf("%d,%d\r\n", i,v))
	}

	defer file.Close()
	defer file2.Close()

	fmt.Println("文件创建完成")

}
