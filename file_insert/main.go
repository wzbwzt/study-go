package main

import (
	"fmt"
	"io"
	"os"
)

func fileInsert() {
	//打开原文件；
	fileObjOld, err := os.OpenFile("./insertOld.txt", os.O_RDWR, 0777)
	if err != nil {
		fmt.Printf("file open failed the err is %v\n", err)
		return
	}
	//创建一个新文件用来储存做插入操作
	fileObjTmp, errTmp := os.OpenFile("./insertTmp.txt", os.O_CREATE|os.O_TRUNC|os.O_RDWR, 0777)
	if errTmp != nil {
		fmt.Printf("tmp failed err is %v\n", errTmp)
		return
	}
	//写入前面部分
	var s1 = make([]byte, 1)
	n, errFirst := fileObjOld.Read(s1[:])
	if errFirst != nil {
		return
	}
	fileObjTmp.Write(s1[:n])
	//写入要插入的文件
	fileObjTmp.Write([]byte("22"))
	//此时光标在文件最后位置不需要再移动光标
	var last = make([]byte, 2)
	for {
		lastN, errlast := fileObjOld.Read(last[:])
		if errlast == io.EOF {
			fileObjTmp.Write(last[:lastN])
			break
		}
		if errlast != nil {
			fmt.Printf("last read failed the err is %v", errlast)
			return
		}
		fileObjTmp.Write(last[:lastN])
	}

	fileObjOld.Close()
	fileObjTmp.Close()
	renameErr := os.Rename("./insertTmp.txt", "./insertOld.txt")
	if renameErr != nil {
		fmt.Printf("rename failed the err is %v", renameErr)
		return
	}

	//移动光标
	// _, err2 := fileObjOld.Seek(1, 0)
	// if err2 != nil {
	// 	fmt.Printf("file seek failed err is %v\n", err2)
	// 	return
	// }
	// var s = make([]byte, 2)
	// for {
	// 	_, errRead := fileObjOld.Read(s[:])
	// 	if errRead == io.EOF {
	// 		return
	// 	}
	// 	if errRead != nil {
	// 		fmt.Printf("file befer read failed err is %v", errRead)
	// 		return
	// 	}
	// }

	// var s = make([]byte, 1)
	// n, err3 := fileObj.Read(s[:])
	// if err3 != nil {
	// 	fmt.Printf("file read failed err is %v\n", err3)
	// 	return
	// }
	// fmt.Println(string(s[:n]))
}

func main() {
	fileInsert()
}
