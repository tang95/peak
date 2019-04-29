package main

import (
	"flag"
	"fmt"
	"github.com/gin-gonic/gin"
	"github.com/gobuffalo/packr"
	log "github.com/sirupsen/logrus"
	"os"
	"peak/config"
	"peak/models"
	"peak/routers"
)

func init() {

	//读取配置文件
	configFilePath := flag.String("C", "config/config.yaml", "config file path")
	flag.Parse()
	err := config.LoadConfiguration(*configFilePath)
	if err != nil {
		log.Error("err parsing config log file", err)
		return
	}

	log.SetFormatter(&log.TextFormatter{
		TimestampFormat: "2006-01-02 15:04:05",
		FullTimestamp:   true,
	})
	log.SetOutput(os.Stdout)
	log.SetLevel(log.InfoLevel)
}

func main() {
	conf := config.GetConfiguration()
	gin.SetMode(conf.SERVER.MODE)
	//初始化路由
	router := gin.Default()

	//注册路由
	routers.RegisterRoute(router)

	//模版文件
	staticBox := packr.NewBox("./templates")
	router.StaticFS("/web", staticBox)
	//初始化数据库连接

	models.Connect()
	//运行
	addr := fmt.Sprintf("%s:%s", conf.SERVER.HOST, conf.SERVER.PORT)
	log.Infof("Listening and serving HTTP on %s\n", addr)
	router.Run(addr)
}
