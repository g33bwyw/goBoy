/**
解决TCP的黏包问题
**/
package proto

import (
	"bufio"
	"bytes"
	"encoding/binary"
	"fmt"
)

/**
字节信息进行写入
*/
func Encode(message string) ([]byte, error) {
	//1.计算字符串的长度
	strLen := int32(len(message))
	//2.创造字节缓冲方式一
	var b []byte
	buf := bytes.NewBuffer(b)
	//方式二
	//buf := new(bytes.Buffer)
	fmt.Println("----buf-----", buf)
	//3.将长度放入到小端中
	sizeErr := binary.Write(buf, binary.LittleEndian, strLen)
	if sizeErr != nil {
		return []byte{}, sizeErr
	}
	//4.将信息放入到字节缓冲中
	writeErr := binary.Write(buf, binary.LittleEndian, []byte(message))
	if writeErr != nil {
		return []byte{}, sizeErr
	}
	fmt.Println("----buf bytes-----", buf.Bytes())
	return buf.Bytes(), nil
}

/**
字节信息读取
*/
func Decode(reader *bufio.Reader) (string, error) {
	//1.读取长度内容
	lengthByte, _ := reader.Peek(4)
	fmt.Println("------lengthByte-----", lengthByte)
	//if err != nil {
	//	return "", err
	//}
	//2.获取字节长度后缓存数据大小
	var length int32
	lengthBuf := bytes.NewBuffer(lengthByte)
	fmt.Println("------lengthBuf-----", lengthBuf)
	err := binary.Read(lengthBuf, binary.LittleEndian, &length)
	fmt.Println("------length-----", length)
	if err != nil {
		return "", err
	}
	//3.检查大小是否一致
	fmt.Println("------reader Buffered-----", reader.Buffered())
	if int32(reader.Buffered()) < length+4 {
		return "", err
	}
	pack := make([]byte, int(4+length))
	_, err = reader.Read(pack)
	if err != nil {
		return "", err
	}
	return string(pack[4:]), nil
}
