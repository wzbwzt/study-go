package main

import (
	"fmt"
	qrcode "github.com/skip2/go-qrcode"
	"image/color"
)

func  main(){
	/*
	WriteFile函数的原型定义如上，它有几个参数，大概意思如下：
	content表示要生成二维码的内容，可以是任意字符串。
	level表示二维码的容错级别，取值有Low、Medium、High、Highest。
	size表示生成图片的width和height，像素单位。
	filename表示生成的文件名路径。

	## RecoveryLevel类型其实是个int,它的定义和常量如下(RecoveryLevel越高，二维码的容错能力越好。)。
	type RecoveryLevel int
	const (
	    // Level L: 7% error recovery.
	    Low RecoveryLevel = iota
	    // Level M: 15% error recovery. Good default choice.
	    Medium
	    // Level Q: 25% error recovery.
	    High
	    // Level H: 30% error recovery.
	    Highest
	)
	*/
	qrcode.WriteFile("https://blog.csdn.net/wzb_wzt",qrcode.Medium,256,"./QRcode/demo.png")

	/*
	Encode 用法和WriteFile函数差不多，只不过返回的是一个[]byte字节数组，这样我们就可以对这个字节数组进行处理了。
	*/
	encode, _ := qrcode.Encode("https://blog.csdn.net/wzb_wzt", qrcode.Medium, 256)
	fmt.Println(encode)

	//自定义设置二维码样式
	//生成一个qrcode struct指针
	code, _ := qrcode.New("https://blog.csdn.net/wzb_wzt", qrcode.High)
	//背景色是绿色
	code.BackgroundColor=color.RGBA{50,205,50,255}
	//前置色是白色
	code.ForegroundColor=color.White
	code.DisableBorder=false
	//生成
	code.WriteFile(256,"./QRcode/demo2.png")



}
