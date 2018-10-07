package controllers

import (
	"encoding/json"
	"github.com/Sirupsen/logrus"
	"github.com/gorilla/websocket"
	"tichu/tichu_websocket/models"
	"tichu/tichu_websocket/protocol"
)

func CallLargeTichu(ws *websocket.Conn, message []byte) {
	var req protocol.CallLargeTichuReq
	_, room, err := controllerInit(ws, message, &req)
	if err != nil {
		logrus.Println("CallLargeTichu controllerInit Error : ", err)
		models.DelUser(ws)
		return
	}

	player := room.Clients[ws]
	room.CallTichu[player.Index] = models.CallTichuNone

	if req.IsCall == 1 {
		room.CallTichu[player.Index] = models.CallTichuLarge
	}

	if len(room.CallTichu) == models.RoomMemberLimit {
		room.State = models.StateChangeCard
		DistributeCard(room, 6)
	}

	for client, player := range room.Clients {
		if len(room.CallTichu) == models.RoomMemberLimit {
			client.WriteJSON(&protocol.DistributeAllCardResp{
				BaseResp:  protocol.NewBaseResp(protocol.RespDistributeAllCard),
				Player:    player,
				CallTichu: room.CallTichu,
			})
		} else {
			client.WriteJSON(&protocol.CallLargeTichuResp{
				BaseResp:  protocol.NewBaseResp(protocol.RespCallLargeTichu),
				CallTichu: room.CallTichu,
			})
		}
	}
}

func ChangeCard(ws *websocket.Conn, message []byte) {
	var req protocol.ChangeCardReq
	_, room, err := controllerInit(ws, message, &req)
	if err != nil {
		logrus.Println("ChangeCard controllerInit Error : ", err)
		models.DelUser(ws)
		return
	}

	player := room.Clients[ws]

	changeCardPlayerCount := 0
	if len(player.GainCard) == models.RoomMemberLimit-1 {
		changeCardPlayerCount++
	}
	for playerIndex, cardIndex := range req.Change {
		card := player.CardList[cardIndex]
		player.CardList = append(player.CardList[:cardIndex], player.CardList[cardIndex+1:]...)

		targetPlayer := room.Players[playerIndex]
		targetPlayer.GainCard = append(targetPlayer.GainCard, card)

		if len(targetPlayer.GainCard) == models.RoomMemberLimit-1 {
			changeCardPlayerCount++
		}
	}

	if changeCardPlayerCount == models.RoomMemberLimit {
		for _, player := range room.Clients {
			player.CardList = append(player.CardList, player.GainCard...)

			if room.CurrentActivePlayer < 0 {
				for _, card := range player.CardList {
					if card.CardType == models.CardTypeMahjong {
						player.IsMyTurn = true
						room.CurrentActivePlayer = player.Index
						break
					}
				}
			}
		}

		room.State = models.StatePlaying
		for client, player := range room.Clients {
			client.WriteJSON(&protocol.StartGameResp{
				BaseResp:            protocol.NewBaseResp(protocol.RespStartGame),
				GainCard:            player.GainCard,
				CurrentActivePlayer: room.CurrentActivePlayer,
			})
		}
	}
}

func CallTichu(ws *websocket.Conn, message []byte) {
	var req protocol.CallTichuReq
	_, room, err := controllerInit(ws, message, &req)
	if err != nil {
		logrus.Println("CallTichu controllerInit Error : ", err)
		models.DelUser(ws)
		return
	}

	player := room.Clients[ws]
	if room.CallTichu[player.Index] > models.CallTichuNone {
		return
	}

	room.CallTichu[player.Index] = models.CallTichuSmall

	for client := range room.Clients {
		client.WriteJSON(&protocol.CallTichuResp{
			BaseResp:  protocol.NewBaseResp(protocol.RespCallTichu),
			CallTichu: room.CallTichu,
		})
	}
}

func SubmitCard(ws *websocket.Conn, message []byte) {
	var req protocol.SubmitCardReq
	_, room, err := controllerInit(ws, message, &req)
	if err != nil {
		logrus.Println("SubmitCard controllerInit Error : ", err)
		models.DelUser(ws)
		return
	}

	// isMyturn check
	player := room.Clients[ws]
	if !player.IsMyTurn {
		// TODO error
		return
	}

	var submitCards models.CardList
	for _, cardIndex := range req.Cards {
		submitCards = append(submitCards, player.CardList[cardIndex])
	}

	//preSubmitted card check
	if room.CanSubmitCard(submitCards) {
		//TODO error
		return
	}

	//change presubmitted card & remove submit card
	submit := models.Submit{
		PlayerIndex: player.Index,
		Class:       models.SubmitClassSingle,
		Cards:       submitCards,
		LastNumber:  1,
	}

	room.Submits = append(room.Submits, submit)

	for _, cardIndex := range req.Cards {
		player.CardList = append(player.CardList[:cardIndex], player.CardList[cardIndex+1:]...)
	}

	MoveNextTurn(room)

	////check game end
	//endPlayerCount := 0
	//for _, p := range room.Players {
	//	if len(p.CardList) == 0 {
	//		endPlayerCount++
	//	}
	//}
	//
	//if endPlayerCount >= 3 {
	//	// end game process
	//
	//} else {
	//	MoveNextTurn(room)
	//}
}

func MoveNextTurn(room *models.Room) {
	room.MoveTurn()
	for client, p := range room.Clients {
		if p.Index == room.CurrentActivePlayer {
			p.IsMyTurn = true
		}

		client.WriteJSON(&protocol.SubmitCardResp{
			BaseResp:            protocol.NewBaseResp(protocol.RespSubmitCard),
			Player:              p,
			CurrentActivePlayer: room.CurrentActivePlayer,
		})
	}
}

func UseBoom(ws *websocket.Conn, message []byte) {
	var req protocol.UseBoomReq
	_, room, err := controllerInit(ws, message, &req)
	if err != nil {
		logrus.Println("UseBoom controllerInit Error : ", err)
		models.DelUser(ws)
		return
	}

	player := room.Clients[ws]

	var cards models.CardList
	for _, value := range req.Cards {
		cards = append(cards, player.CardList[value])
	}

	if !models.IsBoom(cards) {
		// invalid boom
		return
	}

	// TODO send boom result
}

func MoveTurn(ws *websocket.Conn, message []byte) {
	var req protocol.MoveTurnReq
	_, room, err := controllerInit(ws, message, &req)
	if err != nil {
		logrus.Println("MoveTurn controllerInit Error : ", err)
		models.DelUser(ws)
		return
	}

	for inRoomUser, client := range room.Clients {
		if !client.IsConnect {
			continue
		}

		inRoomUser.WriteJSON(&protocol.MoveTurnResp{
			Message: req.Message,
		})
	}
}

func controllerInit(ws *websocket.Conn, message []byte, _struct interface{}) (*models.User, *models.Room, error) {
	err := json.Unmarshal(message, &_struct)
	if err != nil {
		return nil, nil, err
	}

	user, err := models.GetUser(ws)
	if err != nil {
		return nil, nil, err
	}

	room, err := models.GetRoom(user.RoomCode)
	if err != nil {
		return nil, nil, err
	}

	return user, room, nil
}
