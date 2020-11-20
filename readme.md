VSCode安装GO语言依赖工具
===
> 来源：[码农飞龙](https://www.jianshu.com/p/f952042af8ff)

|   |   |   |
|---|---|---|
|名称|描述|链接|
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

1. 创建目录$GOPATH/src/golang.org/x，并切换到该目录
```sh
mkdir -p $GOPATH/src/golang.org/x/
cd $GOPATH/src/golang.org/x/
```
2. 克隆golang.org工具源码
  > 如果不克隆的话，go get -u -v golang.org/xxx肯定是timeout的，所以只能先把它们下载到本地src/golang.org/x/tools
```sh
git clone https://github.com/golang/tools.git
git clone https://github.com/golang/lint.git
```
3. 下载github源码
  > 按照go get -u -v命令，从GitHub上下载代码后还会fetch，我们很可能会在fetch https://golang.org/xxx的时候挂掉，原因你懂的。所以去掉-u选项，禁止从网络更新现有代码。
```sh
# 先从github下载依赖工具的源码，fetch提示timeout不要管
go get -v github.com/ramya-rao-a/go-outline
go get -v github.com/acroca/go-symbols
go get -v github.com/mdempsky/gocode
go get -v github.com/rogpeppe/godef
go get -v github.com/zmb3/gogetdoc
go get -v github.com/fatih/gomodifytags
go get -u github.com/sqs/goreturns
go get -v github.com/cweill/gotests/...
go get -v github.com/josharian/impl
go get -v github.com/haya14busa/goplay/cmd/goplay
go get -v github.com/uudashr/gopkgs/cmd/gopkgs
go get -v github.com/davidrjenni/reftools/cmd/fillstruct
go get -v github.com/alecthomas/gometalinter
```
4. 安装工具
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
  > GOPROXY
  ```sh
  export GOPROXY=https://goproxy.io
  ```
5. 集成到系统环境中
  > 由于我是在用户目录下临时安装的，真正的GO环境是/usr/local/go，所以最后一步无比注意，把~/go/bin下面生成的所有执行文件拷贝到系统环境中
```sh
sudo cp -af $GOPATH/bin/* /usr/local/go/bin/
```
6. 安装问题
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

语法问题
===

```go
s := []int{2, 3, 5, 7, 11, 13}

s = s[:0] // len=0 cap=6 []
s = s[:4] // len=3 cap=5 [3 5 7]
s = s[2:] // len=1 cap=3 [7]
```