package models

import (
	"github.com/stretchr/testify/assert"
	"math/rand"
	"testing"
)

func TestNewCardDeck(t *testing.T) {
	deck := NewCardDeck()

	for baseKey, baseCard := range deck {
		for compareKey, compareCard := range deck {
			if baseKey == compareKey {
				continue
			}

			if baseCard.Color == compareCard.Color &&
				baseCard.CardType == compareCard.CardType &&
				baseCard.Number == compareCard.Number {
				println("wrong deck")
			}
		}
	}

	assert.Equal(t, len(deck), TotalCardCount)
}

func TestIsBoom(t *testing.T) {
	try := 0
	for {
		deck := NewCardDeck()
		cardCount := 4
		if rand.Intn(2) == 1 {
			cardCount = 5
		}

		var cards []Card
		for i := 0; i < cardCount; i++ {
			cards = append(cards, *deck[i])
		}

		if IsBoom(cards) {
			for key, value := range cards {
				println(key, value.CardType, value.Color, value.Number)
			}

			println("try : ", try)

			if cardCount == 5 {
				return
			}
		}

		try++
	}
}
