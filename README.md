# go-scaffolding
[![License](https://img.shields.io/github/license/mashape/apistatus.svg)](https://github.com/go-admin-team/go-admin)

一个基于cobra的golang脚手架，可用于快速开发现代 CLI 应用程序。



## 效果

```shell
# 默认
$ go run main.go
Welcome to use go-app v1.0.0

   ____   ____           _____  ______ ______
  / ___\ /  _ \   ______ \__  \ \____ \\____ \
 / /_/  >  <_> ) /_____/  / __ \|  |_> >  |_> >
 \___  / \____/          (____  /   __/|   __/
/_____/                       \/|__|   |__|

go-app version information:
Version    :  v1.0.0
Go version :  go1.22.0
OS / Arch  :  windows/amd64

Use "go-app -h" or "go-app --help" for more information about a command.


# 查看版本
$ go run main.go  version
Welcome to use go-app v1.0.0

   ____   ____           _____  ______ ______
  / ___\ /  _ \   ______ \__  \ \____ \\____ \
 / /_/  >  <_> ) /_____/  / __ \|  |_> >  |_> >
 \___  / \____/          (____  /   __/|   __/
/_____/                       \/|__|   |__|

go-app version information:
Version    :  v1.0.0
Go version :  go1.22.0
OS / Arch  :  windows/amd64


# 帮助信息
$ go run main.go -h
go-app

Usage:
  go-app [flags]
  go-app [command]

Available Commands:
  help        Help about any command
  version     print version info

Flags:
  -h, --help   help for go-app

Use "go-app [command] --help" for more information about a command.
```



## 组件库

- [cobra](https://github.com/spf13/cobra)：创建强大的现代 CLI 应用程序的库
- [cobra-cli](https://github.com/spf13/cobra-cli)：cobra生成器
- [viper](https://github.com/spf13/viper)： 配置管理解析库
- [zap](https://github.com/uber-go/zap)：高性能日志库
- [lumberjack](https://github.com/natefinch/lumberjack)：日志切割组件



## 本地开发

### 环境要求

go1.22.0

### 用法

```bash
# 项目初始化
git clone https://github.com/stylite1024/golang-scaffolding.git
cd golang-scaffolding
go mod tidy

# 新增命令
go install github.com/spf13/cobra-cli@latest
cobra-cli add [command]

# 修改项目名
打开go.mod文件，修改module,并且修改相关引用包

# 打包到不同平台
make all
```



## 致谢
- [cobra](https://github.com/spf13/cobra)
- [cobra-cli](https://github.com/spf13/cobra-cli)
- [viper](https://github.com/spf13/viper)
- [zap](https://github.com/uber-go/zap)
- [lumberjack](https://github.com/natefinch/lumberjack)
- [go-admin](https://github.com/go-admin-team/go-admin)



## 讨论

发现了错误? 存在某些无意义的东西? 向我发起一个[issue](https://github.com/stylite1024/golang-scaffolding/issues)吧!




## 开源协议

[MIT](https://github.com/go-admin-team/go-admin/blob/master/LICENSE.md)

Copyright (c) 2022 stylite1024
