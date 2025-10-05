// 读取配置：1.读取命令行参数并绑定；2.读取yaml配置文件相关配置
// core/init_conf.go
package core

import (
	"blogx_backend/conf"
	"blogx_backend/flags"
	"fmt"
	"gopkg.in/yaml.v3"
	"os"
)

func ReadConf() (c *conf.Config) {
	byteData, err := os.ReadFile(flags.FlagOptions.File)
	if err != nil {
		panic(err)
	}

	c = new(conf.Config)
	err = yaml.Unmarshal(byteData, c)
	if err != nil {
		panic(fmt.Sprintf("ymal file format error: %s", err))
	}

	fmt.Printf("读取配置文件 %s 成功\n", flags.FlagOptions.File)
	fmt.Println(*c)

	return c
}
