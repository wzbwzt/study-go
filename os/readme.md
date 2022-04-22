`os` 包提供了平台无关的操作系统功能接口。尽管错误处理是 go 风格的，但设计是 Unix 风格的；所以，
失败的调用会返回 error 而非错误码。通常 error 里会包含更多信息。例如，如果使用一个文件名的调用
（如 Open、Stat）失败了，打印错误时会包含该文件名，错误类型将为 `*PathError`，其内部可以解包获得更多信息。

os 包规定为所有操作系统实现的接口都是一致的。有一些某个系统特定的功能，需要使用 syscall 获取。
实际上，os 依赖于 syscall。在实际编程中，我们应该总是优先使用 os 中提供的功能，而不是 syscall。

```go
const (
    O_RDONLY int = syscall.O_RDONLY // 只读模式打开文件
    O_WRONLY int = syscall.O_WRONLY // 只写模式打开文件
    O_RDWR   int = syscall.O_RDWR   // 读写模式打开文件
    O_APPEND int = syscall.O_APPEND // 写操作时将数据附加到文件尾部
    O_CREATE int = syscall.O_CREAT  // 如果不存在将创建一个新文件
    O_EXCL   int = syscall.O_EXCL   // 和 O_CREATE 配合使用，文件必须不存在
    O_SYNC   int = syscall.O_SYNC   // 打开文件用于同步 I/O
    O_TRUNC  int = syscall.O_TRUNC  // 如果可能，打开时清空文件
)

```

其中，`O_RDONLY、O_WRONLY、O_RDWR` 应该只指定一个，剩下的通过 | 操作符来指定。该函数内部会给 flags 加上
syscall.O_CLOEXEC，在 fork 子进程时会关闭通过 OpenFile 打开的文件，即子进程不会重用该文件描述符。

> 注意：由于历史原因，O_RDONLY | O_WRONLY 并非等于 O_RDWR，它们的值一般是 0、1、2
