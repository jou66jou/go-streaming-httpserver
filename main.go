package main

import (
	"log"
	"time"

	"github.com/gin-gonic/gin"
	"github.com/jou66jou/go-livetest/controllers/chat"
	"github.com/jou66jou/go-livetest/routers"
)

var (
	port = "8080"
)

func main() {
	defer func() {
		if r := recover(); r != nil {
			log.Println("main panic: ", r)
			time.Sleep(1 * time.Second)
		}
	}()

	// starts an chatroom service
	go chat.Manager.Start()

	router := gin.New()
	router.Use(gin.Logger())
	router.LoadHTMLGlob("view/*.html")
	router.Static("/static", "static")

	// register router
	routers.Routers(router)
	// router.RunTLS(":"+port, "./certs/server.crt", "./certs/server.key")
	router.Run(":" + port)
}
