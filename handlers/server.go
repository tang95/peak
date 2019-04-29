package handlers

import (
	"fmt"
	"github.com/gin-gonic/gin"
	"github.com/googollee/go-socket.io"
	"github.com/sirupsen/logrus"
)

var SocketIOServer *socketio.Server

func init() {
	var err error
	SocketIOServer, err = socketio.NewServer(nil)
	if err != nil {
		logrus.Error(err)
	}
	SocketIOServer.OnConnect("/", func(s socketio.Conn) error {
		s.SetContext("")
		fmt.Println("connected:", s.ID())
		return nil
	})
}

func SocketIOHandler() gin.HandlerFunc {
	return func(c *gin.Context) {
		origin := c.GetHeader("Origin")
		c.Header("Access-Control-Allow-Credentials", "true")
		c.Header("Access-Control-Allow-Origin", origin)
		SocketIOServer.ServeHTTP(c.Writer, c.Request)
	}
}
