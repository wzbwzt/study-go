package main

import (
	rotatelogs "github.com/lestrrat-go/file-rotatelogs"
	"github.com/rifflock/lfshook"
	log "github.com/sirupsen/logrus"
	"os"
	"time"
)

/*
1.logrus与golang标准库日志模块完全兼容，因此可以使用log“github.com/sirupsen/logrus”替换所有日志导入。

*/


//基本使用
func main1(){
	//以json的形式输出，默认是以TextFormatter文档的形式输出
	log.SetFormatter(&log.JSONFormatter{
		TimestampFormat: "2006.01.02/03:04:05",
	})
	file, _ := os.OpenFile("./logrus/main1.log", os.O_CREATE|os.O_APPEND|os.O_WRONLY, 0777)
	//log输出到io.Wirter接口，默认是os.Stdout
	log.SetOutput(file)
	//输出的错误等级，>=设置的等级；logrus一共有6个等级
	log.SetLevel(log.DebugLevel)
	//WithFields API可以规范使用者按照其提倡的方式记录日志
	log.WithFields(log.Fields{
		"msg":"found not",
	}).Debug("hhhhhh")
	log.WithFields(log.Fields{
		"msg":"found not",
	}).Info("hhhhhh")
	log.WithFields(log.Fields{
		"msg":"found not",
	}).Warn("hhhhhh")
	log.WithFields(log.Fields{
		"msg":"found not",
	}).Error("hhhhhh")
	log.WithFields(log.Fields{
		"msg":"found not",
	}).Fatal("hhhhhh")
	log.WithFields(log.Fields{
		"msg":"found not",
	}).Panic("hhhhhh")

	//输出单条信息
	log.Warn("sfasdf")
}

//进阶使用new一个logger实例
func main2(){
	logger := log.New()
	file, _ := os.OpenFile("./logrus/main2.log", os.O_CREATE|os.O_APPEND|os.O_WRONLY, 0777)
	//设置输出的io.writer
	logger.Out=file
	//设置输出的错误等级
	logger.Level=log.DebugLevel
	//设置输出的格式
	logger.Formatter=&log.JSONFormatter{
		//输出的time字段是自带的，可以对时间输出的格式做限制
		TimestampFormat: "2006.01.02/03:04:05",
	}

	logger.WithFields(log.Fields{
		"msg":"found not",
	}).Warn("warning err")
}

//进阶使用hook
//logrus的hook接口原理是每此写入日志时拦截，修改logrus.Entry。
//使用rotatelogs进行文件的分割存储
func main(){
	filename:="./logrus/main3"
	logger := log.New()
	logger.SetLevel(log.DebugLevel)

	//切割文件
	rotateLog, err:= rotatelogs.New(
		//设置分割的文件明
		filename+".%Y%m%d%H.log",
		// WithLinkName为最新的日志建立软连接，以方便随着找到当前日志文件
		rotatelogs.WithLinkName(filename),
		// WithMaxAge设置最大保存时间(7天)
		//WithMaxAge和WithRotationCount二者只能设置一个，WithRotationCount设置文件清理前最多保存的个数。
		rotatelogs.WithMaxAge(7*24*time.Hour),
		//设置日志分割的时间，这里设置为一小时分割一次
		rotatelogs.WithRotationTime(time.Hour),
	)
	if err != nil {
		logger.Errorf("config rotetelog failed err: %v", err)
	}
	//对不同的等级log设置io.writer
	writerMap := lfshook.WriterMap{
		log.DebugLevel: rotateLog,
		log.InfoLevel:  rotateLog,
		log.WarnLevel:  rotateLog,
		log.ErrorLevel: rotateLog,
		log.FatalLevel: rotateLog,
		log.PanicLevel: rotateLog,
	}
	lfsHook := lfshook.NewHook(writerMap, log.Formatter(&log.JSONFormatter{
		TimestampFormat: "2006.01.02/03:04:05",
	}))
	logger.AddHook(lfsHook)

	i:=0
	for  {

		logger.WithFields(log.Fields{
			"msg":"found not",
		}).Warn("warning err",i)

		<-time.After(time.Second*1)
		i++
	}
}



