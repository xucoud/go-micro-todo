package conf

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

type RabbitMQ struct {
	Mq string
	User string
	Password string
	Host string
	Port string
}

var Ms Mysql
var Se Service
var MQ RabbitMQ

func InitConfig() {
	file, err := ini.Load("./conf/config.ini")
	if err != nil {
		fmt.Println("加载配置问价失败！----", err)
		return
	}
	LoadServiceDAta(file)
	LoadMysqlData(file)
	LoadRabbitMQData(file)

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

func LoadRabbitMQData(file *ini.File) {
	MQ.Mq = file.Section("RabbitMQ").Key("MQ").String()
	MQ.Host = file.Section("RabbitMQ").Key("Host").String()
	MQ.Port = file.Section("RabbitMQ").Key("Port").String()
	MQ.User = file.Section("RabbitMQ").Key("User").String()
	MQ.Password = file.Section("RabbitMQ").Key("Password").String()
}