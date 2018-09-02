package local

import (
	"net/http"
	"github.com/spf13/viper"
	"log"
	"github.com/gorilla/websocket"
	"tichu/tichu_websocket/protocol"
	"encoding/json"
	"tichu/tichu_websocket/controllers"
)

var upgrader = websocket.Upgrader{} // use default options

func StartRouter() {
	http.HandleFunc("/", ProtocolProcess)
	http.ListenAndServe(viper.GetString("http_addr"), nil)
}

func ProtocolProcess(w http.ResponseWriter, r *http.Request) {
	c, err := upgrader.Upgrade(w, r, nil)
	if err != nil {
		log.Print("upgrade:", err)
		return
	}

	defer c.Close()
	for {
		_, message, err := c.ReadMessage()
		if err != nil {
			log.Println("ReadMessage Error : ", err)
			break
		}

		var base protocol.Base
		err = json.Unmarshal(message, &base)
		if err != nil {
			log.Println("Unmarshal Error : ", err)
			break
		}

		switch base.ProtocolType {
		case protocol.CREATE_ROOM:
			var createRoom protocol.CreateRoom
			err = json.Unmarshal(message, &createRoom)
			controllers.CreateRoom(createRoom, c)
		case protocol.JOIN:
			var join protocol.Join
			err = json.Unmarshal(message, &join)
			controllers.Join(join, c)
		}
	}
}