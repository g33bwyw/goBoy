package mylog

import (
	"fmt"
	"runtime"
	"strings"
)

const (
	INFO = iota
	WARN
	ERROR
)

type logLevel uint16

type MyLog struct {
	logLevel logLevel
}

/**
获取日志等级
*/
func parseStringToLogLevel(logLevel string) logLevel {
	logLevel = strings.ToLower(logLevel)
	switch logLevel {
	case "info":
		return INFO
	case "warn":
		return WARN
	case "error":
		return ERROR
	default:
		return INFO
	}
}

/**
获取日志信息
*/
func getInfo(skip int) (fucName,fileName string, line int){
	pc, fileName, line, ok := runtime.Caller(skip)
	if !ok {
		fmt.Println("runtime caller失败")
	}
	fucName = runtime.FuncForPC(pc).Name()
	return

}
