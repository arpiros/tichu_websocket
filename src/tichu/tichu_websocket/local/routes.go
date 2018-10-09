package local

import (
	"encoding/json"
	"github.com/Sirupsen/logrus"
	"github.com/gorilla/websocket"
	"github.com/spf13/viper"
	"net/http"
	"tichu/tichu_websocket/controllers"
	"tichu/tichu_websocket/models"
	"tichu/tichu_websocket/protocol"
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
	switch base.RequestType {
	case protocol.ReqCreateRoom:
		controllers.CreateRoom(ws, message)
	case protocol.ReqJoinRoom:
		controllers.JoinRoom(ws, message)
	case protocol.ReqCallLargeTichu:
		controllers.CallLargeTichu(ws, message)
	case protocol.ReqChangeCard:
		controllers.ChangeCard(ws, message)
	case protocol.ReqCallTichu:
		controllers.CallTichu(ws, message)
	case protocol.ReqSubmitCard:
		controllers.SubmitCard(ws, message)
	case protocol.ReqMoveTurn:
		controllers.MoveTurn(ws, message)
	case protocol.ReqTurnPass:
		controllers.TurnPass(ws, message)
	default:
		logrus.Warnf("Not Found Protocol : %d", base.RequestType)
	}
}

func LogOut(ws *websocket.Conn) {
	user, err := models.GetUser(ws)
	if err != nil {
		return
	}

	models.LeaveRoom(ws, user.RoomCode)
	models.DelUser(ws)
}
