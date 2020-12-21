package main

import (
	"bytes"
	"fmt"
	"io/ioutil"

	"github.com/disintegration/imaging"
	"github.com/micro/go-micro/v2/util/log"

	"github.com/fogleman/gg"
)

//图片的相关处理;缩略图,
func main() {
	//读取本地文件，本地文件尺寸300*400
	imgData, _ := ioutil.ReadFile("../images/juhua.jpg")
	buf := bytes.NewBuffer(imgData)
	image, err := imaging.Decode(buf)
	if err != nil {
		log.Debug(err)
		return
	}

	//生成缩略图，尺寸150*200，并保持到为文件2.jpg
	image = imaging.Resize(image, 150, 0, imaging.Lanczos)
	err = imaging.Save(image, "../images/juhua2.0.jpg")
	if err != nil {
		fmt.Println(err)
	}
	//////////////////////////////////////////////////////////////////////////
	const S = 1024
	dc := gg.NewContext(S, S)
	dc.SetRGBA(0, 0, 0, 0.1)
	for i := 0; i < 360; i += 15 {
		dc.Push()
		dc.RotateAbout(gg.Radians(float64(i)), S/2, S/2)
		dc.DrawEllipse(S/2, S/2, S*7/16, S/8)
		dc.Fill()
		dc.Pop()
	}
	dc.SavePNG("../images/out.png")
	///////////////////////////////////////////////////////////////////////////
	drawBackground()
}

func drawBackground() {
	// 加载图片
	// img, err := gg.LoadPNG("map.png")
	img, err := gg.LoadJPG("../images/juhua.jpg")
	if err != nil {
		panic(err)
	}

	// 获取图片尺寸
	size := img.Bounds().Size()

	width := size.X
	height := size.Y

	// 以加载图片的宽高作为新图片的大小
	dc := gg.NewContext(width, height)

	// 画图
	dc.DrawImage(img, 0, 0)
	//获取的坐标
	//点的颜色设置
	dc.SetRGB(0, 1, 0)
	//点的坐标x，y，半径
	dc.DrawPoint(20, 20, 8)

	// dc.Stroke()   		dc.Fill()   必须有一个绘制呈现
	dc.Fill()

	err = dc.SavePNG("../images/juhua3.0.jpg")
	// 保存新图片，一般quality设置为75即可，最高可设置为100，值越高，质量越好，但是占空间大
	// err = dc.SaveJPG("aifile/maps/1K1"+ks+".jpg", 1)
	if err != nil {
		panic(err)
	}

	//开始压缩图片
	// src, err := imaging.Open("aifile/maps/1K" + ks + ".png")
	// if err != nil {
	// 	panic(err)
	// } else {
	// 	src = imaging.Resize(src, 1250, 0, imaging.Lanczos)
	// 	smallimgPath := "aifile/maps/1Kaaa" + ks + ".png"
	// 	imaging.Save(src, smallimgPath)
	// }

}
