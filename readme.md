# VSCode安装GO语言依赖工具

|名称|描述|链接|
|---|---|---|
|gocode|代码自动补全|https://github.com/mdempsky/gocode|
|go-outline|在当前文件中查找|https://github.com/ramya-rao-a/go-outline|
|go-symbols|在项目路径下查找|https://github.com/acroca/go-symbols|
|gopkgs|自动补全未导入包|https://github.com/uudashr/gopkgs|
|guru|查询所有引用|https://golang.org/x/tools/cmd/guru|
|gorename|重命名符号|https://golang.org/x/tools/cmd/gorename|
|goreturns|格式化代码|https://github.com/sqs/goreturns|
|godef|跳转到声明|https://github.com/rogpeppe/godef|
|godoc|鼠标悬浮时文档提示|https://golang.org/x/tools/cmd/godoc|
|golint|就是lint|https://golang.org/x/lint/golint|
|dlv|调试功能|https://github.com/derekparker/delve/tree/master/cmd/dlv|
|gomodifytags|修改结构体标签|https://github.com/fatih/gomodifytags|
|goplay|运行当前go文件|https://github.com/haya14busa/goplay/|
|impl|新建接口|https://github.com/josharian/impl|
|gotype-live|类型诊断|https://github.com/tylerb/gotype-live|
|gotests|单元测试|https://github.com/cweill/gotests/|
|go-langserver|语言服务|https://github.com/sourcegraph/go-langserver|
|filstruct|结构体成员默认值|https://github.com/davidrjenni/reftools/tree/master/cmd/fillstruct|

# 安装步骤

1. 一般VSCODE会提示自动安装，如果需要可手工安装
```sh
go install github.com/ramya-rao-a/go-outline
go install github.com/acroca/go-symbols
go install github.com/mdempsky/gocode
go install github.com/rogpeppe/godef
go install github.com/zmb3/gogetdoc
go install github.com/fatih/gomodifytags
go install github.com/sqs/goreturns
go install github.com/cweill/gotests/...
go install github.com/josharian/impl
go install github.com/haya14busa/goplay/cmd/goplay
go install github.com/uudashr/gopkgs/cmd/gopkgs
go install github.com/davidrjenni/reftools/cmd/fillstruct
go install github.com/alecthomas/gometalinter
$GOPATH/bin/gometalinter --install
go install golang.org/x/tools/cmd/godoc
go install golang.org/x/lint/golint
go install golang.org/x/tools/cmd/gorename
go install golang.org/x/tools/cmd/goimports
go install golang.org/x/tools/cmd/guru
```
2. 安装问题
  > cannot find package "golang.org/x/net/context/ctxhttp"
```sh
cd $GOPATH/src/golang.org/x
git clone https://github.com/golang/net
```
  > package golang.org/x/dom: unrecognized import path "golang.org/x/dom": https fetch: Get "https://golang.org/x/dom?go-get=1": dial tcp 216.239.37.1:443: i/o timeout
```sh
go env -w GO111MODULE=on
go env -w GOPROXY=https://goproxy.cn,direct
```

# 目录结构
```sh
hello.go # main
|- study
    |- basic.go
    |- concurrency.go
    |- control.go
    |- methods.go
    |- types.go
```

# 接下来去哪？
引用自[Go 指南](https://tour.go-zh.org/concurrency/11)

> 你可以从[安装 Go]([https://go-zh.org/doc/install/]) 开始。
> 一旦安装了 Go，Go [文档](https://go-zh.org/doc/)是一个极好的 应当继续阅读的内容。 它包含了参考、指南、视频等等更多资料。

> 了解如何组织 Go 代码并在其上工作，参阅[此视频](https://www.youtube.com/watch?v=XCsL89YtqCs)，或者阅读[如何编写 Go 代码](https://go-zh.org/doc/code.html)。

> 如果你需要标准库方面的帮助，请参考[包手册](https://go-zh.org/pkg/)。如果是语言本身的帮助，阅读[语言规范](https://go-zh.org/ref/spec)是件令人愉快的事情。

> 进一步探索 Go 的并发模型，参阅 [Go 并发模型(幻灯片)](https://talks.go-zh.org/2012/concurrency.slide)以及[深入 Go 并发模型(幻灯片)](https://www.youtube.com/watch?v=QDDwwePbDtw)并阅读[通过通信共享内存](https://go-zh.org/doc/codewalk/sharemem/)的代码之旅。

> 想要开始编写 Web 应用，请参阅[一个简单的编程环境](https://vimeo.com/53221558)([幻灯片](https://talks.go-zh.org/2012/simple.slide))并阅读[编写 Web 应用](https://go-zh.org/doc/articles/wiki/)的指南。

> [函数：Go 中的一等公民](https://go-zh.org/doc/codewalk/functions/)展示了有趣的函数类型。

> [Go 博客](https://blog.go-zh.org/)有着众多关于 Go 的文章和信息。

> [mikespook 的博客](https://www.mikespook.com/tag/golang/)中有大量中文的关于 Go 的文章和翻译。

> 开源电子书 [Go Web 编程](https://github.com/astaxie/build-web-application-with-golang)和 [Go 入门指南](https://github.com/Unknwon/the-way-to-go_ZH_CN)能够帮助你更加深入的了解和学习 Go 语言。

> 访问 [go-zh.org](https://go-zh.org/) 了解更多内容。