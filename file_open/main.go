package main

import (
	"bufio"
	"fmt"
	"io"
	"io/ioutil"
	"os"
)

/**
os.Open()打开的文件只能read;不可以write
*/
//按字节读取文件
func readFromFile1() {
	//打开文件
	fileObj, err := os.Open("./main.go") //读取的文件可以是绝对路径也可以是相对路径
	if err != nil {
		fmt.Printf("file open err:%v", err)
		return
	}
	//记得关闭文件，defer  函数推出时执行；要放在err判断的后面执行，就是得保证fileObj不为空
	defer fileObj.Close()
	//读取文件
	//.Read()相当于直接从磁盘读取文件
	var tmp = make([]byte, 128) //指定读的长度
	for {
		n, err := fileObj.Read(tmp[:])
		if err == io.EOF {
			fmt.Println("file read end")
			return
		}
		if err != nil {
			fmt.Printf("file read from main.go is err:%v", err)
			return
		}

		// fmt.Printf("读了一共%d个字节", n)
		fmt.Println(string(tmp[:n]))
		// if n < 128 {    //方法不准确
		// 	return
		// }
	}
	// fmt.Println(string(tmp[:]))

}

//按行读取文件
func readfileByBufio() {
	fileObj, err := os.Open("./channel.pdf")
	if err != nil {
		fmt.Printf("file open field err:%v", err)
		return
	}
	defer fileObj.Close()

	//创建一个用来从文件中读内容的对象
	//相当于从buf缓存区读取文件；先从磁盘中读文件放在缓存区；
	reader := bufio.NewReader(fileObj)
	for {
		line, err := reader.ReadString('\n')
		if err == io.EOF {
			fmt.Println("file read end")
			return
		}
		if err != nil {
			fmt.Printf("file read err:%v", err)
			return
		}
		fmt.Print(line)
	}
}

//读取整个文件
func readFileByIoutil() {
	ret, err := ioutil.ReadFile("./channel.md")
	if err != nil {
		fmt.Printf("file read filed err is %v", err)
		return
	}
	fmt.Print(string(ret))

}

func main() {
	// readFromFile1()
	// readfileByBufio()
	readFileByIoutil()
	// fileSize, fileName, mode := getFileInfo("./main.go")
	// fmt.Printf("%vB\n", fileSize)
	// fmt.Printf("%v\n", fileName)
	// fmt.Printf("%v", mode)
}

//获取文件信息
func getFileInfo(filePath string) (fileSize int64, fileName string, mode os.FileMode) {
	fileObj, err := os.Open(filePath)
	if err != nil {
		fmt.Println(err)
		return
	}
	fileInfor, err := fileObj.Stat()
	if err != nil {
		fmt.Println(err)
		return
	}
	fileSize = fileInfor.Size() //获取文件大小
	fileName = fileInfor.Name()
	mode = fileInfor.Mode()

	return
}
