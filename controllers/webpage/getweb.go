package webpage

import (
	"net/http"

	"github.com/gin-gonic/gin"
)

func PullPage(c *gin.Context) {
	c.HTML(http.StatusOK, "webplay.html", nil)
}

func PushPage(c *gin.Context) {
	c.HTML(http.StatusOK, "webpush.html", nil)
}
func PushPage2(c *gin.Context) {
	c.HTML(http.StatusOK, "webpush2.html", nil)
}
func Pushm3u8(c *gin.Context) {
	c.HTML(http.StatusOK, "mplay_m3u8.html", nil)

}
func Pushflv(c *gin.Context) {
	c.HTML(http.StatusOK, "mplay_flv.html", nil)

}

func Chatroom(c *gin.Context) {
	c.HTML(http.StatusOK, "chat.html", nil)

}
