package main

import (
	"goBoy/day05/mylog"
)

func main() {
	//获取日志信息
	//funName,fileName,line := mylog.GetInfo()
	//fmt.Println(funName, fileName, line)
	//for{
	//	consoleLog :=  mylog.NewConsoleLog("warn")
	//	consoleLog.Info("这是info测试")
	//	consoleLog.Warn("这是warn测试")
	//	consoleLog.Error("这是error测试")
	//
	//	time.Sleep(time.Second * 3)
	//}

	//a := fmt.Sprintf("111")
	//fmt.Println(a)

	//fileLog := mylog.NewFileLog("d:", "log.txt", "info")
	//fileLog.Info("这是info测试")

	//file,_ := os.OpenFile("d:/log.txt",os.O_CREATE|os.O_APPEND|os.O_RDWR, 0644)
	//fileInfo, _ := file.Stat()
	//fmt.Println(fileInfo.Size())
	//fileLog :=  mylog.NewFileLog("./","openapi-shop","info", 10*1024)
	//fileLog.Info("这是info测试")
	//fileLog.Warn("这是warn测试")
	//fileLog.Error("这是error测试")
	for {
		fileLog := mylog.NewFileLog("./", "openapi-shop", "info", 10*1024)
		fileLog.Info("这是info测试")
		fileLog.Warn("这是warn测试")
		fileLog.Error("这是error测试")

		//time.Sleep(time.Second * 3)
	}
}
