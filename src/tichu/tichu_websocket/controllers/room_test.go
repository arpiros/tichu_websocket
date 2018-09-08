package controllers

import (
	"testing"
	"tichu/tichu_websocket/models"
)

func TestPop(t *testing.T) {
	cardDeck := models.NewCardDeck()

	for i := 0; i < 8*4; i++ {
		var poppedCard *models.Card
		poppedCard, cardDeck = Pop(cardDeck)
		println(poppedCard.CardType, poppedCard.Number, poppedCard.Color)
		println(len(cardDeck))
	}

	for i := 0; i < 6*4; i++ {
		var poppedCard *models.Card
		poppedCard, cardDeck = Pop(cardDeck)
		println(poppedCard.CardType, poppedCard.Number, poppedCard.Color)
		println(len(cardDeck))
	}
}

func TestDistributeCard(t *testing.T) {
	room := &models.Room{
		//Players: make([]*models.Player, 4),
		CardDeck: models.NewCardDeck(),
	}

	for i := 0; i < models.RoomMemberLimit; i++ {
		newPlayer := &models.Player{
			Index:      len(room.Players) + 1,
			TeamNumber: len(room.Players) % models.TeamCount,
			IsConnect:  true,
		}

		room.Players = append(room.Players, newPlayer)
	}


	DistributeCard(room, 8)
	DistributeCard(room, 6)
}