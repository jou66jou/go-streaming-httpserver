package routers

import (
	"github.com/gin-gonic/gin"
	"github.com/jou66jou/go-livetest/controllers/chat"
	"github.com/jou66jou/go-livetest/controllers/webpage"
)

func Routers(r *gin.Engine) *gin.Engine {

	//	web page get
	r.GET("/webpush", webpage.PushPage)
	r.GET("/webplay", webpage.PullPage)
	r.GET("/mplaym3u8", webpage.Pushm3u8)
	r.GET("/mplayflv", webpage.Pushflv)
	r.GET("/webpush2", webpage.PushPage2)
	r.GET("/chatroom", webpage.Chatroom)
	//	websocket
	r.GET("/chat", chat.Chatroom)
	return r
}
