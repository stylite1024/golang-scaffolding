package common

import (
	"fmt"
	"go-app/pkg/tools"
	"runtime"
)

var (
	// Version Info
	AppName = "go-app"
	Version = "v1.0.0"
	// LogoContent go-admin ascii显示，减少静态文件依赖
	// 设计tagg: https://patorjk.com/software/taag
	// ascii转bytes: https://www.toolhelper.cn/EncodeDecode/ByteArray  https://onlinetools.com/ascii/convert-ascii-to-bytes
	LogoContent = []byte{10, 32, 32, 32, 32, 32, 32, 32, 32, 32, 32, 32, 32, 32, 32, 32, 32, 32, 32, 32, 32, 32, 32, 32, 32, 32, 32, 32, 32, 32, 32, 32, 32, 32, 32, 32, 32, 32, 32, 32, 32, 32, 32, 32, 32, 32, 32, 32, 10, 32, 32, 32, 95, 95, 95, 95, 32, 32, 32, 95, 95, 95, 95, 32, 32, 32, 32, 32, 32, 32, 32, 32, 32, 32, 95, 95, 95, 95, 95, 32, 32, 95, 95, 95, 95, 95, 95, 32, 95, 95, 95, 95, 95, 95, 32, 32, 10, 32, 32, 47, 32, 95, 95, 95, 92, 32, 47, 32, 32, 95, 32, 92, 32, 32, 32, 95, 95, 95, 95, 95, 95, 32, 92, 95, 95, 32, 32, 92, 32, 92, 95, 95, 95, 95, 32, 92, 92, 95, 95, 95, 95, 32, 92, 32, 10, 32, 47, 32, 47, 95, 47, 32, 32, 62, 32, 32, 60, 95, 62, 32, 41, 32, 47, 95, 95, 95, 95, 95, 47, 32, 32, 47, 32, 95, 95, 32, 92, 124, 32, 32, 124, 95, 62, 32, 62, 32, 32, 124, 95, 62, 32, 62, 10, 32, 92, 95, 95, 95, 32, 32, 47, 32, 92, 95, 95, 95, 95, 47, 32, 32, 32, 32, 32, 32, 32, 32, 32, 32, 40, 95, 95, 95, 95, 32, 32, 47, 32, 32, 32, 95, 95, 47, 124, 32, 32, 32, 95, 95, 47, 32, 10, 47, 95, 95, 95, 95, 95, 47, 32, 32, 32, 32, 32, 32, 32, 32, 32, 32, 32, 32, 32, 32, 32, 32, 32, 32, 32, 32, 32, 32, 32, 92, 47, 124, 95, 95, 124, 32, 32, 32, 124, 95, 95, 124, 32, 32, 32, 32, 10}
)

func Tip() {
	tipStr := `Welcome to use ` + tools.Green(AppName+" "+Version)
	info := fmt.Sprintf(`%s version information:
Version    :  %s
Go version :  %s
OS / Arch  :  %s/%s
`, AppName, Version, runtime.Version(), runtime.GOOS, runtime.GOARCH)
	fmt.Printf("%s", tipStr)
	fmt.Println(tools.Green(string(LogoContent)))
	fmt.Println(info)
}

func HelpTip()  {
	helpStr := tools.Cyan(`Use "go-app -h" or "go-app --help" for more information about a command.`)
	fmt.Printf("%s\n", helpStr)
}
