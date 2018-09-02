package main

import (
	"flag"
	"log"
	"net/url"
	"os"
	"os/signal"
	"github.com/gorilla/websocket"
	"time"
	"tichu/tichu_websocket/protocol"
	"fmt"
)

var addr = flag.String("addr", "localhost:1015", "http service address")

func main() {
	flag.Parse()
	log.SetFlags(0)

	interrupt := make(chan os.Signal, 1)
	signal.Notify(interrupt, os.Interrupt)

	u := url.URL{Scheme: "ws", Host: *addr, Path: "/"}
	log.Printf("connecting to %s", u.String())

	c, _, err := websocket.DefaultDialer.Dial(u.String(), nil)
	if err != nil {
		log.Fatal("dial:", err)
	}
	defer c.Close()

	done := make(chan struct{})

	go func() {
		defer close(done)
		//for {
			var createRoom protocol.CreateRoom
			err := c.ReadJSON(&createRoom)
			if err != nil {
				log.Println("read:", err)
				return
			}
			log.Printf("recv: %v", createRoom.ProtocolType)
		//}
	}()

	ticker := time.NewTicker(time.Second)
	defer ticker.Stop()

	for {
		select {
		case <-done:
			return
		case t := <-ticker.C:
			//var createRoom protocol.CreateRoom
			//createRoom.ProtocolType = protocol.CREATE_ROOM
			//err := c.WriteJSON(createRoom)
			//if err != nil {
			//	log.Println("write:", err)
			//	return
			//}
			//
			fmt.Printf("t : %v\n", t.String())

			var joinRoom protocol.Join
			joinRoom.RoomNumber = 0
			err = c.ReadJSON(&joinRoom)
			if err != nil {
				log.Println("read:", err)
				return
			}
			log.Printf("recv: %v", joinRoom.ProtocolType)
		case <-interrupt:
			log.Println("interrupt")

			// Cleanly close the connection by sending a close message and then
			// waiting (with timeout) for the server to close the connection.
			err := c.WriteMessage(websocket.CloseMessage, websocket.FormatCloseMessage(websocket.CloseNormalClosure, ""))
			if err != nil {
				log.Println("write close:", err)
				return
			}
			select {
			case <-done:
			case <-time.After(time.Second):
			}
			return
		}
	}
}
