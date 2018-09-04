package controllers

import (
	"github.com/gorilla/websocket"
	"tichu/tichu_websocket/protocol"
	"encoding/json"
	"github.com/Sirupsen/logrus"
	"tichu/tichu_websocket/models"
)

func MoveTurn(ws *websocket.Conn, message []byte) {
	var moveTurnReq protocol.MoveTurnReq
	err := json.Unmarshal(message, &moveTurnReq)
	if err != nil {
		logrus.Println("Unmarshal Error : ", err)
		models.DelUser(ws)
		return
	}

	user, err := models.GetUser(ws)
	if err != nil {
		logrus.Warnf("MoveTurn GetUser Error : %v", err)
		return
	}

	//TODO check user state
	room, err := models.GetRoom(user.RoomCode)
	if err != nil {
		logrus.Warnf("MoveTurn GetRoom Error : %v", err)
		return
	}

	for inRoomUser, client := range room.Clients {
		if !client {
			continue
		}

		inRoomUser.WriteJSON(&protocol.MoveTurnResp{
			Message: moveTurnReq.Message,
		})
	}
}