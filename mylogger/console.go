package mylogger

import (
	"fmt"
	"time"
)

//往控制台打印错误日志

// ConsoleLogger 日志结构体
type ConsoleLogger struct {
	Level Loglevel
}

//NewConsoleLog 构造函数
func NewConsoleLog(levelStr string) ConsoleLogger {
	logLevel, err := parseLoglevel(levelStr)
	if err != nil {
		panic(err)
	}
	return ConsoleLogger{
		Level: logLevel,
	}
}

//判断哪些错误级别打印 返回bool
func (c ConsoleLogger) enable(level Loglevel) bool {
	return c.Level <= level
}

//封装打印内容函数
func (c ConsoleLogger) wrLogToCon(msg string, logLev Loglevel, a ...interface{}) {
	if c.enable(logLev) {
		now := time.Now()
		msg = fmt.Sprintf(msg, a...)
		strlog := loglevelToStr(logLev)
		fileName, funcName, line := getInfo(3)
		fmt.Printf("[%s][%s][%s|%s|%v] %s\n", now.Format("2006/01/02 15:04:05"), strlog, fileName, funcName, line, msg)
	}

}

//Debug debug级别
func (c ConsoleLogger) Debug(msg string, a ...interface{}) {
	c.wrLogToCon(msg, DEBUG, a...)

}

//Trace Trace级别
func (c ConsoleLogger) Trace(msg string, a ...interface{}) {
	c.wrLogToCon(msg, TRACE, a...)

}

//Info Info级别
func (c ConsoleLogger) Info(msg string, a ...interface{}) {
	c.wrLogToCon(msg, INFO, a...)

}

//Warning Warning级别
func (c ConsoleLogger) Warning(msg string, a ...interface{}) {
	c.wrLogToCon(msg, WARNING, a...)

}

//Error error级别
func (c ConsoleLogger) Error(msg string, a ...interface{}) {
	c.wrLogToCon(msg, ERROR, a...)

}

//Fatal fatal级别
func (c ConsoleLogger) Fatal(msg string, a ...interface{}) {
	c.wrLogToCon(msg, FATAL, a...)

}
