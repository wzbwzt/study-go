package mylogger

import (
	"fmt"
	"os"
	"path"
	"time"
)

//往文件中记录日志

//FileLogger 对象结构体
type FileLogger struct {
	Level       Loglevel
	filePath    string
	fileName    string
	maxFileSize int64
	fileObj     *os.File //记录所有错误的文件句柄
	errFileObj  *os.File //记录error文件句柄
	logChan     chan *chanMsg
}

// 建立一个通道结构体来存储信息
type chanMsg struct {
	level      Loglevel
	msg        string
	funcName   string
	fileObj    *os.File
	errFileObj *os.File
	timeStamp  string
	line       int
	allMsg     string
}

//ChanSize  通道的缓存大小
var ChanSize = 5000

//NewFileLog 构造函数
func NewFileLog(levelStr, fp, fn string, maxSize int64) *FileLogger {
	levelInt, err := parseLoglevel(levelStr)
	if err != nil {
		panic(err)
	}
	f := &FileLogger{
		Level:       levelInt,
		filePath:    fp,
		fileName:    fn,
		maxFileSize: maxSize,
		logChan:     make(chan *chanMsg, ChanSize),
	}
	err = f.initFile() //打开文件
	if err != nil {
		panic(err)
	}
	//开启一个后台的goroutine来写日志
	go f.getLogFromChan()
	return f
}

//文件创建与打开  获得文件句柄
func (f *FileLogger) initFile() error {
	fileFullPath := path.Join(f.filePath, f.fileName)
	//开启log文件
	fileObj, err := os.OpenFile(fileFullPath+".log", os.O_CREATE|os.O_APPEND|os.O_WRONLY, 0777)
	if err != nil {
		fmt.Printf("open log failed the err is %v", err)
		return err
	}
	//开启errLog文件
	errFileObj, err := os.OpenFile(fileFullPath+".err.log", os.O_CREATE|os.O_APPEND|os.O_WRONLY, 0777)
	if err != nil {
		fmt.Printf("open log failed the err is %v", err)
		return err
	}
	f.fileObj = fileObj
	f.errFileObj = errFileObj
	return nil
}

//文件句柄关闭
func (f *FileLogger) clossFile() {
	f.fileObj.Close()
	f.errFileObj.Close()

}

//判断文件大小 根据指定的文件大小来切割文件
//1.获取日志文件的大小
//2.判断日志文件是否大于指定的文件大小;大于指定文件size的话切割文件
//3.关闭之前打开的日志文件句柄
//4.按照指定格式创建打开一个新的日志文件
//5.返回新文件的句柄
func (f *FileLogger) cutBySize() error {
	//1
	fileInfo, err := f.fileObj.Stat()
	if err != nil {
		fmt.Printf("stat file failed err is %v", err)
		return err
	}
	errFileInfo, err := f.errFileObj.Stat()
	if err != nil {
		fmt.Printf("stat file failed err is %v", err)
		return err
	}
	fileSize := fileInfo.Size()       //记录所有日志的文件size
	errFileSize := errFileInfo.Size() //记录errr日志文件的大小
	//2
	if fileSize >= f.maxFileSize {
		//3
		f.fileObj.Close()
		//4
		fileFullpath := path.Join(f.filePath, f.fileName)
		now := time.Now()
		timeStr := now.Format("20060102150405")
		newFileObj, err := os.OpenFile(fileFullpath+timeStr+".log", os.O_CREATE|os.O_APPEND|os.O_WRONLY, 0777)
		if err != nil {
			fmt.Printf("newLogFile  open failed err is %v", err)
			return err
		}
		f.fileObj = newFileObj
		return nil
	}
	if errFileSize >= f.maxFileSize {
		f.errFileObj.Close()
		fileFullpath := path.Join(f.filePath, f.fileName)
		now := time.Now()
		timeStr := now.Format("20060102150405")
		newErrFileObj, err := os.OpenFile(fileFullpath+timeStr+".err.log", os.O_CREATE|os.O_APPEND|os.O_WRONLY, 0777)
		if err != nil {
			fmt.Printf("newLogFile  open failed err is %v", err)
			return err
		}
		f.errFileObj = newErrFileObj
		return nil
	}
	return nil
}

//异步写入日志 往通道内写入日志
func (f *FileLogger) wrLogToChan(msg string, lever Loglevel, a ...interface{}) {
	if f.enable(lever) {
		now := time.Now()
		msg = fmt.Sprintf(msg, a...)
		strlog := loglevelToStr(lever)
		fileName, funcName, line := getInfo(4)
		formatMsg := fmt.Sprintf("[%v][%v:%v:%v]{%v}%s\n", now.Format("2006/01/02 15:04:05"), fileName, funcName, line, strlog, msg)
		logChanStru := &chanMsg{
			level:      lever,
			msg:        msg,
			funcName:   funcName,
			fileObj:    f.fileObj,
			errFileObj: f.errFileObj,
			timeStamp:  now.Format("2006/01/02 15:04:05"),
			line:       line,
			allMsg:     formatMsg,
		}
		select {
		case f.logChan <- logChanStru:
		default:
			//通道满的话就丢掉，确保不会堵塞
		}
	}
}

//从通道中取日志信息并写入文件
func (f *FileLogger) getLogFromChan() {
	for {
		select {
		case logMsg := <-f.logChan:
			err := f.cutBySize()
			if err != nil {
				panic(err)
			}
			fmt.Fprintf(f.fileObj, logMsg.allMsg)
			if logMsg.level >= ERROR {
				err := f.cutBySize()
				if err != nil {
					panic(err)
				}
				//要记录的日志大于error时还需要单独记一封
				fmt.Fprintf(f.errFileObj, logMsg.allMsg)
			}
		default:
			time.Sleep(time.Millisecond * 500) //取不到日志的话就让他休息，让出CPU
		}
	}

}

//日志写入文件具体操作
func (f *FileLogger) wrLogToFile(msg string, lever Loglevel, a ...interface{}) {
	f.wrLogToChan(msg, lever, a...)

}

//判断哪些错误级别写入 返回bool
func (f *FileLogger) enable(level Loglevel) bool {
	return f.Level <= level
}

//Debug debug级别
func (f *FileLogger) Debug(msg string, a ...interface{}) {
	f.wrLogToFile(msg, DEBUG, a...)
}

//Trace Trace级别
func (f *FileLogger) Trace(msg string, a ...interface{}) {
	f.wrLogToFile(msg, TRACE, a...)
}

//Info Info级别
func (f *FileLogger) Info(msg string, a ...interface{}) {
	f.wrLogToFile(msg, INFO, a...)
}

//Warning Warning级别
func (f *FileLogger) Warning(msg string, a ...interface{}) {
	f.wrLogToFile(msg, WARNING, a...)
}

//Error error级别
func (f *FileLogger) Error(msg string, a ...interface{}) {
	f.wrLogToFile(msg, ERROR, a...)
}

//Fatal fatal级别
func (f *FileLogger) Fatal(msg string, a ...interface{}) {
	f.wrLogToFile(msg, FATAL, a...)
}
