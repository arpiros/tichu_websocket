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
		if !client.IsConnect {
			continue
		}

		inRoomUser.WriteJSON(&protocol.MoveTurnResp{
			Message: moveTurnReq.Message,
		})
	}
}

func StartGame(room *models.Room) {
	room.CardDeck = models.NewCardDeck()

	DistributeCard(room, 8)
	for key, value := range room.Clients {
		key.WriteJSON(value.CardList)
	}
}

func Pop(deck []*models.Card) (*models.Card, []*models.Card) {
	return deck[len(deck)-1], deck[:len(deck)-1]
}

func DistributeCard(room *models.Room, cardCount int) {
	for i := 0; i < cardCount; i++ {
		for _, value := range room.Players {
			var poppedCard *models.Card
			poppedCard, room.CardDeck = Pop(room.CardDeck)
			value.CardList = append(value.CardList, poppedCard)
		}
	}
}