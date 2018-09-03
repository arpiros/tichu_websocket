package controllers

import (
	"tichu/tichu_websocket/protocol"
	"encoding/json"
	"tichu/tichu_websocket/models"
	"github.com/gorilla/websocket"
	"github.com/Sirupsen/logrus"
)

func CreateRoom(ws *websocket.Conn, message []byte) {
	var createRoomReq protocol.CreateRoomReq
	err := json.Unmarshal(message, &createRoomReq)
	if err != nil {
		logrus.Println("Unmarshal Error : ", err)
		models.DelUser(ws)
		return
	}
	room := models.CreateRoom(ws)
	user, err := models.GetUser(ws)
	if err != nil {

	}
	user.RoomCode = room.RoomCode

	ws.WriteJSON(&protocol.CreateRoomResp{
		RoomCode: room.RoomCode,
	})
}

func JoinRoom(ws *websocket.Conn, message []byte) {
	var joinRoomReq protocol.JoinRoomReq
	err := json.Unmarshal(message, &joinRoomReq)
	if err != nil {
		logrus.Println("Unmarshal Error : ", err)
		models.DelUser(ws)
		return
	}

	models.JoinRoom(ws, joinRoomReq.RoomCode)
}
