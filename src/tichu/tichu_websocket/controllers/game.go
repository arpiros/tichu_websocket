package controllers

import (
	"github.com/gorilla/websocket"
	"tichu/tichu_websocket/protocol"
	"encoding/json"
	"github.com/Sirupsen/logrus"
	"tichu/tichu_websocket/models"
)

func CallLargeTichu(ws *websocket.Conn, message []byte) {
	var callLageTichu protocol.CallLargeTichuReq
	err := json.Unmarshal(message, &callLageTichu)
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

	room, err := models.GetRoom(user.RoomCode)
	if err != nil {
		logrus.Warnf("MoveTurn GetRoom Error : %v", err)
		return
	}

	player := room.Clients[ws]
	room.CallTichu[player.PlayerIndex] = models.CallTichuNone

	if callLageTichu.IsCall {
		room.CallTichu[player.PlayerIndex] = models.CallTichuLarge
	}

	if len(room.CallTichu) == models.RoomMemberLimit {
		DistributeCard(room, 6)
	}

	for client, player := range room.Clients {
		if len(room.CallTichu) == models.RoomMemberLimit {
			client.WriteJSON(&protocol.DistributeAllCardResp{
				Player: player,
			})
		} else {
			client.WriteJSON(&protocol.CallLargeTichuResp{
				CallTichu: room.CallTichu,
			})
		}
	}
}

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
		if !client.IsConnect {
			continue
		}

		inRoomUser.WriteJSON(&protocol.MoveTurnResp{
			Message: moveTurnReq.Message,
		})
	}
}
