package websocket

import (
	"fmt"
	"log"
	"github.com/gorilla/websocket"
)

type Client struct {
	ID string
	Conn *websocket.Conn
	Pool *Pool
}


type Message struct {
    Type int    `json:"type"`
    Body string `json:"body"`
}

// Declared function on Client Types
func (c *Client) Read() {
	defer func() { // Ensure clean up
		c.Pool.Unregister <- c
		c.Conn.Close()
	}()

	for {
		messageType, p, err := c.Conn.ReadMessage()
		if err != nil {
			log.Println(err)
			return
		}

		message := Message{Type: messageType, Body: string(p)}
		c.Pool.Broadcast <- message  // Broadcast message to every client in Pool
		fmt.Printf("Mesage Received: %+v\n", message)
	}
} 