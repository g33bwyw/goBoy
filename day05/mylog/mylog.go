package mylog

import (
	"fmt"
	"path"
	"runtime"
	"strings"
)

const (
	INFO = iota
	WARN
	ERROR
)

type logLevel uint16

type MyLog interface {
	Info(str string, arg ...interface{})
	Warn(str string, arg ...interface{})
	Error(str string, arg ...interface{})
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
将日志等级解析成字符串
*/
func parseLogLevelToString(lv logLevel) string {
	switch lv {
	case INFO:
		return "info"
	case WARN:
		return "warn"
	case ERROR:
		return "error"
	default:
		return "info"
	}
}

/**
获取日志信息
*/
func getInfo(skip int) (fucName, fileName string, line int) {
	pc, fileName, line, ok := runtime.Caller(skip)
	if !ok {
		fmt.Println("runtime caller失败")
	}
	fucName = runtime.FuncForPC(pc).Name()
	funcSlice := strings.Split(fucName, ".")
	fucName = funcSlice[1]
	fileName = path.Base(fileName)
	return

}
