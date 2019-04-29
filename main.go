package main

import (
	"flag"
	"github.com/gin-gonic/gin"
	log "github.com/sirupsen/logrus"
	"os"
	"peak/config"
	"peak/models"
	"peak/routers"
)

func init() {
	log.SetFormatter(&log.TextFormatter{})
	log.SetOutput(os.Stdout)
	log.SetLevel(log.WarnLevel)
}

func main() {
	//初始化路由
	router := gin.Default()

	//注册路由
	routers.RegisterRoute(router)

	//读取配置文件
	configFilePath := flag.String("C", "config/config.yaml", "config file path")
	err := config.LoadConfiguration(*configFilePath);
	if err != nil {
		log.Error("err parsing config log file", err)
		return
	}
	//初始化数据库连接

	models.Connect()
	//运行
	router.Run() // listen and serve on 0.0.0.0:8080
}
