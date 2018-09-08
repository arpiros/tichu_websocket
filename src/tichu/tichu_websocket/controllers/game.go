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
	room.CallTichu[player.Index] = models.CallTichuNone

	if callLageTichu.IsCall {
		room.CallTichu[player.Index] = models.CallTichuLarge
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

func ChangeCard(ws *websocket.Conn, message []byte) {
	var changeCardReq protocol.ChangeCardReq
	err := json.Unmarshal(message, &changeCardReq)
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

	changeCardPlayerCount := 0
	if len(player.GainCard) == models.RoomMemberLimit - 1 {
		changeCardPlayerCount++
	}
	for playerIndex, cardIndex := range changeCardReq.Change {
		card := player.CardList[cardIndex]
		player.CardList = append(player.CardList[:cardIndex], player.CardList[cardIndex+1:]...)

		targetPlayer := room.Players[playerIndex]
		targetPlayer.GainCard = append(targetPlayer.GainCard, card)

		if len(targetPlayer.GainCard) == models.RoomMemberLimit - 1 {
			changeCardPlayerCount++
		}
	}


	if changeCardPlayerCount == models.RoomMemberLimit {
		for _, player := range room.Clients {
			player.CardList = append(player.CardList, player.GainCard...)

			if room.CurrentActivePlayer < 0 {
				for _, card := range player.CardList {
					if card.CardType == models.CardTypeSparrow {
						player.IsMyTurn = true
						room.CurrentActivePlayer = player.Index
						break
					}
				}
			}
		}

		for client, player := range room.Clients {
			client.WriteJSON(&protocol.StartGameResp{
				Player: player,
				CurrentActivePlayer: room.CurrentActivePlayer,
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
