# go1.18 新增三大功能之-工作区模式

## Go v1.18 版本新增 Workspaces 模式主要解决的问题

当在项目中需要导入一些私有 Git 仓库或未发布到 Git 仓库的依赖项时，或需要修改三方依赖模块的代码时，一般解决方法是将代码下载到本地，在 go.mod 文件中使用 Go Modules 模式中的 replace 指令替换为本地目录路径。
因为 go.mod 文件也是和项目一起提交到 Git 仓库，每个开发人员的本地目录路径可能不一样，所以，就需要把在远程仓库拉取到的 go.mod 文件中 replace 的目录路径手动替换为自己的本地目录路径之后，才可以正常使用。

而 workspaces 就可以在本地创建一个工作区，go.mod 文件不变的情况下，下载依赖到项目中，并且更改，当执行项目时先从工作区查找依赖包

## 例子

```bash
# 进入 Home 目录
cd ~
# 创建工作区目录
mkdir workspace
# 进入工作区目录
cd workspace
# 创建 Go 项目所在的目录
mkdir hello
# 进入 Go 项目所在的目录
cd hello
# 初始化
go mod init github.com/weirubo/hello
# 添加依赖项
go get golang.org/x/example
# 创建 main.go 文件
touch main.go
```

执行 `go run main.go` output:olleh
注意：
workspace 目录下执行`go run hello/main.go`报错:go: go.mod file not found in current directory or any parent directory; see 'go help modules' (当前夫级目录中有 go.mod 文件所以不会报错!!)

这时可以创建一个工作区:

```bash
# 进入 Home 目录下的 workspace 目录
cd ~/workspace
# 初始化工作区
go work init ./hello

```

执行完以上命令，我们可以发现在 workspace 目录下创建了一个 go.work 文件，查看该文件，我们可以发现包含两个指令，分别是 go 和 use。
其中 go 指令，是指定使用 Go 的哪个版本编译项目，类似 go.mod 文件中的 go 指令。
其中 use 指令，是指在构建项目时，hello 目录中的模块是主模块。

在创建工作区后，我们进入工作区目录，运行 main.go 文件。
`go run hello/main.go` output:olleh

但是进入工作区外同样会报错:
`go run workspace/hello/main.go` 报错:workspace/hello/main.go:5:2: no required module provides package golang.org/x/example/stringutil: go.mod file not found in current directory or any parent directory; see 'go help modules'

### 如果想要修改一个依赖文件

进入 hello 目录 clone 依赖文件 `git clone https://go.googlesource.com/example`
进入工作区目录执行`go work use ./hello/example`

(go work use 命令将 example 模块添加到工作区，我们就可以使用我们下载到本地的依赖模块 example 的代码，而不再使用 GOMODCACHE 中的 example 模块的代码。
我们项目组成员只需维护自己本地的工作区，不必再手动修改 go.mod 文件中 replace 指令中的本地目录路径。)
