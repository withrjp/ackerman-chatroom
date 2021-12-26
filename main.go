package main

import (
	"github.com/gin-gonic/gin"
)

func main() {
	hub := newHub()
	go hub.run()
	r := gin.Default()
	r.GET("/ws", func(c *gin.Context) {
		serveWs(hub, c.Writer, c.Request)
	})
	r.Run()
}
