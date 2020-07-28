# toolkit 

## 作用
用于提供工作效率的工具箱，里面有各种工具，就比如真实工具箱中里面有扳手，各种大小的起子，钳子等
某些场景下确实可以达到事半功倍的效果

## 安装
### 源码安装
有 `go` 语言环境的可以直接用源码进行编译运行

```
git clone https://github.com/russellgao/toolkit.git
cd toolkit
make
```

### 二进制
可以直接在release 页面进行下载对应的操作系统的二进制文件

https://github.com/russellgao/toolkit/releases/

## 用法
###  本机运行
可以通过如下命令进行
```
gwz:toolkit gaoweizong$ tkctl --help

tkctl is a toolkit entrypoint,run `tkctl --help` get more information.

Usage:
  tkctl [flags]
  tkctl [command]

Available Commands:
  help        Help about any command
  replace     文本替换，支持正则替换和非正则替换，类似与linux下的sed，但比sed更好用，而且可以跨平台使用
  secret      生成随机密码，支持1～100位长度，可以指定是否包含特殊字符
  version     tkctl version

Flags:
  -h, --help      help for tkctl
  -v, --version   show the version and exit

Use "tkctl [command] --help" for more information about a command.

```

tkctl 中的子命令会不断更新，某个具体的功能请查看`Available Commands:`下的帮助文档，如文本替换

```shell script
 tkctl replace --help 
文本替换，支持正则替换和非正则替换，类似与linux下的sed，但比sed更好用，而且可以跨平台使用

Usage:
  tkctl replace [flags]

Flags:
  -d, --dirs string      需要替换的目录, 默认为当前路径 (default ".")
  -h, --help             help for replace
  -m, --mode string      替换的模式，支持正则（regexp）和非正则（text）两种模式，默认非正则， (default "text")
  -p, --pattern string   需要替换的pattern  [required]
  -r, --repl string      目标字符串  [required]
```

### docker
如果本地有docker环境，也可以不用下载二进制的制品，可以通过docker 环境直接运行

```shell script
docker run -it --rm russellgao/toolkit:latest tkctl --help

# 如果有需要可以把目录挂载进去
docker run -it -v /data:/data --rm russellgao/toolkit:latest tkctl --help
```

## 适用范围
可以跨平台使用

- mac
- windows
- linux

## 开发环境
- go 1.14.2

## 支持的功能
### 1.0.0
- 文本正则替换
- 生成随机密码

## 未来展望
- 期望可以成为一个完整的工具箱，可以解决日常工作中的繁杂事情。

## 关于作者
云原生爱好者，专注于devops，ops，kubernetes，云原生，算法等领域，可以关注作者个人公众号
，不定期会有干货分享，可以及时推送给您，欢迎随时交流哟！

![](russellgao.jpg)

