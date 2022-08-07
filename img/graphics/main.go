package main

import (
	"bytes"
	"encoding/base64"
	"fmt"
	"image"
	"image/color"
	"image/draw"
	"image/jpeg"
	"io/ioutil"
	"os"

	"github.com/disintegration/imaging"
	"github.com/fogleman/gg"
	"github.com/labstack/gommon/log"
)

//图片的相关处理;缩略图,图片的压缩，画图
func main() {
	//读取本地文件，本地文件尺寸300*400
	imgData, _ := ioutil.ReadFile("../images/dog.png")
	buf := bytes.NewBuffer(imgData)
	image, err := imaging.Decode(buf)
	if err != nil {
		log.Debug(err)
		return
	}

	// compressImageResource(imgData)

	//生成缩略图，尺寸150*200，并保持到为文件2.jpg
	image = imaging.Resize(image, 24, 24, imaging.Lanczos)
	var bufNew bytes.Buffer
	err = jpeg.Encode(&bufNew, image, &jpeg.Options{Quality: 90})
	// err = imaging.Save(image, "../images/dog3.0.jpg")
	if err != nil {
		fmt.Println(err)
	}
	base64_bufNew := base64.StdEncoding.EncodeToString(bufNew.Bytes())
	fmt.Println(base64_bufNew)

	//////////////////////////////////////////////////////////////////////////
	// const S = 1024
	// dc := gg.NewContext(S, S)
	// dc.SetRGBA(0, 0, 0, 0.1)
	// for i := 0; i < 360; i += 15 {
	// 	dc.Push()
	// 	dc.RotateAbout(gg.Radians(float64(i)), S/2, S/2)
	// 	dc.DrawEllipse(S/2, S/2, S*7/16, S/8)
	// 	dc.Fill()
	// 	dc.Pop()
	// }
	// dc.SavePNG("../images/out.png")
	///////////////////////////////////////////////////////////////////////////

	// drawBackground()

}

//压缩图片
func compressImageResource(data []byte) []byte {
	img, _, err := image.Decode(bytes.NewReader(data))
	if err != nil {
		return data
	}
	buf := bytes.Buffer{}
	err = jpeg.Encode(&buf, img, &jpeg.Options{Quality: 60})
	if err != nil {
		return data
	}
	if buf.Len() > len(data) {
		return data
	}
	return buf.Bytes()
}

func drawBackground() {
	// 加载图片
	img, err := gg.LoadJPG("../images/dog.png")
	if err != nil {
		panic(err)
	}

	//获取图片尺寸
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
	// dc.DrawPoint(20, 20, 8)

	dc.Stroke() //dc.Fill()   必须有一个绘制呈现

	err = dc.SavePNG("../images/dog30.0.jpg")

	// 保存新图片，一般quality为图片质量最高为设置为100，值越高，质量越好，但是占空间大
	err = gg.SaveJPG("../images/dog1.0.png", img, 1)
	// err = gg.SaveJPG("../images/dog90.0.png", img, 90)
	// err = gg.SaveJPG("../images/dog70.0.jpg", img, 70)
	if err != nil {
		panic(err)
	}

	//生成缩略图
	// src, err := imaging.Open("aifile/maps/1K" + ks + ".png")
	// if err != nil {
	// 	panic(err)
	// } else {
	// 	src = imaging.Resize(src, 1250, 0, imaging.Lanczos)
	// 	smallimgPath := "aifile/maps/1Kaaa" + ks + ".png"
	// 	imaging.Save(src, smallimgPath)
	// }

}

func main2() {
	f1, err := os.Open("out2.0.jpg")
	if err != nil {
		panic(err)
	}
	defer f1.Close()
	f2, err := os.Open("out.jpg")
	if err != nil {
		panic(err)
	}
	defer f2.Close()

	f3, err := os.Create("3.jpg")
	if err != nil {
		panic(err)
	}
	defer f3.Close()

	m1, err := jpeg.Decode(f1)
	if err != nil {
		panic(err)
	}
	bounds := m1.Bounds()
	m2, err := jpeg.Decode(f2)
	if err != nil {
		panic(err)
	}
	m := image.NewRGBA(bounds)
	white := color.RGBA{255, 255, 255, 255}
	draw.Draw(m, bounds, &image.Uniform{white}, image.ZP, draw.Src)
	draw.Draw(m, bounds, m1, image.ZP, draw.Src)
	draw.Draw(m, image.Rect(100, 200, 300, 600), m2, image.Pt(250, 60), draw.Src)
	err = jpeg.Encode(f3, m, &jpeg.Options{Quality: 90})
	if err != nil {
		panic(err)
	}
	fmt.Printf("ok\n")
}
