package main

import (
	"blogx_backend/core"
	"blogx_backend/flags"
	"blogx_backend/global"
	"github.com/sirupsen/logrus"
)

func main() {
	flags.Parse()
	global.Config = core.ReadConf()
	core.InitLogrus()

	logrus.Warnf("warning...")
	logrus.Debug("Debug...")
	logrus.Error("error...")
	logrus.Infof("info...")
}
