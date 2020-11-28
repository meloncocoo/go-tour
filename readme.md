VSCode安装GO语言依赖工具
===

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

安装步骤
---

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

3. 目录结构
```sh
hello.go # main
|- study
    |- basic.go
    |- concurrency.go
    |- control.go
    |- methods.go
    |- types.go
```