package user

import (
	"go_code/go-micro-todo/user/config"
	"go_code/go-micro-todo/user/model"
)

func Init() {
	config.InitConfig()
	model.InitDB()
}

func main() {
	defer model.Db.Close()
}