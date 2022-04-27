go generate 在执行的时候会自动注入以下环境变量:

```sh $GOARCH 系统架构: arm, amd64 等
$GOOS 操作系统: linux, windows 等
$GOFILE 当前执行的命令所处的文件名
$GOLINE 当前执行的命令在文件中的行号
$GOPACKAGE 执行的命令所处的文件的包名
$DOLLAR $ 符号

```
