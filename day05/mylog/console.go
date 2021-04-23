package mylog

import (
	"fmt"
	"time"
)

type consoleLog struct {
	lv  logLevel
	dlv logLevel
}

func NewConsoleLog(defaultLevel string) *consoleLog {
	dlv := parseStringToLogLevel(defaultLevel)
	return &consoleLog{
		dlv: dlv,
	}
}

//格式化日志
func (c *consoleLog) log(str string, arg ...interface{}) {
	if c.lv >= c.dlv {
		funcName, fileName, line := getInfo(3)
		msg := fmt.Sprintf(str, arg...)
		fmt.Printf("[%s] [%s] [%s:%s:%d] %s\n", time.Now().Format("2006-01-02 15:04:05"),
			parseLogLevelToString(c.lv), funcName, fileName, line, msg)
	}
}

func (c *consoleLog) Info(str string, arg ...interface{}) {
	c.lv = INFO
	c.log(str, arg...)
}

func (c *consoleLog) Warn(str string, arg ...interface{}) {
	c.lv = WARN
	c.log(str, arg...)
}

func (c *consoleLog) Error(str string, arg ...interface{}) {
	c.lv = ERROR
	c.log(str, arg...)
}
