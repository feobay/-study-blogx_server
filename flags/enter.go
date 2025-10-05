// flags/enter.go
// 绑定参数
package flags

import (
	"flag"
)

type Options struct {
	File    string
	DB      bool // go run main.go -db 迁移数据库
	Version bool // go run main.go -version 指定运行版本
}

var FlagOptions = new(Options) // 入口参数

func Parse() {
	// 以下绑定操作都是将配置的值写入给第一个入参
	// -f filename 设置启动时读取的配置文件（默认为settings.yaml）
	flag.StringVar(&FlagOptions.File, "f", "settings.yaml", "配置文件") // f参数名称(也就是说用的时候直接-f即可)，绑定值默认为 `settings.yaml`, 用法标签设置为 “配置文件”

	// -db 启用即为标志要进行数据库迁移，默认为关闭
	flag.BoolVar(&FlagOptions.DB, "db", false, "数据库迁移")

	// -v 指定版本
	flag.BoolVar(&FlagOptions.Version, "v", false, "指定版本")
	flag.Parse() // 前面是绑定参数并设定默认值，这一行解析是将命令行的参数覆盖赋值给对应变量
	//fmt.Println("FlagOptions is", *FlagOptions)
}
