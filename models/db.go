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
	addrs := fmt.Sprintf("%s:%s", config.MONGO.HOST, config.MONGO.PORT)
	dialInfo := &mgo.DialInfo{Addrs: []string{addrs}, Source: config.MONGO.DB, Username: config.MONGO.USER, Password: config.MONGO.PASS}
	DB_SESSION, ERR = mgo.DialWithInfo(dialInfo)
	if ERR != nil {
		panic("连接数据库发生错误！")
	}
	DB = DB_SESSION.DB(config.MONGO.DB)
}
