##代码组织结构

####概述

* go 项目存储于一个工作空间workspace
* 一个工作空间包含多个仓库repo
* 各个仓库包含一个或多个包package
* 各个包由单一目录下的一个或多个.go文件组成
* 包的目录路径作为引用的路径，即import + 目录地址

####工作空间

工作空间是一个目录（该目录指定目录结构）,
包含src 和 bin子目录,src存放.go源文件，bin存放可执行命令

使用go build ， go install 会创建可执行文件到bin目录

src 目录一般包含多个代码仓库, 如下面例子
```
bin/
   hello     # executable command
   output    # executable command
src/
  github.com/golang/example
    .git/
    hello/
      hello.go
    output/
      main.go
      main_est.go
    stringutil/
      reverse.go
      reverse_test.go
  golang.org/x/image/
    .git/
    bmp/
      reader.go
      writer.go
  ....
```
上述例子，工作空间包含了两个仓库(example 和 image). example仓库包含
了两个命令(hello和output) 和一个library(stringutil).image仓库包含了bmp包和其他包

一个典型的工作空间包含多个仓库（包含多个包和命令）。大多数Go程序员使用单一的工作空间来开发多个项目

####NOTE: 不要用symbolic links 去链接工作空间中的文件或目录

####GOPATH
GOPATH环境变量指定了工作空间的目录，默认指向home目录下的一个名为go的文件目录. Unix下为$HOME/go

如果你想要建立不同的目录，你需要设置GOPATH 指向那个目录.
注意GOPATH目录与go的安装目录不一定是相同的

可以使用go env命令去查看当前的GOPATH.
如果没有设置环境变量，则输出默认的目录地址

为了方便起见，需要将工作空间下的bin目录加入到PATH环境变量中
$ export PATH=$PATH:$(go env GOPATH)/bin
$ export GOPATH=$(go env GOPATH)

想要学习GOPATH变量的详细细节，可以使用go help gopath

####Import paths
import path是一个可以唯一标识包的字符串。包的import
path是它在工作空间下的目录或一个远程的仓库

标准库中的包以短路径的方式引用，如"fmt" 或 "net/http",
对于你自己的包，必须选择一个基本路径base path .
该路径不能与标准库未来可能添加的组件或其他的外部依赖相冲突

如果你有一个仓库，那么你应用这个仓库的源地址作为base path

####注意：在build代码时不需要事先提交到远程的仓库. 实际上你可以选择
任意的路径名,只要与go生态中的其他library不冲突就好


假设我们以github.com/user作为我们的base path.在工作空间中创建一个目
录
$ mkdir -p $GOPATH/src/github.com/user

写一个简单的程序。首先选择包路径为github.com/user/hello 并且创建
一个目录
$ mkdir $GOPATH/src/github.com/user/hello

接下来创建一个hello目录下创建一个hello.go文件.
```
package main
import "fmt"
func main() {
  fmt.Println("hello world")
}
```
现在你可以build 和 install
$ go install github.com/user/hello

####注意 你可以从任何目录下运行上面的命令, go tool
会从GOPATH目录下 github.com/user/hello目录查找源码

你也可以从包目录下直接运行go install
$ cd $GOPATH/src/github.com/user/hello
$ go install

这个命令会build出hello命令程序,创建一个可执行的二进制文件。然后安装这个二进制文件到工作空间的bin目录下,命名为hello.

go tool 在build 和 install时 仅会输出错误信息,
如果没有错误信息则执行成功。

然后你可以运行命令如下
$ $GOPATH/bin/hello
hello, world.

或者你可以把$GOPATH/bin 加到PATH变量中.


