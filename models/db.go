package models

import (
	"fmt"
	"github.com/globalsign/mgo"
	config2 "peak/config"
)

var DB_SESSION *mgo.Session
var DB *mgo.Database
var ERR error

func Connect() {
	config := config2.GetConfiguration()
	addrs := fmt.Sprintf("%s:%s", config.MONGOHOST, config.MONGOPORT)
	dialInfo := &mgo.DialInfo{Addrs: []string{addrs}, Source: config.MONGODB, Username: config.MONGOUSER, Password: config.MONGOPASS}
	DB_SESSION, ERR = mgo.DialWithInfo(dialInfo)
	if ERR != nil {
		panic("连接数据库发生错误！")
	}
	DB = DB_SESSION.DB(config.MONGODB)
}
