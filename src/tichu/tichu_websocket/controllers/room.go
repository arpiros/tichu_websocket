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

	user, err := models.GetUser(ws)
	if err != nil {
		logrus.Errorf("CreateRoom GetUser Error : %v", err)
		return
	}

	if user.State != models.UserState_None {
		logrus.Errorf("CreateRoom Wrong State, State : %v", user.State)
		return
	}

	room := models.CreateRoom(ws)

	user.RoomCode = room.RoomCode
	user.State = models.UserState_InRoom

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

	user, err := models.GetUser(ws)
	if err != nil {
		logrus.Warnf("JoinRoom GetUser Error : %v", err)
		return
	}

	if user.State != models.UserState_None {
		logrus.Errorf("JoinRoom Wrong State, State : %v", user.State)
		return
	}

	models.JoinRoom(ws, joinRoomReq.RoomCode)
	user.RoomCode = joinRoomReq.RoomCode
	user.State = models.UserState_InRoom

	ws.WriteJSON(&protocol.JoinRoomResp{
		User: *user,
	})
}
