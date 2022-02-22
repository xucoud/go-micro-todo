package config

import (
	"fmt"
	"github.com/go-ini/ini"
)

type Service struct {
	AppMode string
	HttpPort string
}

type Mysql struct {
	Db string
	DbHost string
	DbPort string
	DbUser string
	DbPassword string
	DbName string
}

var Ms Mysql
var Se Service

func InitConfig() {
	file, err := ini.Load("./config/config.ini")
	if err != nil {
		fmt.Println("加载配置问价失败！----", err)
		return
	}
	LoadServiceDAta(file)
	LoadMysqlData(file)
}

func LoadMysqlData(file *ini.File) {
	Ms.DbName = file.Section("Mysql").Key("DbName").String()
	Ms.Db = file.Section("Mysql").Key("Db").String()
	Ms.DbHost = file.Section("Mysql").Key("DbHost").String()
	Ms.DbPort = file.Section("Mysql").Key("DbPort").String()
	Ms.DbUser = file.Section("Mysql").Key("DbUser").String()
	Ms.DbPassword = file.Section("Mysql").Key("DbPassword").String()
}

func LoadServiceDAta(file *ini.File) {
	Se.AppMode = file.Section("Service").Key("AppMode").String()
	Se.HttpPort = file.Section("Service").Key("HttpPort").String()
}