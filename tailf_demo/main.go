package main

//tailf demo   用于日志收集系统的Log Agent工作流程中的第一步读日志
import (
	"fmt"
	"time"

	"github.com/hpcloud/tail"
)

func main() {
	fileName := "./my.log"
	config := tail.Config{
		ReOpen:    true,                                 //重新打开
		Follow:    true,                                 //是否跟随
		Location:  &tail.SeekInfo{Offset: 0, Whence: 2}, //从文件哪里开始读
		MustExist: false,                                //文件不存在不报错
		Poll:      true,
	}
	tails, err := tail.TailFile(fileName, config)
	if err != nil {
		fmt.Println("tail failed err:", err)
		return
	}
	var (
		line *tail.Line
		ok   bool
	)

	for {
		line, ok = <-tails.Lines
		if !ok {
			fmt.Println("tail file close reopen,filename:%s\n", tails.Filename)
			time.Sleep(time.Second * 1)
			continue
		}
		fmt.Println("line:", line.Text)

	}

}
