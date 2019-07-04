package chat

import (
	"encoding/json"
	"fmt"
	"io/ioutil"
	"net/http"
	"regexp"
	"strconv"
	"strings"

	"github.com/gin-gonic/gin"
	"github.com/gofrs/uuid"
	"github.com/gorilla/websocket"
)

type ClientManager struct {
	clients    map[*Client]bool // conn map by pointer key
	broadcast  chan []byte      // broadcast channel
	register   chan *Client
	unregister chan *Client
}

var Manager = ClientManager{
	broadcast:  make(chan []byte),
	register:   make(chan *Client),
	unregister: make(chan *Client),
	clients:    make(map[*Client]bool),
}

type Message struct {
	Sender  string `json:"sender,omitempty"` // Sender id
	Event   string `json:"event,omitempty"`  // Message event. ex: chat, connCount
	Content string `json:"content,omitempty"`
}

// listen register and broadcast channel
func (Manager *ClientManager) Start() {
	for {
		select {
		case conn := <-Manager.register:
			Manager.clients[conn] = true
			WriteCount(conn)

		case conn := <-Manager.unregister:
			if _, ok := Manager.clients[conn]; ok {

				close(conn.send)
				delete(Manager.clients, conn)
				WriteCount(conn)
			}
		case message := <-Manager.broadcast:
			for conn := range Manager.clients {
				select {
				case conn.send <- message:
				}
			}
		}
	}
}

// broadcast to all client
func (Manager *ClientManager) AllSend(message []byte) {

	for conn := range Manager.clients {
		conn.send <- message
	}
}

// Chatroom : websocket new conntion
func Chatroom(c *gin.Context) {
	conn, err := (&websocket.Upgrader{CheckOrigin: func(r *http.Request) bool { return true }}).Upgrade(c.Writer, c.Request, nil)
	if err != nil {
		fmt.Println("new client error: " + err.Error())
		return
	}

	// get uuid
	id, _ := uuid.NewV4()

	// get chinese nikename by uuid
	reId := strings.Replace(id.String(), "-", "", -1)
	client := &Client{nic: getNikeName(reId[0:20]), socket: conn, send: make(chan []byte)}

	// into Manger.register channel
	Manager.register <- client

	// websocket go read and write
	go client.Read()
	go client.Write()
}

// push online count to client
func WriteCount(conn *Client) {
	jsonMessage, _ := json.Marshal(&Message{Event: "connCount", Content: strconv.Itoa(len(Manager.clients))})
	Manager.AllSend(jsonMessage)
}

// get chinese nikename
func getNikeName(uuid string) string {
	resp, err := http.Get("https://wtf.hiigara.net/api/run/titlegen/" + uuid + "?event=ManualRun")
	if err != nil {
		return uuid
	}
	body, err := ioutil.ReadAll(resp.Body)
	if err != nil {
		return uuid
	}
	u := map[string]interface{}{}
	err = json.Unmarshal(body, &u)
	if err != nil {
		return uuid
	}
	re, _ := regexp.Compile("em=](.*)\\[\\/=em")
	str := fmt.Sprintf("%v", u["text"])
	nic := re.FindStringSubmatch(str)
	return nic[1]
}
