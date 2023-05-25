package main

/*
#cgo LDFLAGS: -L/usr/local/lib

#include <stdio.h>
#include <stdlib.h>
#define REPEAT_LIMIT 3              // CGO会保留C代码块中的宏定义
typedef struct{                     // 自定义结构体
    int repeat_time;
    char* str;
}blob;
int SayHello(blob* pblob) {  // 自定义函数
    for ( ;pblob->repeat_time < REPEAT_LIMIT; pblob->repeat_time++){
        puts(pblob->str);
    }
    return 0;
}
*/
import "C" //这里可看作封装的伪包C, 这条语句要紧挨着上面的注释块，不可在它俩之间间隔空行！
import (
	"fmt"
	"unsafe"
)

// CGO 中使用 #cgo 关键字可以设置编译阶段和链接阶段的相关参数，可以使用 ${SRCDIR} 来表示 Go 包当前目录的绝对路径。
//使用 C.结构名 或 C.struct_结构名 可以在 Go 代码段中定义 C 对象，并通过成员名访问结构体成员。

func main() {
	cblob := C.blob{} // 在GO程序中创建的C对象，存储在Go的内存空间 //
	cblob.repeat_time = 0

	cblob.str = C.CString("Hello, World\n") // C.CString 会在C的内存空间申请一个C语言字符串对象，再将Go字符串拷贝到C字符串

	ret := C.SayHello(&cblob) // &cblob 取C语言对象cblob的地址

	fmt.Println("ret", ret)
	fmt.Println("repeat_time", cblob.repeat_time)

	C.free(unsafe.Pointer(cblob.str)) // C.CString 申请的C空间内存不会自动释放，需要显示调用C中的free释放
}
