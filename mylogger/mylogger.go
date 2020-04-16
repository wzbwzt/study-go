package mylogger

import (
	"errors"
	"fmt"
	"path"
	"runtime"
	"strings"
)

//该文件放一些共用的方法及信息
// 往终端写日志相关内容

// Loglevel 定义级别类型
type Loglevel uint16

//UNKNOWN 其他等级
const (
	UNKNOWN Loglevel = iota
	DEBUG
	TRACE
	INFO
	WARNING
	ERROR
	FATAL
)

//将错误级别str转换为Loglevel
func parseLoglevel(s string) (Loglevel, error) {
	s = strings.ToLower(s)
	switch s {
	case "debug":
		return DEBUG, nil
	case "trace":
		return TRACE, nil
	case "info":
		return INFO, nil
	case "warning":
		return WARNING, nil
	case "error":
		return ERROR, nil
	case "fatal":
		return FATAL, nil
	default:
		err := errors.New("无效的日志级别")
		return UNKNOWN, err
	}
}

//将错误等级数字转换为string
func loglevelToStr(level Loglevel) (strLog string) {
	switch level {
	case DEBUG:
		return "debug"
	case TRACE:
		return "trace"
	case INFO:
		return "info"
	case WARNING:
		return "warning"
	case ERROR:
		return "error"
	case FATAL:
		return "fatal"
	default:
		return "unknow"
	}
}

//获取报错位置所在行数、funcName、文件名
func getInfo(skip int) (fileName, funcName string, line int) {
	pc, file, line, ok := runtime.Caller(skip)
	if !ok {
		fmt.Println("get info failed")
		return
	}
	fileName = path.Base(file)
	funcName = runtime.FuncForPC(pc).Name()
	return
}
