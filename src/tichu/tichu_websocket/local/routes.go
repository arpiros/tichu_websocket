package local

import (
	"net/http"
	"github.com/spf13/viper"
	"github.com/gorilla/websocket"
	"tichu/tichu_websocket/protocol"
	"encoding/json"
	"tichu/tichu_websocket/models"
	"tichu/tichu_websocket/controllers"
	"github.com/Sirupsen/logrus"
)

var upgrader = websocket.Upgrader{} // use default options

func StartRouter() {
	http.HandleFunc("/", ProtocolProcess)
	err := http.ListenAndServe(viper.GetString("http_addr"), nil)
	if err != nil {
		logrus.Fatal("ListenAndServe: ", err)
	}
}

func ProtocolProcess(w http.ResponseWriter, r *http.Request) {
	upgrader.CheckOrigin = func(r *http.Request) bool { return true }
	ws, err := upgrader.Upgrade(w, r, nil)
	if err != nil {
		logrus.Print("upgrade:", err)
		return
	}

	defer ws.Close()
	for {
		_, message, err := ws.ReadMessage()
		if err != nil {
			logrus.Println("ReadMessage Error : ", err)
			models.DelUser(ws)
			break
		}

		var base protocol.Base
		err = json.Unmarshal(message, &base)
		if err != nil {
			logrus.Println("Unmarshal Error : ", err)
			models.DelUser(ws)
			break
		}

		models.AddUser(ws)

		switch base.ProtocolType {
		case protocol.CREATE_ROOM:
			controllers.CreateRoom(ws, message)
		case protocol.JOIN:
			controllers.JoinRoom(ws, message)
		default:
			logrus.Warnf("Not Found Protocol : %d", base.ProtocolType)
		}
	}
}
