package core

import (
	"fmt"
	"gopkg.in/yaml.v3"
	"os"
)

var confPath = "settings.yaml"

type System struct {
	IP   string `yaml:"ip"`
	Port string `yaml:"port"`
}

type Config struct {
	System System `yaml:"system"`
}

func ReadConf() {
	byteData, err := os.ReadFile(confPath)
	if err != nil {
		panic(err)
	}

	var conf Config
	err = yaml.Unmarshal(byteData, &conf)
	if err != nil {
		panic(fmt.Sprintf("ymal file format error: %s", err))
	}

	fmt.Println(conf)
}
