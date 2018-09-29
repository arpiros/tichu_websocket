package models

import (
	"github.com/stretchr/testify/assert"
	"math/rand"
	"sort"
	"testing"
	"tichu/tichu_websocket/util"
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

		var cards CardList
		for i := 0; i < cardCount; i++ {
			cards = append(cards, deck[i])
		}

		if IsBoom(cards) {
			sort.Sort(cards)
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

func TestIsStrait(t *testing.T) {
	try := 0
	for {
		deck := NewCardDeck()
		cardCount := util.RandomRange(1, 13)

		var cards CardList
		for i := 0; i < cardCount; i++ {
			cards = append(cards, deck[i])
		}

		if IsStrait(cards) {
			sort.Sort(cards)
			for key, value := range cards {
				println(key, value.CardType, value.Color, value.Number)
			}

			println("try : ", try)
			return
		}

		try++
	}
}

func TestIsSingle(t *testing.T) {
	try := 0
	for {
		deck := NewCardDeck()
		cardCount := util.RandomRange(1, 13)

		var cards CardList
		for i := 0; i < cardCount; i++ {
			cards = append(cards, deck[i])
		}

		if IsSingle(cards) {
			sort.Sort(cards)
			for key, value := range cards {
				println(key, value.CardType, value.Color, value.Number)
			}

			println("try : ", try)
			return
		}

		try++
	}
}

func TestIsPair(t *testing.T) {
	try := 0
	for {
		deck := NewCardDeck()
		cardCount := util.RandomRange(1, 13)

		var cards CardList
		for i := 0; i < cardCount; i++ {
			cards = append(cards, deck[i])
		}

		if IsPair(cards) {
			sort.Sort(cards)
			for key, value := range cards {
				println(key, value.CardType, value.Color, value.Number)
			}

			println("try : ", try)
			return
		}

		try++
	}
}

func TestIsTriple(t *testing.T) {
	try := 0
	for {
		deck := NewCardDeck()
		cardCount := util.RandomRange(1, 13)

		var cards CardList
		for i := 0; i < cardCount; i++ {
			cards = append(cards, deck[i])
		}

		if IsTriple(cards) {
			sort.Sort(cards)
			for key, value := range cards {
				println(key, value.CardType, value.Color, value.Number)
			}

			println("try : ", try)
			return
		}

		try++
	}
}


func TestIsStraitPair(t *testing.T) {
	try := 0
	for {
		deck := NewCardDeck()
		cardCount := 4

		var cards CardList
		for i := 0; i < cardCount; i++ {
			cards = append(cards, deck[i])
		}

		if IsStraitPair(cards) {
			sort.Sort(cards)
			for key, value := range cards {
				println(key, value.CardType, value.Color, value.Number)
			}

			println("try : ", try)
			return
		}

		try++
	}
}

func TestIsFullHouse(t *testing.T) {
	try := 0
	for {
		deck := NewCardDeck()
		cardCount := 5

		var cards CardList
		for i := 0; i < cardCount; i++ {
			cards = append(cards, deck[i])
		}

		if IsFullHouse(cards) {
			sort.Sort(cards)
			for key, value := range cards {
				println(key, value.CardType, value.Color, value.Number)
			}

			println("try : ", try)
			return
		}

		try++
	}
}
