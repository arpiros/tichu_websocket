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
	http.HandleFunc("/", PreProtocolProcess)
	err := http.ListenAndServe(viper.GetString("http_addr"), nil)
	if err != nil {
		logrus.Fatal("ListenAndServe: ", err)
	}
}

func PreProtocolProcess(w http.ResponseWriter, r *http.Request) {
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
			LogOut(ws)
			break
		}

		var base protocol.RequestBase
		err = json.Unmarshal(message, &base)
		if err != nil {
			logrus.Println("Unmarshal Error : ", err)
			LogOut(ws)
			break
		}

		models.AddUser(ws)

		ProtocolProcess(base, ws, message)
	}
}

func ProtocolProcess(base protocol.RequestBase, ws *websocket.Conn, message []byte) {
	switch base.ProtocolType {
	case protocol.CreateRoom:
		controllers.CreateRoom(ws, message)
	case protocol.JoinRoom:
		controllers.JoinRoom(ws, message)
	case protocol.CallLargeTichu:
		controllers.CallLargeTichu(ws, message)
	case protocol.ChangeCard:
		controllers.ChangeCard(ws, message)
	case protocol.MoveTurn:
		controllers.MoveTurn(ws, message)
	default:
		logrus.Warnf("Not Found Protocol : %d", base.ProtocolType)
	}
}

func LogOut(ws *websocket.Conn) {
	user, err := models.GetUser(ws)
	if err != nil {

	}

	models.LeaveRoom(ws, user.RoomCode)
	models.DelUser(ws)
}