package main

import (
	"time"

	"github.com/wzbwzt/studyGo/mylogger"
)

func main() {
	// fileObj, err := os.OpenFile("./sys.log", os.O_CREATE|os.O_APPEND|os.O_WRONLY, 0777)
	// if err != nil {
	// 	fmt.Printf("file open failed err is %v", err)
	// 	return
	// }
	// log.SetOutput(fileObj)
	// for {
	// log.Println("this is a log demo;")
	// time.Sleep(time.Second * 2)
	// }

	//控制台打印日志测试
	//只有当级别大于或等于传入的错误级别时才会打印输出
	// conLog := mylogger.NewConsoleLog("warning")
	// for {

	// 	time.Sleep(time.Second * 2)
	// 	conLog.Debug("this is  debug")
	// 	conLog.Trace("this is  debug")
	// 	conLog.Info("this is info")
	// 	conLog.Warning("this is  warning")
	// 	id := 123
	// 	name := "tiezhu"
	// 	conLog.Error("this is  error|id:%d,name:%s", id, name)
	// 	conLog.Fatal("this is  fatal")
	// }

	//文件输出打印日志测试
	fileLog := mylogger.NewFileLog("info", "./", "syslog", 5*1024)
	for {
		fileLog.Debug("this is  debug")
		fileLog.Trace("this is  debug")
		fileLog.Info("this is info")
		fileLog.Warning("this is  warning")
		id := 123
		name := "tiezhu"
		fileLog.Error("this is  error|id:%d,name:%s", id, name)
		fileLog.Fatal("this is  fatal")
		time.Sleep(time.Second * 2)
	}

}
