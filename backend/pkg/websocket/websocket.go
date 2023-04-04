package websocket

import (
	"log"
	"net/http"

	"github.com/gorilla/websocket"
)

var upgrader = websocket.Upgrader{
	ReadBufferSize:  1024,
	WriteBufferSize: 1024,
	// For checking origin of requests:
	// Allow all origins of incoming connection requests.
	// As this projects gets more complex I can upgrade later.
	CheckOrigin:     func(r *http.Request) bool { return true },
}

func Upgrade(w http.ResponseWriter, r *http.Request) (*websocket.Conn, error) {
	conn, err := upgrader.Upgrade(w, r, nil)
	if err != nil {
		log.Println(err)
		return nil, err
	}
	return conn, nil
}

// // Define a reader to listen for new messages sent to our Websocket endpoint
// func Reader(conn *websocket.Conn) {
// 	for {
// 		messageType, p, err := conn.ReadMessage()
// 		if err != nil {
// 			log.Println(err)
// 			return
// 		}

// 		fmt.Println(string(p)) // Read meassage, p

// 		if err2 := conn.WriteMessage(messageType, p); err2 != nil { // Write same message back to confirm
// 			log.Println(err2) // log any errors arising during writing back step
// 		}
// 	}
// }

// func Writer(conn *websocket.Conn) {
// 	for { // Loop until an error from no more messages
// 		fmt.Println("Sending...")
// 		messageType, r, err := conn.NextReader()

// 		if err != nil {
// 			fmt.Println(err)
// 		}

// 		w, err := conn.NextWriter(messageType)

// 		if err != nil {
// 			fmt.Println(err)
// 		}

// 		if _, err := io.Copy(w, r); err != nil {
// 			fmt.Println(err)
// 		}

// 		if err := w.Close(); err != nil {
// 			fmt.Println(err)
// 			return
// 		}

// 	}
// }
