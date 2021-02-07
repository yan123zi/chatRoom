package main

import (
	"chatRoom/server/app"
	"chatRoom/server/config"
	"flag"
	"github.com/jinzhu/gorm"
	"github.com/sirupsen/logrus"
)

var configFilePath = flag.String("config", "./chat-room.yaml", "配置文件")

func init() {
	flag.Parse()
	_, err := config.InitAppConfig(*configFilePath)
	if err != nil {
		logrus.Errorf("init app err:%s", err.Error())
		return
	}
	db,err:=gorm.Open(config.AppConfig.MysqlUrl)
	if err != nil {
		logrus.Errorf("init app db err:%s",err.Error())
		return
	}
	config.Db=db
}

func main() {
	app.InitIris()
}
