package main

import (
	"bytes"
	"fmt"
	"io/ioutil"

	"github.com/disintegration/imaging"
	"github.com/micro/go-micro/v2/util/log"
)

//图片的相关处理;缩略图,
func main() {
	//读取本地文件，本地文件尺寸300*400
	imgData, _ := ioutil.ReadFile("./images/juhua.jpg")
	buf := bytes.NewBuffer(imgData)
	image, err := imaging.Decode(buf)
	if err != nil {
		log.Debug(err)
		return
	}

	//生成缩略图，尺寸150*200，并保持到为文件2.jpg
	image = imaging.Resize(image, 150, 200, imaging.Lanczos)
	err = imaging.Save(image, "d:/2.jpg")
	if err != nil {
		fmt.Println(err)
	}

}
