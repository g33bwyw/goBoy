package mylog

import (
	fmt "fmt"
	"time"
)

type consoleLog struct {
	MyLog
}

func (c *consoleLog) NewConsoleLog(level string) *consoleLog {
	return &consoleLog{
		MyLog{
			logLevel: parseStringToLogLevel(level),
		},
	}
}

//格式化日志
func (c *consoleLog) FormatLog(fomate string, arg ...interface{}) {
	funcName, fileName, line := getInfo(2)
	msg := fmt.Sprintf(fomate, arg)
	fmt.Printf("[%s][%s][%s-%s:%d]:%s", time.Now().Format("2016-01-02 15:04:05"),
		c.MyLog.logLevel, funcName, fileName, line, msg)
}
