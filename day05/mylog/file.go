package mylog

import (
	"fmt"
	"os"
	"path"
	"time"
)

type fileLog struct {
	filepath string
	filename string
	maxSize  int64
	lv       logLevel
	dlv      logLevel
}

func NewFileLog(filepath, filename, defaultLevel string, maxSize int64) *fileLog {
	dlv := parseStringToLogLevel(defaultLevel)
	return &fileLog{
		filepath: filepath,
		filename: filename,
		dlv:      dlv,
		maxSize:  maxSize,
	}
}

/**
获取文件大小
*/
func (f *fileLog) getFileSize(file *os.File) int64 {
	fileInfo, _ := file.Stat()
	return fileInfo.Size()
}

//格式化日志
func (f *fileLog) log(str string, arg ...interface{}) error {
	if f.lv >= f.dlv {
		//error级别以下的记录到info文件中去
		fileName := ""
		if f.lv >= ERROR {
			fileName = f.filename + ".error"
		} else {
			fileName = f.filename + ".info"
		}
		filePath := path.Join(f.filepath, fileName)
		fileObj, err := os.OpenFile(filePath, os.O_CREATE|os.O_APPEND|os.O_RDWR, 0644)
		if err != nil {
			return err
		}
		fileSize := f.getFileSize(fileObj)
		//当超过了指定文件的大小
		if fileSize > f.maxSize {
			fileObj.Close()
			newFilePath := filePath + time.Now().Format("20060102150405")
			err := os.Rename(filePath, newFilePath)
			if err != nil {
				return err
			}
			fileObj, err = os.OpenFile(filePath, os.O_CREATE|os.O_APPEND|os.O_RDWR, 0644)
		}
		funcName, fileName, line := getInfo(3)
		msg := fmt.Sprintf(str, arg...)
		fmt.Fprintf(fileObj, "[%s] [%s] [%s:%s:%d] %s\n", time.Now().Format("2006-01-02 15:04:05"),
			parseLogLevelToString(f.lv), funcName, fileName, line, msg)
		fileObj.Close()

	}
	return nil
}

func (f *fileLog) Info(str string, arg ...interface{}) {
	f.lv = INFO
	f.log(str, arg...)
}

func (f *fileLog) Warn(str string, arg ...interface{}) {
	f.lv = WARN
	f.log(str, arg...)
}

func (f *fileLog) Error(str string, arg ...interface{}) {
	f.lv = ERROR
	f.log(str, arg...)
}
