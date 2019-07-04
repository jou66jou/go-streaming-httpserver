package chat

import (
	"encoding/json"

	"github.com/gorilla/websocket"
)

type Client struct {
	nic    string //nikename
	socket *websocket.Conn
	send   chan []byte //save message
}

// read client's message
func (c *Client) Read() {
	defer func() {
		Manager.unregister <- c
		c.socket.Close()
	}()

	for {
		_, message, err := c.socket.ReadMessage()
		if err != nil {
			Manager.unregister <- c
			c.socket.Close()
			break
		}
		jsonMessage, _ := json.Marshal(&Message{Sender: c.nic, Content: string(message), Event: "chat"})
		// fmt.Println(c.id + " : " + string(message))
		Manager.broadcast <- jsonMessage
	}
}

// write message to client
func (c *Client) Write() {
	defer func() {
		c.socket.Close()
	}()

	for {
		select {
		case message, ok := <-c.send:
			if !ok {
				c.socket.WriteMessage(websocket.CloseMessage, []byte{})
				return
			}

			c.socket.WriteMessage(websocket.TextMessage, message)
		}
	}
}
